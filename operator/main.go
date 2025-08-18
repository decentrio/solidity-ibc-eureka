package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	tendermintContract "operator/bindings/SP1ICS07Tendermint"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

const (
	flagOnlyOnce       = "only-once"
	flagPrivateCluster = "private-cluster"
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

			return nil
		},
	}
	cmd.Flags().Bool(flagOnlyOnce, false, "run only once")
	cmd.Flags().Bool(flagPrivateCluster, false, "enable private cluster")
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

			return nil
		},
	}
	cmd.Flags().String(flagProofType, "groth16", "the type of proof to use (groth16, plonk)")
	cmd.Flags().Int64(flagTrustedBlock, 0, "the trusted block height, if <height> is 0 then catch latest block")
	cmd.Flags().String(flagOutput, "json", "the output structure for the genesis state (json, file)")
	cmd.Flags().String(flagOutputPath, "./data/genesis.json", "the path to the output file for the genesis state")
	cmd.Flags().String(flagTrustLevel, "2/3", "the trust level for the genesis state (e.g., 2/3)")
	cmd.Flags().String(flagTrustingPeriod, "24", "the trusting period for the genesis state")
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
			return nil
		},
	}
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
func parseTrustThreshold(value string) (tendermintContract.IICS07TendermintMsgsTrustThreshold, error) {
	parts := strings.Split(value, "/")
	if len(parts) != 2 {
		return tendermintContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid trust threshold format: %s (expected format: 'numerator/denominator')", value)
	}

	numerator, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return tendermintContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid numerator: %s", parts[0])
	}

	denominator, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return tendermintContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid denominator: %s", parts[1])
	}

	if denominator == 0 {
		return tendermintContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("denominator cannot be zero")
	}

	return tendermintContract.IICS07TendermintMsgsTrustThreshold{
		Numerator:   uint8(numerator),
		Denominator: uint8(denominator),
	}, nil
}
