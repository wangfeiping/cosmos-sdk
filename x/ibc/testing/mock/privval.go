package mock

import (
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

var _ tmtypes.PrivValidator = PV{}

// MockPV implements PrivValidator without any safety or persistence.
// Only use it for testing.
type PV struct {
	PrivKey crypto.PrivKey
}

func NewPV() PV {
	return PV{ed25519.GenPrivKey()}
}

// GetPubKey implements PrivValidator interface
func (pv PV) GetPubKey() (crypto.PubKey, error) {
	return pv.PrivKey.PubKey(), nil
}

// SignVote implements PrivValidator interface
func (pv PV) SignVote(chainID string, vote *tmproto.Vote) error {
	signBytes := tmtypes.VoteSignBytes(chainID, vote)
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	vote.Signature = sig
	return nil
}

// SignProposal implements PrivValidator interface
func (pv PV) SignProposal(chainID string, proposal *tmproto.Proposal) error {
	signBytes := tmtypes.ProposalSignBytes(chainID, proposal)
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	proposal.Signature = sig
	return nil
}
