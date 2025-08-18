package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	tendermintContract "operator/bindings/SP1ICS07Tendermint"
)

const (
	flagOnlyOnce       = "only-once"
	flagPrivateCluster = "private-cluster"
	flagProofType      = "proof-type"
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
	cmd := &cobra.Command{}
	cmd.Flags().String(flagProofType, "groth16", "the type of proof to use (groth16, plonk)")
	return cmd
}

func Fixtures(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{}
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
