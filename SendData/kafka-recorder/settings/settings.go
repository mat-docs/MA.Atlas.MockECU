// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package settings

import (
	"gitlab.com/atrico/container"
)

type RecordSettings interface {
	RootSettings
}

type ReplaySettings interface {
	RootSettings
}

type RootSettings interface {
	// Record file
	File() string
	// Kafka brokers
	Brokers() []string
	// Kafka topic
	Topic() string
	// Timeout (ms)
	Timeout() int
	// Alternate config file
	ConfigFile() string
	// Generate more detailed output
	Verbose() bool
	// Level of logging to generate (panic,fatal,error,warning,info,debug,trace)
	LogLevel() string
	// Writer for logger (stdout[:json], stderr[:json], file[:json]=path)
	LogWriter() string
	// Format for timestamp of log entries
	LogTimeFormat() string
	// Log stack trace for errors
	LogStackTrace() bool
}

// Register the settings
func RegisterSettings(c container.Container) {
	settings := theSettings{}
	c.Singleton(func() RecordSettings { return &settings })
	c.Singleton(func() ReplaySettings { return &settings })
	c.Singleton(func() RootSettings { return &settings })
}

// Force all settings to be recalculated on next request
func ResetCaches() {
	fileSettingCache.Reset()
	brokersSettingCache.Reset()
	topicSettingCache.Reset()
	timeoutSettingCache.Reset()
	configFileSettingCache.Reset()
	verboseSettingCache.Reset()
	logLevelSettingCache.Reset()
	logWriterSettingCache.Reset()
	logTimeFormatSettingCache.Reset()
	logStackTraceSettingCache.Reset()
}

// Stub object for settings interface
type theSettings struct{}
