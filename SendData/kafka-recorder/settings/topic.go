// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/atrico/core"
	"gitlab.com/atrico/viperEx/v2"
)

const topicSettingName = "Topic"
const topicSettingCmdline = "topic"
const topicSettingEnvVar = "TOPIC"

// Cached value
var topicSettingCache = core.NewCachedValue(func() string { return viper.GetString(topicSettingName) })

// Fetch the setting
func (theSettings) Topic() string {
	return topicSettingCache.GetValue()
}

func AddTopicFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(topicSettingName, "Kafka topic").Cmdline(topicSettingCmdline).EnvVar(topicSettingEnvVar).AddTo(flagSet)
}

// SECTION-END
