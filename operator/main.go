package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	tendermintContract "operator/bindings/SP1ICS07Tendermint"
	tendermintClient "operator/client"
	"operator/keys"
	"operator/runner"
)

const (
	flagOnlyOnce       = "only-once"
	flagProofType      = "proof-type"
	flagOutput         = "output"
	flagOutputPath     = "output-path"
	flagTrustLevel     = "trust-level"
	flagTrustingPeriod = "trusting-period"
	flagTrustedBlock   = "trusted-block"
	flagMembership     = "membership"
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

			trustedBlock, err := cmd.Flags().GetInt64(flagTrustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get trusted block: %w", err)
			}
			trustingPeriod, err := cmd.Flags().GetUint32(flagTrustingPeriod)
			if err != nil {
				return fmt.Errorf("failed to get trusting period: %w", err)
			}
			trustLevel, err := cmd.Flags().GetString(flagTrustLevel)
			if err != nil {
				return fmt.Errorf("failed to get trust level from flag: %w", err)
			}
			proofType, err := cmd.Flags().GetString(flagProofType)
			if err != nil {
				return fmt.Errorf("failed to get proof type from flag: %w", err)
			}

			genesis, err := tendermintClient.GetGenesis(trustedBlock, trustingPeriod, trustLevel, proofType)
			if err != nil {
				return fmt.Errorf("failed to get genesis: %w", err)
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
		Short: "membership <key_path> <is_base64> <membership_type>",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := godotenv.Load()
			if err != nil {
				return fmt.Errorf("error loading .env file: %v", err)
			}

			// Read RPC endpoint from environment variable
			tendermintRpcEndpoint := os.Getenv("TENDERMINT_RPC_URL")
			if tendermintRpcEndpoint == "" {
				return fmt.Errorf("TENDERMINT_RPC_URL environment variable is required in .env file")
			}
			tendermintRpcClient, err := rpchttp.New(tendermintRpcEndpoint, "/websocket")
			if err != nil {
				return fmt.Errorf("failed to create RPC client: %w", err)
			}
			ethRpcEndpoint := os.Getenv("ETH_RPC_URL")
			if ethRpcEndpoint == "" {
				return fmt.Errorf("ETH_RPC_URL environment variable is required in .env file")
			}

			hexAddress := os.Getenv("CONTRACT_ADDRESS")
			if hexAddress == "" {
				return fmt.Errorf("CONTRACT_ADDRESS environment variable is required in .env file")
			}

			privKey := os.Getenv("PRIVATE_KEY")
			if privKey == "" {
				return fmt.Errorf("PRIVATE_KEY environment variable is required in .env file")
			}
			privateKey, err := keys.RestoreKey(privKey)
			if err != nil {
				return fmt.Errorf("failed to restore private key: %w", err)
			}

			chainIdEth := os.Getenv("CHAIN_ID")
			if chainIdEth == "" {
				return fmt.Errorf("CHAIN_ID environment variable is required in .env file")
			}

			chainIdInt := big.NewInt(0)
			chainIdInt, ok := chainIdInt.SetString(chainIdEth, 10)
			if !ok {
				return fmt.Errorf("invalid chain id: %v", err)
			}

			ethClient, err := ethclient.Dial(ethRpcEndpoint)
			if err != nil {
				return fmt.Errorf("failed to create Ethereum client: %w", err)
			}

			trustedBlock, err := cmd.Flags().GetInt64(flagTrustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get trusted block: %w", err)
			}
			trustingPeriod, err := cmd.Flags().GetUint32(flagTrustingPeriod)
			if err != nil {
				return fmt.Errorf("failed to get trusting period: %w", err)
			}
			trustLevel, err := cmd.Flags().GetString(flagTrustLevel)
			if err != nil {
				return fmt.Errorf("failed to get trust level from flag: %w", err)
			}
			proofType, err := cmd.Flags().GetString(flagProofType)
			if err != nil {
				return fmt.Errorf("failed to get proof type from flag: %w", err)
			}
			genesis, err := tendermintClient.GetGenesis(trustedBlock, trustingPeriod, trustLevel, proofType)
			if err != nil {
				return fmt.Errorf("failed to get genesis: %w", err)
			}

			tendermintAddr := common.HexToAddress(hexAddress)
			ics07Tendermint, err := tendermintContract.NewContractSP1ICS07Tendermint(
				tendermintAddr,
				ethClient,
			)
			if err != nil {
				return fmt.Errorf("failed to create ICS07 Tendermint contract: %w", err)
			}

			isMembership, err := cmd.Flags().GetBool(flagMembership)
			if err != nil {
				return fmt.Errorf("failed to get membership flag: %w", err)
			}

			publicKey, err := keys.PublicKey(privateKey)
			if err != nil {
				log.Fatal(err)
			}

			fromAddress := crypto.PubkeyToAddress(*publicKey)
			nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}
			gasPrice, err := ethClient.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainIdInt)
			if err != nil {
				return fmt.Errorf("failed to create auth transactor: %w", err)
			}
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)     // in wei
			auth.GasLimit = uint64(300000) // in units
			auth.GasPrice = gasPrice

			kvPairs, proofs, err := runner.RunMembership(tendermintRpcClient, args[0], trustedBlock, args[1] == "true")
			if err != nil {
				return err
			}

			membershipType, err := strconv.Atoi(args[2])
			if err != nil {
				return fmt.Errorf("invalid membership type: %w", err)
			}
			if isMembership {
				msg := tendermintContract.ILightClientMsgsMsgVerifyMembership{
					Height:                tendermintContract.IICS02ClientMsgsHeight(genesis.TrustedClientState.LatestHeight),
					KvPairs:               kvPairs,
					MerkleProofs:          proofs,
					AppHash:               genesis.TrustedConsensusState.Root,
					TrustedConsensusState: tendermintContract.IICS07TendermintMsgsConsensusState(genesis.TrustedConsensusState),
					MembershipType:        uint8(membershipType),
				}

				tx, err := ics07Tendermint.VerifyMembership(auth, msg)
				if err != nil {
					return fmt.Errorf("failed to verify membership: %w", err)
				}
				fmt.Println("tx hash: ", tx.Hash().Hex())
			} else {
				msg := tendermintContract.ILightClientMsgsMsgVerifyNonMembership{
					Height:                tendermintContract.IICS02ClientMsgsHeight(genesis.TrustedClientState.LatestHeight),
					KvPairs:               kvPairs,
					MerkleProofs:          proofs,
					AppHash:               genesis.TrustedConsensusState.Root,
					TrustedConsensusState: tendermintContract.IICS07TendermintMsgsConsensusState(genesis.TrustedConsensusState),
					MembershipType:        uint8(membershipType),
				}

				tx, err := ics07Tendermint.VerifyNonMembership(auth, msg)
				if err != nil {
					return fmt.Errorf("failed to verify non-membership: %w", err)
				}
				fmt.Println("tx hash: ", tx.Hash().Hex())
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
	cmd.Flags().Bool(flagMembership, true, "verify membership/non-membership proof")
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
