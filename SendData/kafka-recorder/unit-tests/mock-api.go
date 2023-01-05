// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package unit_tests

import (
	"gitlab.com/atrico/container"

	"gitlab.com/atrico/kafka-recorder/api"
	"gitlab.com/atrico/kafka-recorder/settings"
)

type mockApiSettings interface {
	settings.RecordSettings
	settings.ReplaySettings
	settings.RootSettings
}

type mockApi struct {
	cmd    []string
	config mockApiSettings
}

type mockApiFactory mockApi

var results map[string]interface{}

func (m mockApi) Run() error {
	results = make(map[string]interface{})
	if m.cmd[0] == "root" {
		results["TheCommand"] = []string{""}
	} else {
		results["TheCommand"] = m.cmd
	}
	results["File"] = m.config.File()
	results["Brokers"] = m.config.Brokers()
	results["Topic"] = m.config.Topic()
	results["Timeout"] = m.config.Timeout()
	results["ConfigFile"] = m.config.ConfigFile()
	results["Verbose"] = m.config.Verbose()
	results["LogLevel"] = m.config.LogLevel()
	results["LogWriter"] = m.config.LogWriter()
	results["LogTimeFormat"] = m.config.LogTimeFormat()
	results["LogStackTrace"] = m.config.LogStackTrace()
	return nil
}

func (f mockApiFactory) Create() api.Runnable {
	return mockApi(f)
}

func registerMockApiFactories(c container.Container) {
	c.Singleton(func(config settings.RootSettings) api.RecordApiFactory {
		return mockApiFactory{[]string{"record"}, config.(mockApiSettings)}
	})
	c.Singleton(func(config settings.RootSettings) api.ReplayApiFactory {
		return mockApiFactory{[]string{"replay"}, config.(mockApiSettings)}
	})
}
