// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/atrico/container"

	"gitlab.com/atrico/kafka-recorder/api"
)

type RecordCmd commandInfo

func RegisterCmdRecord(c container.Container) {
	c.Singleton(func(apiFactory api.RecordApiFactory) RecordCmd { return RecordCmd(createRecordCommand(apiFactory)) })
}

func createRecordCommand(apiFactory api.Factory) commandInfo {
	cmd := &cobra.Command{
		Use:   "record",
		Short: "Record kafka messages",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return apiFactory.Create().Run()
		},
	}
	return commandInfo{cmd, "record"}
}
