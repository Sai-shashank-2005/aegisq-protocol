package ledger

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func TestFullChainWithDilithium(t *testing.T) {

	signer, err := crypto.NewDilithiumSigner()
	if err != nil {
		t.Fatal(err)
	}
	defer signer.Close()

	// Create validator identity
	node, err := identity.NewNodeIdentity("validator-1", signer)
	if err != nil {
		t.Fatal(err)
	}

	// Create validator set
	vs := consensus.NewValidatorSet()
	vs.AddValidator("validator-1", node.PublicKey)

	// --- Genesis block must contain a transaction ---
	genesisTx := transaction.NewTransaction(
		node,
		"genesis_payload",
		"genesis_data",
	)

	if err := genesisTx.SignWithIdentity(node); err != nil {
		t.Fatal(err)
	}

	genesis := block.NewBlock(
		0,
		[]byte("genesis"),
		[]*transaction.Transaction{genesisTx},
	)

	if err := genesis.Finalize(node); err != nil {
		t.Fatal(err)
	}

	ledger := NewLedger(genesis, vs)

	// --- Create next block with transaction ---
	tx := transaction.NewTransaction(
		node,
		"block_payload",
		"block_data",
	)

	if err := tx.SignWithIdentity(node); err != nil {
		t.Fatal(err)
	}

	newBlock := block.NewBlock(
		1,
		genesis.Hash,
		[]*transaction.Transaction{tx},
	)

	if err := newBlock.Finalize(node); err != nil {
		t.Fatal(err)
	}

	if err := ledger.AddBlock(newBlock, signer, node.PublicKey); err != nil {
		t.Fatal(err)
	}

	if err := ledger.ValidateChain(signer); err != nil {
		t.Fatal("Chain validation failed:", err)
	}
}