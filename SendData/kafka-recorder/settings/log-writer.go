// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const logWriterSettingName = "Log.Writer"
const logWriterSettingCmdline = "log-writer"
const logWriterSettingDefaultVal = "stdout"

// Cached value
var logWriterSettingCache = core.NewCachedValue(func() string { return viper.GetString(logWriterSettingName) })

// Fetch the setting
func (theSettings) LogWriter() string {
	return logWriterSettingCache.GetValue()
}

func AddLogWriterFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(logWriterSettingName, "Writer for logger (stdout[:json], stderr[:json], file[:json]=path)").Cmdline(logWriterSettingCmdline).DefaultVal(logWriterSettingDefaultVal).AddTo(flagSet)
}

// SECTION-END
