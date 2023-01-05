// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package api

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rcrowley/go-metrics"
	"github.com/rs/zerolog/log"
	"gitlab.com/atrico/container"

	"gitlab.com/atrico/kafka-recorder/api/internal"
	"gitlab.com/atrico/kafka-recorder/settings"
)

type ReplayApi Runnable
type ReplayApiFactory Factory

type replayApiFactory struct {
	container.Container
}

func (f replayApiFactory) Create() Runnable {
	RegisterApiReplay(f.Container)
	var theApi ReplayApi
	f.Container.Make(&theApi)
	return theApi
}

// SECTION-END

func RegisterApiReplay(c container.Container) {
	internal.RegisterCommon(c)
	c.Singleton(func(config settings.ReplaySettings, verboseService settings.VerboseService) ReplayApi {
		return replayApi{config, verboseService}
	})
}

type replayApi struct {
	settings.ReplaySettings
	settings.VerboseService
}

// Replay kafka messages
func (svc replayApi) Run() (err error) {
	log.Debug().Msgf("Replaying file %s to %s/%s", svc.File(), svc.Brokers(), svc.Topic())
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Metadata.Full = false
	metrics.UseNilMetrics = true
	var producer sarama.AsyncProducer
	if producer, err = sarama.NewAsyncProducer(svc.Brokers(), config); err == nil {
		var file *os.File
		if file, err = os.Open(svc.File()); err == nil {
			defer file.Close()
			var key, value []byte
			more := true
			for more && err == nil {
				if key, value, err = internal.ReadMessage(file); err == nil {
					log.Trace().Msgf("Writing %s => %s", string(key), string(value))
					pMsg := sarama.ProducerMessage{
						Topic: svc.Topic(),
						Key:   dataWrapper(key),
						Value: dataWrapper(value),
					}
					producer.Input() <- &pMsg
					select {
					case <-producer.Successes():
					// Nothing to do
					case err = <-producer.Errors():
						if err != io.EOF {
							log.Error().Err(err).Msg("Producer error")
						}
					case <-time.After(time.Millisecond * time.Duration(svc.Timeout())):
						log.Debug().Msg("Timeout")
						err = errors.New("producer timeout")
					}
				}
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return err
}

type dataWrapper []byte

func (d dataWrapper) Encode() ([]byte, error) {
	return d, nil
}
func (d dataWrapper) Length() int {
	return len(d)
}
