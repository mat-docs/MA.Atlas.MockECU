// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package common

type MockSettings struct {
	TheCommand       []string
	FileVar          string
	BrokersVar       []string
	TopicVar         string
	TimeoutVar       int
	ConfigFileVar    string
	VerboseVar       bool
	LogLevelVar      string
	LogWriterVar     string
	LogTimeFormatVar string
	LogStackTraceVar bool
}

func (s MockSettings) File() string {
	return s.FileVar
}
func (s MockSettings) Brokers() []string {
	return s.BrokersVar
}
func (s MockSettings) Topic() string {
	return s.TopicVar
}
func (s MockSettings) Timeout() int {
	return s.TimeoutVar
}
func (s MockSettings) ConfigFile() string {
	return s.ConfigFileVar
}
func (s MockSettings) Verbose() bool {
	return s.VerboseVar
}
func (s MockSettings) LogLevel() string {
	return s.LogLevelVar
}
func (s MockSettings) LogWriter() string {
	return s.LogWriterVar
}
func (s MockSettings) LogTimeFormat() string {
	return s.LogTimeFormatVar
}
func (s MockSettings) LogStackTrace() bool {
	return s.LogStackTraceVar
}

func NewMockSettings(cmd []string) MockSettings {
	return MockSettings{
		TheCommand:       cmd,
		BrokersVar:       make([]string, 0),
		TimeoutVar:       30000,
		LogLevelVar:      "info",
		LogWriterVar:     "stdout",
		LogTimeFormatVar: "2006-01-02T15:04:05Z07:00",
	}
}
