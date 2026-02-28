package integration

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/ledger"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/scheduler"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func TestFullConsensusFlow(t *testing.T) {

	// --- Create 4 validators ---
	signer := &crypto.Ed25519Signer{}

	validators := make(map[string]*identity.NodeIdentity)
	vs := consensus.NewValidatorSet()

	for i := 1; i <= 4; i++ {
		id := "v" + string(rune('0'+i))
		node, err := identity.NewNodeIdentity(id, signer)
		if err != nil {
			t.Fatal(err)
		}
		validators[id] = node
		vs.AddValidator(id, node.PublicKey)
	}

	// --- Scheduler ---
	s := scheduler.NewRoundRobinScheduler(vs)

	// --- BFT ---
	votePool := consensus.NewVotePool(vs)
	finality := consensus.NewFinalityEngine(votePool)

	// --- Genesis ---
	genTx := transaction.NewTransaction(validators["v1"], "genesis", "init")
	genTx.SignWithIdentity(validators["v1"])

	genesis := block.NewBlock(
		0,
		0,
		[]byte("genesis"),
		[]*transaction.Transaction{genTx},
	)

	if err := genesis.Finalize(validators["v1"]); err != nil {
		t.Fatal(err)
	}

	ldg := ledger.NewLedger(genesis, vs)

	// --- Round 1 ---
	height := 1
	view := 0

	leaderID, err := s.GetLeader(height, view)
	if err != nil {
		t.Fatal(err)
	}

	leader := validators[leaderID]

	// Leader proposes block
	tx := transaction.NewTransaction(leader, "data", "payload")
	tx.SignWithIdentity(leader)

	newBlock := block.NewBlock(
		height,
		view,
		genesis.Hash,
		[]*transaction.Transaction{tx},
	)

	if err := newBlock.Finalize(leader); err != nil {
		t.Fatal(err)
	}

	blockHash := string(newBlock.Hash)

	// --- PREPARE ---
	for id := range validators {
		votePool.AddVote(consensus.Vote{
			ValidatorID: id,
			BlockHash:   blockHash,
			View:        view,
			Type:        consensus.Prepare,
		})
	}

	if !finality.TryPrepare(height, blockHash, view) {
		t.Fatal("Prepare quorum not reached")
	}

	// --- COMMIT ---
	for id := range validators {
		votePool.AddVote(consensus.Vote{
			ValidatorID: id,
			BlockHash:   blockHash,
			View:        view,
			Type:        consensus.Commit,
		})
	}

	if err := finality.TryCommit(height, blockHash, view); err != nil {
		t.Fatal("Commit failed:", err)
	}

	if !finality.IsFinalized(height, blockHash) {
		t.Fatal("Block not finalized")
	}

	// Append after finality
	if err := ldg.AddBlock(newBlock, signer, leader.PublicKey); err != nil {
		t.Fatal(err)
	}

	if err := ldg.ValidateChain(signer); err != nil {
		t.Fatal("Chain invalid after consensus:", err)
	}
}