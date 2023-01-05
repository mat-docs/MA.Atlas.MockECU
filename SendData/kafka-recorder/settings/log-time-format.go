// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const logTimeFormatSettingName = "Log.TimeFormat"
const logTimeFormatSettingCmdline = "log-timeformat"
const logTimeFormatSettingDefaultVal = "2006-01-02T15:04:05Z07:00"

// Cached value
var logTimeFormatSettingCache = core.NewCachedValue(func() string { return viper.GetString(logTimeFormatSettingName) })

// Fetch the setting
func (theSettings) LogTimeFormat() string {
	return logTimeFormatSettingCache.GetValue()
}

func AddLogTimeFormatFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(logTimeFormatSettingName, "Format for timestamp of log entries").Cmdline(logTimeFormatSettingCmdline).DefaultVal(logTimeFormatSettingDefaultVal).AddTo(flagSet)
}

// SECTION-END
