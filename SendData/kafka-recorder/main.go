// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package main

import (
	"os"

	"gitlab.com/atrico/container"
	"gitlab.com/atrico/core/errorsEx"

	"gitlab.com/atrico/kafka-recorder/api"
	"gitlab.com/atrico/kafka-recorder/cmd"
	"gitlab.com/atrico/kafka-recorder/settings"
)

func main() {
	c := register()
	var cmdFactory cmd.CommandFactory
	c.Make(&cmdFactory)
	cmd := cmdFactory.Create()

	err := cmd.Execute()
	var returnCode int
	if errEx, ok := err.(errorsEx.Error); ok {
		returnCode = errEx.ReturnCode()
	} else if err != nil {
		returnCode = 1
	}
	os.Exit(returnCode)
}

func register() container.Container {
	c := container.NewContainer()
	settings.RegisterSettings(c)
	settings.RegisterVerboseService(c)
	cmd.RegisterCmd(c)
	api.RegisterApiFactories(c)
	return c
}
