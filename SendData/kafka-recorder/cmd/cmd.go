// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package cmd

import (
	"path"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/atrico/container"
)

// Create all the commands
type CommandFactory interface {
	Create() *cobra.Command
}

type commandInfo struct {
	cmd  *cobra.Command
	path string
}
type commandFactory []commandInfo

// Register Commands
func RegisterCmd(c container.Container) {
	RegisterCmdVersion(c)
	RegisterCmdRecord(c)
	RegisterCmdReplay(c)
	RegisterCmdRoot(c)
	c.Singleton(func(version VersionCmd, record RecordCmd, replay ReplayCmd, root RootCmd) CommandFactory {
		cmdFactory := commandFactory{
			commandInfo(version),
			commandInfo(record),
			commandInfo(replay),
			commandInfo(root),
		}
		sort.Slice(cmdFactory, func(i, j int) (less bool) {
			// Sort by shortest path first, then alphabetically (root first)
			lenDiff := len(strings.Split(cmdFactory[i].path, "/")) - len(strings.Split(cmdFactory[j].path, "/"))
			if lenDiff != 0 {
				return lenDiff < 0
			}
			return cmdFactory[i].path < cmdFactory[j].path
		})
		return cmdFactory
	})
}

func (c commandFactory) Create() *cobra.Command {
	cobra.OnInitialize(initConfig)
	commands := make(map[string]*cobra.Command, 1)
	for _, cmdInfo := range c {
		if cmdInfo.path != "." {
			parent := path.Dir(cmdInfo.path)
			commands[parent].AddCommand(cmdInfo.cmd)
		}
		commands[cmdInfo.path] = cmdInfo.cmd
	}
	return commands["."]
}
