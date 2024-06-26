package cmd

import (
	"github.com/ignite/cli/v28/ignite/services/plugin"
	"github.com/spf13/pflag"
)

const (
	flagConfig = "config"
)

// GetCommands returns the list of explorer app commands.
func GetCommands() []*plugin.Command {
	return []*plugin.Command{
		{
			Use:     "hermes [command]",
			Short:   "Connect blockchains with an IBC Hermes Relayer",
			Aliases: []string{"h"},
			Flags: []*plugin.Flag{
				{Name: flagConfig, Usage: "set a custom Hermes config file", Shorthand: "c", Type: plugin.FlagTypeString},
			},
			Commands: []*plugin.Command{
				{
					Use:   "exec [args...]",
					Short: "Execute a hermes raw command",
				},
				{
					Use:   "start [chain-a-id] [chain-a-rpc]",
					Short: "Start the Hermes relayer",
				},
				{
					Use:   "keys [command]",
					Short: "Start the Hermes relayer",
					Commands: []*plugin.Command{
						{
							Use:   "add [chain-id] [mnemonic]",
							Short: "Add a new key from mnemonic to Hermes relayer",
						},
						{
							Use:   "file [chain-id] [filepath]",
							Short: "Add a new key from a key file to Hermes relayer",
						},
						{
							Use:   "list [chain-id]",
							Short: "List Hermes relayer keys",
						},
						{
							Use:   "delete [chain-id] [key-name]",
							Short: "Delete a key from Hermes relayer",
						},
					},
				},
				{
					Use:   "configure [chain-a-id] [chain-a-rpc] [chain-a-grpc] [chain-b-id] [chain-b-rpc] [chain-b-grpc]",
					Short: "Configure the Hermes relayer creating the config file, client, channels and connection",
					Flags: []*plugin.Flag{
						{Name: flagChainAPortID, DefaultValue: "transfer", Usage: "port ID of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBPortID, DefaultValue: "transfer", Usage: "port ID of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAPortID, DefaultValue: "transfer", Usage: "port ID of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainACCVConsumerChain, DefaultValue: "false", Usage: "only specify true if the chain A is a CCV consumer", Type: plugin.FlagTypeBool},
						{Name: flagChainBCCVConsumerChain, DefaultValue: "false", Usage: "only specify true if the chain B is a CCV consumer", Type: plugin.FlagTypeBool},
						{Name: flagChainAEventSourceURL, Usage: "WS event source url of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBEventSourceURL, Usage: "WS event source url of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAEventSourceMode, DefaultValue: "push", Usage: "WS event source mode of the chain A (event source url should be set to use this flag)", Type: plugin.FlagTypeString},
						{Name: flagChainBEventSourceMode, DefaultValue: "push", Usage: "WS event source mode of the chain B (event source url should be set to use this flag)", Type: plugin.FlagTypeString},
						{Name: flagChainAEventSourceBatchDelay, DefaultValue: "500ms", Usage: "WS event source batch delay time of the chain A (event source url should be set to use this flag)", Type: plugin.FlagTypeString},
						{Name: flagChainBEventSourceBatchDelay, DefaultValue: "500ms", Usage: "WS event source batch delay time of the chain B (event source url should be set to use this flag)", Type: plugin.FlagTypeString},
						{Name: flagChainARPCTimeout, DefaultValue: "10s", Usage: "RPC timeout of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBRPCTimeout, DefaultValue: "10s", Usage: "RPC timeout of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainATrustedNode, DefaultValue: "false", Usage: "enable trusted node on the chain A", Type: plugin.FlagTypeBool},
						{Name: flagChainBTrustedNode, DefaultValue: "false", Usage: "enable trusted node on the chain B", Type: plugin.FlagTypeBool},
						{Name: flagChainAAccountPrefix, DefaultValue: "cosmos", Usage: "account prefix of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBAccountPrefix, DefaultValue: "cosmos", Usage: "account prefix of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAKeyName, DefaultValue: "wallet", Usage: "hermes account name of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBKeyName, DefaultValue: "wallet", Usage: "hermes account name of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAAddressType, DefaultValue: "cosmos", Usage: "address type of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBAddressType, DefaultValue: "cosmos", Usage: "address type of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAStorePrefix, DefaultValue: "ibc", Usage: "store prefix of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBStorePrefix, DefaultValue: "ibc", Usage: "store prefix of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainADefaultGas, DefaultValue: "100000", Usage: "default gas used for transactions on chain A", Type: plugin.FlagTypeUint64},
						{Name: flagChainBDefaultGas, DefaultValue: "100000", Usage: "default gas used for transactions on chain B", Type: plugin.FlagTypeUint64},
						{Name: flagChainAMaxGas, DefaultValue: "400000", Usage: "max gas used for transactions on chain A", Type: plugin.FlagTypeUint64},
						{Name: flagChainBMaxGas, DefaultValue: "400000", Usage: "max gas used for transactions on chain B", Type: plugin.FlagTypeUint64},
						{Name: flagChainAGasPrice, DefaultValue: "0.025stake", Usage: "gas price used for transactions on chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBGasPrice, DefaultValue: "0.025stake", Usage: "gas price used for transactions on chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAGasMultiplier, DefaultValue: "1.1", Usage: "gas multiplier used for transactions on chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBGasMultiplier, DefaultValue: "1.1", Usage: "gas multiplier used for transactions on chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAMaxMsgNum, DefaultValue: "30", Usage: "max message number used for transactions on chain A", Type: plugin.FlagTypeUint64},
						{Name: flagChainBMaxMsgNum, DefaultValue: "30", Usage: "max message number used for transactions on chain B", Type: plugin.FlagTypeUint64},
						{Name: flagChainAMaxTxSize, DefaultValue: "2097152", Usage: "max transaction size on chain A", Type: plugin.FlagTypeUint64},
						{Name: flagChainBMaxTxSize, DefaultValue: "2097152", Usage: "max transaction size on chain B", Type: plugin.FlagTypeUint64},
						{Name: flagChainAClockDrift, DefaultValue: "5s", Usage: "clock drift of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBClockDrift, DefaultValue: "5s", Usage: "clock drift of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAMaxBlockTime, DefaultValue: "30s", Usage: "maximum block time of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBMaxBlockTime, DefaultValue: "30s", Usage: "maximum block time of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainATrustingPeriod, DefaultValue: "14days", Usage: "trusting period of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBTrustingPeriod, DefaultValue: "14days", Usage: "trusting period of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainATrustThresholdNumerator, DefaultValue: "2", Usage: "trusting threshold numerator of the chain A", Type: plugin.FlagTypeUint64},
						{Name: flagChainBTrustThresholdNumerator, DefaultValue: "2", Usage: "trusting threshold numerator of the chain B", Type: plugin.FlagTypeUint64},
						{Name: flagChainATrustThresholdDenominator, DefaultValue: "3", Usage: "trusting threshold denominator of the chain A", Type: plugin.FlagTypeUint64},
						{Name: flagChainBTrustThresholdDenominator, DefaultValue: "3", Usage: "trusting threshold denominator of the chain B", Type: plugin.FlagTypeUint64},
						{Name: flagChainAMemoPrefix, Usage: "memo prefix of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBMemoPrefix, Usage: "memo prefix of the chain B", Type: plugin.FlagTypeString},
						{Name: flagChainAFaucet, Usage: "faucet URL of the chain A", Type: plugin.FlagTypeString},
						{Name: flagChainBFaucet, Usage: "faucet URL of the chain B", Type: plugin.FlagTypeString},
						{Name: flagTelemetryEnabled, DefaultValue: "false", Usage: "enable hermes telemetry", Type: plugin.FlagTypeBool},
						{Name: flagTelemetryHost, DefaultValue: "127.0.0.1", Usage: "hermes telemetry host", Type: plugin.FlagTypeString},
						{Name: flagTelemetryPort, DefaultValue: "3001", Usage: "hermes telemetry port", Type: plugin.FlagTypeUint64},
						{Name: flagRestEnabled, DefaultValue: "false", Usage: "enable hermes rest", Type: plugin.FlagTypeBool},
						{Name: flagRestHost, DefaultValue: "127.0.0.1", Usage: "hermes rest host", Type: plugin.FlagTypeString},
						{Name: flagRestPort, DefaultValue: "3000", Usage: "hermes rest port", Type: plugin.FlagTypeUint64},
						{Name: flagModeChannelsEnabled, DefaultValue: "true", Usage: "enable hermes channels", Type: plugin.FlagTypeBool},
						{Name: flagModeClientsEnabled, DefaultValue: "true", Usage: "enable hermes clients", Type: plugin.FlagTypeBool},
						{Name: flagModeClientsMisbehaviour, DefaultValue: "true", Usage: "enable hermes clients misbehaviour", Type: plugin.FlagTypeBool},
						{Name: flagModeClientsRefresh, DefaultValue: "true", Usage: "enable hermes client refresh time", Type: plugin.FlagTypeBool},
						{Name: flagModeConnectionsEnabled, DefaultValue: "true", Usage: "enable hermes connections", Type: plugin.FlagTypeBool},
						{Name: flagModePacketsEnabled, DefaultValue: "true", Usage: "enable hermes packets", Type: plugin.FlagTypeBool},
						{Name: flagModePacketsClearInterval, DefaultValue: "100", Usage: "hermes packet clear interval", Type: plugin.FlagTypeUint64},
						{Name: flagModePacketsClearOnStart, DefaultValue: "true", Usage: "enable hermes packets clear on start", Type: plugin.FlagTypeBool},
						{Name: flagModePacketsTxConfirmation, DefaultValue: "true", Usage: "hermes packet transaction confirmation", Type: plugin.FlagTypeBool},
						{Name: flagAutoRegisterCounterpartyPayee, DefaultValue: "false", Usage: "auto register the counterparty payee on a destination chain to the relayer's address on the source chain", Type: plugin.FlagTypeBool},
						{Name: flagGenerateWallets, DefaultValue: "false", Usage: "automatically generate wallets if they do not exist", Type: plugin.FlagTypeBool},
						{Name: flagOverwriteConfig, DefaultValue: "false", Usage: "overwrite the current config if it already exists", Type: plugin.FlagTypeBool},
					},
				},
			},
		},
	}
}

func getConfig(flags *pflag.FlagSet) string {
	config, _ := flags.GetString(flagConfig)
	return config
}
