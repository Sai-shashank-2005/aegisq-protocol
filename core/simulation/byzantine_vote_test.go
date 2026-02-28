package simulation

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
)

func TestByzantineEquivocationAttack(t *testing.T) {

	vs := consensus.NewValidatorSet()

	// 4 validators (f = 1)
	vs.AddValidator("v1", []byte("k1"))
	vs.AddValidator("v2", []byte("k2"))
	vs.AddValidator("v3", []byte("k3")) // Byzantine
	vs.AddValidator("v4", []byte("k4"))

	vp := consensus.NewVotePool(vs)

	view := 1

	hashA := "blockA"
	hashB := "blockB"

	// Honest votes for blockA
	vp.AddVote(consensus.Vote{
		ValidatorID: "v1",
		BlockHash:   hashA,
		View:        view,
		Type:        consensus.Prepare,
	})

	vp.AddVote(consensus.Vote{
		ValidatorID: "v2",
		BlockHash:   hashA,
		View:        view,
		Type:        consensus.Prepare,
	})

	// Byzantine equivocation
	vp.AddVote(consensus.Vote{
		ValidatorID: "v3",
		BlockHash:   hashA,
		View:        view,
		Type:        consensus.Prepare,
	})

	vp.AddVote(consensus.Vote{
		ValidatorID: "v3",
		BlockHash:   hashB, // conflicting block
		View:        view,
		Type:        consensus.Prepare,
	})

	// Honest v4 votes blockA
	vp.AddVote(consensus.Vote{
		ValidatorID: "v4",
		BlockHash:   hashA,
		View:        view,
		Type:        consensus.Prepare,
	})

	// Check quorum for blockA
	if !vp.HasQuorum(hashA, view, consensus.Prepare) {
		t.Fatal("Expected quorum for blockA not reached")
	}

	// Now check if blockB incorrectly forms quorum
	if vp.HasQuorum(hashB, view, consensus.Prepare) {
		t.Fatal("Byzantine equivocation formed illegal quorum for blockB")
	}
}