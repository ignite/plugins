package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ignite/apps/hermes/pkg/hermes"
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
	c.PersistentFlags().StringP(flagConfig, "c", hermes.DefaultConfigPath(), "path to Hermes config file")

	// add sub commands.
	c.AddCommand(
		NewHermesConfigure(),
		NewHermesStart(),
		NewHermesKeys(),
		NewHermesExecute(),
	)

	return c
}

func getConfig(cmd *cobra.Command) string {
	config, _ := cmd.Flags().GetString(flagConfig)
	return config
}