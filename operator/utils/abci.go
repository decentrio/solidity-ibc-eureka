package utils

import (
	"bytes"
	"context"
	"fmt"

	tendermintContract "operator/bindings/SP1ICS07Tendermint"

	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/gogoproto/proto"
	ics23 "github.com/cosmos/ics23/go"
)

func ProvePath(client *http.HTTP, path [][]byte, targetHeight uint64) ([]byte, *tendermintContract.IMembershipMsgsMerkleProof, error) {
	abciResp, err := client.ABCIQueryWithOptions(context.Background(), fmt.Sprintf("store/%s/key", string(path[0])), bytes.Join(path[1:], nil), rpcclient.ABCIQueryOptions{
		Height: 0,
		Prove:  true,
	})
	if err != nil {
		return nil, nil, err
	}
	if abciResp.Response.Height+1 != int64(targetHeight) {
		return nil, nil, fmt.Errorf("invalid proof height, expected %v, got %v", targetHeight-1, abciResp.Response.Height)
	}

	if !bytes.Equal(abciResp.Response.Key, path[1]) {
		return nil, nil, fmt.Errorf("invalid proof hkey mismatch, expected %v, got %v", abciResp.Response.Key, path[1])
	}

	merkleProof, err := convertTmToIcsMerkleProof(*abciResp.Response.ProofOps)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert to merkle proof, %s", err.Error())
	}

	if len(merkleProof.Proofs) == 0 {
		return nil, nil, fmt.Errorf("empty merkle proof")
	}

	return abciResp.Response.Value, merkleProof, err
}

func convertTmToIcsMerkleProof(proofOps crypto.ProofOps) (*tendermintContract.IMembershipMsgsMerkleProof, error) {
	merkleProofs := make([]tendermintContract.IMembershipMsgsCommitmentProof, proofOps.Size())

	for _, p := range proofOps.Ops {
		var commitmentProof ics23.CommitmentProof
		err := proto.Unmarshal(p.Data, &commitmentProof)
		if err != nil {
			return nil, err
		}

		existProof := commitmentProof.GetExist()
		if existProof == nil {
			return nil, fmt.Errorf("invalid proof, cannot find existence proof")
		}

		path := make([]tendermintContract.IMembershipMsgsInnerOp, len(existProof.Path))
		for _, i := range existProof.Path {
			path = append(path, tendermintContract.IMembershipMsgsInnerOp{
				HashOp: uint8(i.Hash),
				Prefix: i.Prefix,
				Suffix: i.Suffix,
			})
		}

		merkleProofs = append(merkleProofs, tendermintContract.IMembershipMsgsCommitmentProof{
			ExistenceProof: tendermintContract.IMembershipMsgsExistenceProof{
				Key:   existProof.Key,
				Value: [32]byte(existProof.Value),
				Leaf: tendermintContract.IMembershipMsgsLeafOp{
					HashOp:       uint8(existProof.Leaf.Hash),
					PrehashKey:   uint8(existProof.Leaf.PrehashKey),
					PrehashValue: uint8(existProof.Leaf.PrehashValue),
					Prefix:       existProof.Leaf.Prefix,
				},
				Path: path,
			},
		})
	}

	return &tendermintContract.IMembershipMsgsMerkleProof{
		Proofs: merkleProofs,
	}, nil
}

func BytesToBytes32(data []byte) [32]byte {
	var result [32]byte
	copy(result[:], data)
	return result
}
