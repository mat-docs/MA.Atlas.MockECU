// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const logStackTraceSettingName = "Log.EnableStackTrace"
const logStackTraceSettingCmdline = "log-stacktrace"

// Cached value
var logStackTraceSettingCache = core.NewCachedValue(func() bool { return viper.GetBool(logStackTraceSettingName) })

// Fetch the setting
func (theSettings) LogStackTrace() bool {
	return logStackTraceSettingCache.GetValue()
}

func AddLogStackTraceFlag(flagSet *pflag.FlagSet) {
	viperEx.BoolSetting(logStackTraceSettingName, "Log stack trace for errors").Cmdline(logStackTraceSettingCmdline).AddTo(flagSet)
}

// SECTION-END
