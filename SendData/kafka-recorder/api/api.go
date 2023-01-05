// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package api

import (
	"gitlab.com/atrico/container"

	"gitlab.com/atrico/kafka-recorder/settings"
)

type Runnable interface {
	Run() error
}

type Factory interface {
	Create() Runnable
}

func RegisterApiFactories(c container.Container) {
	settings.RegisterVerboseService(c)
	c.Singleton(func() RecordApiFactory { return recordApiFactory{c} })
	c.Singleton(func() ReplayApiFactory { return replayApiFactory{c} })
}
