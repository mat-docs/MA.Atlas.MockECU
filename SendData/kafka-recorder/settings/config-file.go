// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const configFileSettingName = "ConfigFile"
const configFileSettingCmdline = "config-file"

// Cached value
var configFileSettingCache = core.NewCachedValue(func() string { return viper.GetString(configFileSettingName) })

// Fetch the setting
func (theSettings) ConfigFile() string {
	return configFileSettingCache.GetValue()
}

func AddConfigFileFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(configFileSettingName, "Alternate config file").Cmdline(configFileSettingCmdline).AddTo(flagSet)
}

// SECTION-END
