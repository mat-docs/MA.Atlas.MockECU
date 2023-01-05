// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const brokersSettingName = "Brokers"
const brokersSettingCmdline = "broker"
const brokersSettingShortcut = 'b'
const brokersSettingEnvVar = "BROKERS"

// Cached value
var brokersSettingCache = core.NewCachedValue(func() []string { return viperEx.GetStringSlice(brokersSettingName) })

// Fetch the setting
func (theSettings) Brokers() []string {
	return brokersSettingCache.GetValue()
}

func AddBrokersFlag(flagSet *pflag.FlagSet) {
	viperEx.StringArraySetting(brokersSettingName, "Kafka brokers").Cmdline(brokersSettingCmdline).CmdlineShortcut(brokersSettingShortcut).EnvVar(brokersSettingEnvVar).AddTo(flagSet)
}

// SECTION-END
