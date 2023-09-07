package cmd

import (
	"github.com/spf13/cobra"
)

const (
	flagConfig = "config"
)

// NewRelayer creates a new relayer command that holds
// some other sub commands related to hermes relayer.
func NewRelayer() *cobra.Command {
	c := &cobra.Command{
		Use:           "hermes [command]",
		Aliases:       []string{"h"},
		Short:         "Hermes relayer wrapper",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// configure flags.
	// TODO if add the config file don't need to create it
	c.PersistentFlags().StringP(flagConfig, "c", "", "Use a custom config instead create a new one")

	// add sub commands.
	c.AddCommand(
		NewHermesKeys(),
		NewHermesConfigure(),
		NewHermesStart(),
		NewHermesExecute(),
	)

	return c
}