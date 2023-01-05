// Generated 2023-01-05 15:14:25 by go-framework v2.1.3
// SECTION-START: Framework
package api

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"gitlab.com/atrico/container"

	"gitlab.com/atrico/kafka-recorder/api/internal"
	"gitlab.com/atrico/kafka-recorder/settings"
)

type RecordApi Runnable
type RecordApiFactory Factory

type recordApiFactory struct {
	container.Container
}

func (f recordApiFactory) Create() Runnable {
	RegisterApiRecord(f.Container)
	var theApi RecordApi
	f.Container.Make(&theApi)
	return theApi
}

// SECTION-END

func RegisterApiRecord(c container.Container) {
	internal.RegisterCommon(c)
	c.Singleton(func(config settings.RecordSettings, verboseService settings.VerboseService) RecordApi {
		return recordApi{config, verboseService}
	})
}

type recordApi struct {
	settings.RecordSettings
	settings.VerboseService
}

// Record kafka messages
func (svc recordApi) Run() (err error) {
	config := sarama.NewConfig()
	var master sarama.Consumer
	var consumer sarama.PartitionConsumer
	if master, err = sarama.NewConsumer(svc.Brokers(), config); err == nil {
		if consumer, err = master.ConsumePartition(svc.Topic(), 0, 0); err == nil {
			var file *os.File
			if file, err = os.Create(svc.File()); err == nil {
				defer file.Close()
				more := true
				for more && err == nil {
					// Read data
					select {
					case kMsg, open := <-consumer.Messages():
						if open {
							err = internal.WriteMessage(file, kMsg)
						}
					case err = <-consumer.Errors():
						log.Error().Err(err).Msg("Consumer error")
					case <-time.After(time.Millisecond * time.Duration(svc.Timeout())):
						log.Debug().Msg("Timeout")
						more = false
					}
				}
			}
		}
	}
	return err
}
