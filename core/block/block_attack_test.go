package block

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func TestAttack_ModifyTransactionInsideBlock(t *testing.T) {
	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := transaction.NewTransaction(node, "QmCID", "Original")
	tx.SignWithIdentity(node)

	block := NewBlock(1, []byte("prevhash"), []*transaction.Transaction{tx})
	block.Finalize(node)

	block.Transactions[0].Metadata = "HACKED"

	valid, _ := block.Verify(signer, node.PublicKey)

	if valid {
		t.Fatal("Attack succeeded: block should be invalid after tx mutation")
	}
}

func TestAttack_SignatureReplay(t *testing.T) {
	signer := &crypto.Ed25519Signer{}

	node1, _ := identity.NewNodeIdentity("validator-1", signer)
	node2, _ := identity.NewNodeIdentity("validator-2", signer)

	tx := transaction.NewTransaction(node1, "QmCID", "Test")
	tx.SignWithIdentity(node1)

	valid, _ := tx.Verify(signer)

	if !valid {
		t.Fatal("Transaction should verify with correct key")
	}

	// Now try verifying with wrong public key manually
	ok := signer.Verify(node2.PublicKey, []byte("wronghash"), tx.Signature)
	if ok {
		t.Fatal("Replay attack succeeded")
	}
}

func TestAttack_ModifyBlockHash(t *testing.T) {
	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := transaction.NewTransaction(node, "QmCID", "Test")
	tx.SignWithIdentity(node)

	block := NewBlock(1, []byte("prevhash"), []*transaction.Transaction{tx})
	block.Finalize(node)

	block.Hash = []byte("fakehash")

	valid, _ := block.Verify(signer, node.PublicKey)

	if valid {
		t.Fatal("Attack succeeded: forged hash accepted")
	}
}