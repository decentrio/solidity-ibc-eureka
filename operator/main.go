package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	updateClientContract "operator/bindings/UpdateClient"
	tendermintClient "operator/client"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
)

const (
	flagOnlyOnce       = "only-once"
	flagProofType      = "proof-type"
	flagOutput         = "output"
	flagOutputPath     = "output-path"
	flagTrustLevel     = "trust-level"
	flagTrustingPeriod = "trusting-period"
	flagTrustedBlock   = "trusted-block"
)

func main() {
	zLogger, _ := zap.NewProduction(zap.AddStacktrace(zap.DPanicLevel))
	defer zLogger.Sync()

	logger := zLogger.Sugar()

	rootCmd := &cobra.Command{
		Use:   "operator [command]",
		Short: "command for operator",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(
		Start(zLogger),
		Genesis(zLogger),
		Fixtures(zLogger),
	)

	err := rootCmd.Execute()
	if err != nil {
		logger.Fatal(err)
	}
}

func Start(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO:
			return nil
		},
	}
	cmd.Flags().Bool(flagOnlyOnce, false, "run only once")
	// cmd.Flags().String(flagConfigPath, ".operator/config.yaml", "the path to your operator priv and pub key")
	return cmd
}

func Genesis(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genesis",
		Short: "genesis",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := godotenv.Load()
			if err != nil {
				return fmt.Errorf("error loading .env file: %v", err)
			}

			// Read RPC endpoint from environment variable
			rpcEndpoint := os.Getenv("TENDERMINT_RPC_URL")
			if rpcEndpoint == "" {
				return fmt.Errorf("TENDERMINT_RPC_URL environment variable is required in .env file")
			}
			client, err := rpchttp.New(rpcEndpoint, "/websocket")
			if err != nil {
				return fmt.Errorf("failed to create RPC client: %w", err)
			}

			trustedBlock, err := cmd.Flags().GetInt64(flagTrustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get trusted block: %w", err)
			}

			status, err := client.Status(context.Background())
			if err != nil {
				return fmt.Errorf("failed to get status: %w", err)
			}

			if trustedBlock == 0 {
				// get latest block
				trustedBlock = status.SyncInfo.LatestBlockHeight
			}

			trustedLightBlock, err := tendermintClient.GetLightBlock(client, trustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get light block: %w", err)
			}
			_ = trustedLightBlock

			unbondingPeriod, err := tendermintClient.GetUnbondingTime(client)
			if err != nil {
				return fmt.Errorf("failed to get unbonding time: %w", err)
			}

			trustingPeriod, err := cmd.Flags().GetUint32(flagTrustingPeriod)
			if err != nil {
				return fmt.Errorf("failed to get trusting period: %w", err)
			}

			if trustingPeriod == 0 {
				trustingPeriod = uint32(unbondingPeriod * 2 / 3)
			}

			if trustingPeriod > uint32(unbondingPeriod) {
				return fmt.Errorf("trusting period %d cannot be greater than unbonding period %d", trustingPeriod, uint32(unbondingPeriod))
			}

			chainId := trustedLightBlock.SignedHeader.Header.ChainID
			revision := clienttypes.ParseChainID(chainId)
			trustLevel, err := cmd.Flags().GetString(flagTrustLevel)
			if err != nil {
				return fmt.Errorf("failed to get trust level from flag: %w", err)
			}
			trustThreshold, err := parseTrustThreshold(trustLevel)
			if err != nil {
				return fmt.Errorf("failed to parse trust level: %w", err)
			}

			proofType, err := cmd.Flags().GetString(flagProofType)
			if err != nil {
				return fmt.Errorf("failed to get proof type from flag: %w", err)
			}

			var zkAlgorithm tendermintClient.SupportedZkAlgorithm
			switch proofType {
			case "groth16":
				zkAlgorithm = tendermintClient.Groth16
			case "plonk":
				zkAlgorithm = tendermintClient.Plonk
			default:
				return fmt.Errorf("unsupported proof type: %s, supported types are: groth16, plonk", proofType)
			}

			clientState := updateClientContract.IICS07TendermintMsgsClientState{
				ChainId:    chainId,
				TrustLevel: trustThreshold,
				LatestHeight: updateClientContract.IICS02ClientMsgsHeight{
					RevisionNumber: revision,
					RevisionHeight: uint64(trustedLightBlock.SignedHeader.Header.Height),
				},
				IsFrozen:        false,
				ZkAlgorithm:     uint8(zkAlgorithm),
				TrustingPeriod:  trustingPeriod,
				UnbondingPeriod: uint32(unbondingPeriod),
			}

			consensusState := updateClientContract.IICS07TendermintMsgsConsensusState{
				Timestamp:          big.NewInt(trustedLightBlock.SignedHeader.Header.Time.UnixMilli()),
				Root:               bytesToBytes32(trustedLightBlock.SignedHeader.Header.AppHash),
				NextValidatorsHash: bytesToBytes32(trustedLightBlock.SignedHeader.NextValidatorsHash),
			}

			genesis := tendermintClient.SP1ICS07TendermintGenesis{
				TrustedClientState:    clientState,
				TrustedConsensusState: consensusState,
			}

			outputType, err := cmd.Flags().GetString(flagOutput)
			if err != nil {
				return fmt.Errorf("failed to get output type from flag: %w", err)
			}

			data, err := json.Marshal(genesis)
			if err != nil {
				return fmt.Errorf("failed to marshal genesis state: %w", err)
			}

			switch outputType {
			case "json":
				fmt.Println(string(data))
			case "file":
				outputDir, err := cmd.Flags().GetString(flagOutputPath)
				if err != nil {
					return fmt.Errorf("failed to get output path from flag: %w", err)
				}

				if err := os.WriteFile(outputDir, data, 0644); err != nil {
					return fmt.Errorf("failed to write genesis state to file: %w", err)
				}
			default:
				return fmt.Errorf("unsupported output type: %s, supported types are: json, file", outputType)
			}

			return nil
		},
	}
	cmd.Flags().String(flagProofType, "groth16", "the type of proof to use (groth16, plonk)")
	cmd.Flags().Int64(flagTrustedBlock, 0, "the trusted block height, if <height> is 0 then catch latest block")
	cmd.Flags().String(flagOutput, "json", "the output structure for the genesis state (json, file)")
	cmd.Flags().String(flagOutputPath, "./data/genesis.json", "the path to the output file for the genesis state")
	cmd.Flags().String(flagTrustLevel, "2/3", "the trust level for the genesis state (e.g., 2/3)")
	cmd.Flags().Uint32(flagTrustingPeriod, 0, "the trusting period for the genesis state")
	return cmd
}

func Fixtures(logger *zap.Logger) *cobra.Command {
	fixturesCmd := &cobra.Command{
		Use:   "fixtures",
		Short: "fixtures",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	fixturesCmd.AddCommand(
		UpdateClientCmd(logger),
		MembershipCmd(logger),
		Misbehaviour(logger),
	)

	return fixturesCmd
}

// UpdateClientCmd creates a command to update the client
func UpdateClientCmd(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-client",
		Short: "update client",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

// MembershipCmd creates a command to manage membership
func MembershipCmd(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "membership",
		Short: "membership",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := godotenv.Load()
			if err != nil {
				return fmt.Errorf("error loading .env file: %v", err)
			}

			// Read RPC endpoint from environment variable
			rpcEndpoint := os.Getenv("TENDERMINT_RPC_URL")
			if rpcEndpoint == "" {
				return fmt.Errorf("TENDERMINT_RPC_URL environment variable is required in .env file")
			}
			client, err := rpchttp.New(rpcEndpoint, "/websocket")
			if err != nil {
				return fmt.Errorf("failed to create RPC client: %w", err)
			}

			trustedBlock, err := cmd.Flags().GetInt64(flagTrustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get trusted block: %w", err)
			}

			status, err := client.Status(context.Background())
			if err != nil {
				return fmt.Errorf("failed to get status: %w", err)
			}

			if trustedBlock == 0 {
				// get latest block
				trustedBlock = status.SyncInfo.LatestBlockHeight
			}

			trustedLightBlock, err := tendermintClient.GetLightBlock(client, trustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get light block: %w", err)
			}
			_ = trustedLightBlock

			unbondingPeriod, err := tendermintClient.GetUnbondingTime(client)
			if err != nil {
				return fmt.Errorf("failed to get unbonding time: %w", err)
			}

			trustingPeriod, err := cmd.Flags().GetUint32(flagTrustingPeriod)
			if err != nil {
				return fmt.Errorf("failed to get trusting period: %w", err)
			}

			if trustingPeriod == 0 {
				trustingPeriod = uint32(unbondingPeriod * 2 / 3)
			}

			if trustingPeriod > uint32(unbondingPeriod) {
				return fmt.Errorf("trusting period %d cannot be greater than unbonding period %d", trustingPeriod, uint32(unbondingPeriod))
			}

			chainId := trustedLightBlock.SignedHeader.Header.ChainID
			revision := clienttypes.ParseChainID(chainId)
			trustLevel, err := cmd.Flags().GetString(flagTrustLevel)
			if err != nil {
				return fmt.Errorf("failed to get trust level from flag: %w", err)
			}
			trustThreshold, err := parseTrustThreshold(trustLevel)
			if err != nil {
				return fmt.Errorf("failed to parse trust level: %w", err)
			}

			proofType, err := cmd.Flags().GetString(flagProofType)
			if err != nil {
				return fmt.Errorf("failed to get proof type from flag: %w", err)
			}

			var zkAlgorithm tendermintClient.SupportedZkAlgorithm
			switch proofType {
			case "groth16":
				zkAlgorithm = tendermintClient.Groth16
			case "plonk":
				zkAlgorithm = tendermintClient.Plonk
			default:
				return fmt.Errorf("unsupported proof type: %s, supported types are: groth16, plonk", proofType)
			}

			clientState := updateClientContract.IICS07TendermintMsgsClientState{
				ChainId:    chainId,
				TrustLevel: trustThreshold,
				LatestHeight: updateClientContract.IICS02ClientMsgsHeight{
					RevisionNumber: revision,
					RevisionHeight: uint64(trustedLightBlock.SignedHeader.Header.Height),
				},
				IsFrozen:        false,
				ZkAlgorithm:     uint8(zkAlgorithm),
				TrustingPeriod:  trustingPeriod,
				UnbondingPeriod: uint32(unbondingPeriod),
			}

			consensusState := updateClientContract.IICS07TendermintMsgsConsensusState{
				Timestamp:          big.NewInt(trustedLightBlock.SignedHeader.Header.Time.UnixMilli()),
				Root:               bytesToBytes32(trustedLightBlock.SignedHeader.Header.AppHash),
				NextValidatorsHash: bytesToBytes32(trustedLightBlock.SignedHeader.NextValidatorsHash),
			}

			genesis := tendermintClient.SP1ICS07TendermintGenesis{
				TrustedClientState:    clientState,
				TrustedConsensusState: consensusState,
			}

			return nil
		},
	}
	cmd.Flags().String(flagProofType, "groth16", "the type of proof to use (groth16, plonk)")
	cmd.Flags().Int64(flagTrustedBlock, 0, "the trusted block height, if <height> is 0 then catch latest block")
	cmd.Flags().String(flagOutput, "json", "the output structure for the genesis state (json, file)")
	cmd.Flags().String(flagOutputPath, "./data/genesis.json", "the path to the output file for the genesis state")
	cmd.Flags().String(flagTrustLevel, "2/3", "the trust level for the genesis state (e.g., 2/3)")
	cmd.Flags().Uint32(flagTrustingPeriod, 0, "the trusting period for the genesis state")
	return cmd
}

func Misbehaviour(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "misbehaviour",
		Short: "misbehaviour",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

// parseTrustThreshold parses a trust threshold fraction string like "2/3"
func parseTrustThreshold(value string) (updateClientContract.IICS07TendermintMsgsTrustThreshold, error) {
	parts := strings.Split(value, "/")
	if len(parts) != 2 {
		return updateClientContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid trust threshold format: %s (expected format: 'numerator/denominator')", value)
	}

	numerator, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return updateClientContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid numerator: %s", parts[0])
	}

	denominator, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return updateClientContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid denominator: %s", parts[1])
	}

	if denominator == 0 {
		return updateClientContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("denominator cannot be zero")
	}

	return updateClientContract.IICS07TendermintMsgsTrustThreshold{
		Numerator:   uint8(numerator),
		Denominator: uint8(denominator),
	}, nil
}

func bytesToBytes32(data []byte) [32]byte {
	var result [32]byte
	copy(result[:], data)
	return result
}
