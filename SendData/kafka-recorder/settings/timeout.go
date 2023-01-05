// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const timeoutSettingName = "Timeout"
const timeoutSettingCmdline = "timeout"
const timeoutSettingShortcut = 't'
const timeoutSettingEnvVar = "TIMEOUT"
const timeoutSettingDefaultVal = 30000

// Cached value
var timeoutSettingCache = core.NewCachedValue(func() int { return viper.GetInt(timeoutSettingName) })

// Fetch the setting
func (theSettings) Timeout() int {
	return timeoutSettingCache.GetValue()
}

func AddTimeoutFlag(flagSet *pflag.FlagSet) {
	viperEx.IntSetting(timeoutSettingName, "Timeout (ms)").Cmdline(timeoutSettingCmdline).CmdlineShortcut(timeoutSettingShortcut).EnvVar(timeoutSettingEnvVar).DefaultVal(timeoutSettingDefaultVal).AddTo(flagSet)
}

// SECTION-END
