package cmd

import (
	"github.com/ignite/cli/v28/ignite/services/plugin"
)

// GetCommands returns the list of app commands.
func GetCommands() []*plugin.Command {
	return []*plugin.Command{
		{
			Use:      "connect [command]",
			Aliases:  []string{"c"},
			Short:    "Interact with any Cosmos SDK based blockchain using Ignite Connect",
			Long:     "Ignite Connect is an app that allows you to interact with any Cosmos SDK based blockchain.\n It leverages AutoCLI from client/v2 and is inspired by the Hubl tool",
			Commands: []*plugin.Command{},
		},
	}
}