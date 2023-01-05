// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const fileSettingName = "File"
const fileSettingCmdline = "file"
const fileSettingShortcut = 'f'
const fileSettingEnvVar = "DATAFILE"

// Cached value
var fileSettingCache = core.NewCachedValue(func() string { return viper.GetString(fileSettingName) })

// Fetch the setting
func (theSettings) File() string {
	return fileSettingCache.GetValue()
}

func AddFileFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(fileSettingName, "Record file").Cmdline(fileSettingCmdline).CmdlineShortcut(fileSettingShortcut).EnvVar(fileSettingEnvVar).AddTo(flagSet)
}

// SECTION-END
