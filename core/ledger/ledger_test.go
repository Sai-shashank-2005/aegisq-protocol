package ledger

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func setupLedger(t *testing.T) (*Ledger, *identity.NodeIdentity, crypto.Signer) {

	signer := &crypto.Ed25519Signer{}

	node, err := identity.NewNodeIdentity("validator-1", signer)
	if err != nil {
		t.Fatal(err)
	}

	vs := consensus.NewValidatorSet()
	vs.AddValidator("validator-1", node.PublicKey)

	// Genesis block must contain at least one transaction
	tx := createDummyTransaction(t, node)

	genesis := block.NewBlock(0, []byte("genesis"), []*transaction.Transaction{tx})
	if err := genesis.Finalize(node); err != nil {
		t.Fatal(err)
	}

	ledger := NewLedger(genesis, vs)

	return ledger, node, signer
}

func createDummyTransaction(t *testing.T, node *identity.NodeIdentity) *transaction.Transaction {

	tx := transaction.NewTransaction(
		node,
		"dummy_payload",
		"test_data",
	)

	err := tx.SignWithIdentity(node)
	if err != nil {
		t.Fatal(err)
	}

	return tx
}

func TestLedgerAddBlock(t *testing.T) {

	ledger, node, signer := setupLedger(t)

	tx := createDummyTransaction(t, node)

	newBlock := block.NewBlock(
		1,
		ledger.GetLastBlock().Hash,
		[]*transaction.Transaction{tx},
	)

	if err := newBlock.Finalize(node); err != nil {
		t.Fatal(err)
	}

	if err := ledger.AddBlock(newBlock, signer, node.PublicKey); err != nil {
		t.Fatal("Failed to add valid block:", err)
	}
}

func TestLedgerRejectWrongIndex(t *testing.T) {

	ledger, node, signer := setupLedger(t)

	tx := createDummyTransaction(t, node)

	newBlock := block.NewBlock(
		2, // wrong index
		ledger.GetLastBlock().Hash,
		[]*transaction.Transaction{tx},
	)

	if err := newBlock.Finalize(node); err != nil {
		t.Fatal(err)
	}

	if err := ledger.AddBlock(newBlock, signer, node.PublicKey); err == nil {
		t.Fatal("Block with wrong index should fail")
	}
}

func TestLedgerRejectWrongPreviousHash(t *testing.T) {

	ledger, node, signer := setupLedger(t)

	tx := createDummyTransaction(t, node)

	newBlock := block.NewBlock(
		1,
		[]byte("wrong_hash"),
		[]*transaction.Transaction{tx},
	)

	if err := newBlock.Finalize(node); err != nil {
		t.Fatal(err)
	}

	if err := ledger.AddBlock(newBlock, signer, node.PublicKey); err == nil {
		t.Fatal("Block with wrong previous hash should fail")
	}
}

func TestLedgerValidateChain(t *testing.T) {

	ledger, node, signer := setupLedger(t)

	tx := createDummyTransaction(t, node)

	newBlock := block.NewBlock(
		1,
		ledger.GetLastBlock().Hash,
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