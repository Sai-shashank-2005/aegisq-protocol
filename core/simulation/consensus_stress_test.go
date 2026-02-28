package simulation

import (
	"testing"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/ledger"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/scheduler"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func TestConsensusWith1000SyntheticTransactions(t *testing.T) {

	signer := &crypto.Ed25519Signer{}

	// -------------------------
	// Create 4 validators
	// -------------------------
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

	// -------------------------
	// Genesis Block
	// -------------------------
	genTx, err := GenerateSyntheticTransaction(validators["v1"])
	if err != nil {
		t.Fatal(err)
	}

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

	// -------------------------
	// Generate 1000 synthetic transactions
	// -------------------------
	txs, err := GenerateBulkTransactions(validators["v1"], 50000)
	if err != nil {
		t.Fatal(err)
	}

	// -------------------------
	// Determine Leader (Layer 7)
	// -------------------------
	height := 1
	view := 0

	s := scheduler.NewRoundRobinScheduler(vs)

	leaderID, err := s.GetLeader(height, view)
	if err != nil {
		t.Fatal(err)
	}

	leader := validators[leaderID]

	// -------------------------
	// Build Block (by scheduled leader only)
	// -------------------------
	start := time.Now()

	newBlock := block.NewBlock(
		height,
		view,
		genesis.Hash,
		txs,
	)

	if err := newBlock.Finalize(leader); err != nil {
		t.Fatal(err)
	}

	elapsed := time.Since(start)
	t.Log("Block finalize time:", elapsed)

	// -------------------------
	// BFT Prepare + Commit
	// -------------------------
	votePool := consensus.NewVotePool(vs)
	finality := consensus.NewFinalityEngine(votePool)

	blockHash := string(newBlock.Hash)

	// PREPARE
	for id := range validators {
		votePool.AddVote(consensus.Vote{
			ValidatorID: id,
			BlockHash:   blockHash,
			View:        view,
			Type:        consensus.Prepare,
		})
	}

	if !finality.TryPrepare(height, blockHash, view) {
		t.Fatal("Prepare quorum failed")
	}

	// COMMIT
	for id := range validators {
		votePool.AddVote(consensus.Vote{
			ValidatorID: id,
			BlockHash:   blockHash,
			View:        view,
			Type:        consensus.Commit,
		})
	}

	if err := finality.TryCommit(height, blockHash, view); err != nil {
		t.Fatal(err)
	}

	if !finality.IsFinalized(height, blockHash) {
		t.Fatal("Finality failed")
	}

	// -------------------------
	// Append to ledger
	// -------------------------
	if err := ldg.AddBlock(newBlock, signer, leader.PublicKey); err != nil {
		t.Fatal(err)
	}

	if err := ldg.ValidateChain(signer); err != nil {
		t.Fatal("Ledger validation failed")
	}

	t.Log("Consensus completed successfully with 50000 synthetic transactions")
}