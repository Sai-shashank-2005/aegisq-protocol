package ledger

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func createBlock(index int, prevHash []byte, node *identity.NodeIdentity) *block.Block {
	tx := transaction.NewTransaction(node, "data", "meta")
	tx.SignWithIdentity(node)

	b := block.NewBlock(index, prevHash, []*transaction.Transaction{tx})
	b.Finalize(node)

	return b
}

func TestLedgerAddBlock(t *testing.T) {

	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	genesis := createBlock(0, []byte("genesis"), node)

	ledger := NewLedger(genesis)

	block1 := createBlock(1, genesis.Hash, node)

	err := ledger.AddBlock(block1, signer, node.PublicKey)
	if err != nil {
		t.Fatal("Failed to add valid block")
	}
}

func TestLedgerRejectWrongIndex(t *testing.T) {

	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	genesis := createBlock(0, []byte("genesis"), node)
	ledger := NewLedger(genesis)

	blockWrong := createBlock(5, genesis.Hash, node)

	err := ledger.AddBlock(blockWrong, signer, node.PublicKey)
	if err == nil {
		t.Fatal("Ledger accepted block with wrong index")
	}
}

func TestLedgerRejectWrongPreviousHash(t *testing.T) {

	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	genesis := createBlock(0, []byte("genesis"), node)
	ledger := NewLedger(genesis)

	blockWrong := createBlock(1, []byte("fakehash"), node)

	err := ledger.AddBlock(blockWrong, signer, node.PublicKey)
	if err == nil {
		t.Fatal("Ledger accepted block with wrong previous hash")
	}
}

func TestLedgerValidateChain(t *testing.T) {

	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	genesis := createBlock(0, []byte("genesis"), node)
	ledger := NewLedger(genesis)

	block1 := createBlock(1, genesis.Hash, node)
	ledger.AddBlock(block1, signer, node.PublicKey)

	block2 := createBlock(2, block1.Hash, node)
	ledger.AddBlock(block2, signer, node.PublicKey)

	err := ledger.ValidateChain(signer, node.PublicKey)
	if err != nil {
		t.Fatal("Valid chain rejected")
	}
}
