// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const verboseSettingName = "Verbose"
const verboseSettingCmdline = "verbose"
const verboseSettingShortcut = 'v'
const verboseSettingEnvVar = "VERBOSE"

// Cached value
var verboseSettingCache = core.NewCachedValue(func() bool { return viper.GetBool(verboseSettingName) })

// Fetch the setting
func (theSettings) Verbose() bool {
	return verboseSettingCache.GetValue()
}

func AddVerboseFlag(flagSet *pflag.FlagSet) {
	viperEx.BoolSetting(verboseSettingName, "Generate more detailed output").Cmdline(verboseSettingCmdline).CmdlineShortcut(verboseSettingShortcut).EnvVar(verboseSettingEnvVar).AddTo(flagSet)
}

// SECTION-END
