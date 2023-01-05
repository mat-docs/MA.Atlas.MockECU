// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const logLevelSettingName = "Log.Level"
const logLevelSettingCmdline = "log-level"
const logLevelSettingEnvVar = "GO_LOGLEVEL"
const logLevelSettingDefaultVal = "info"

// Cached value
var logLevelSettingCache = core.NewCachedValue(func() string { return viper.GetString(logLevelSettingName) })

// Fetch the setting
func (theSettings) LogLevel() string {
	return logLevelSettingCache.GetValue()
}

func AddLogLevelFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(logLevelSettingName, "Level of logging to generate (panic,fatal,error,warning,info,debug,trace)").Cmdline(logLevelSettingCmdline).EnvVar(logLevelSettingEnvVar).DefaultVal(logLevelSettingDefaultVal).AddTo(flagSet)
}

// SECTION-END
