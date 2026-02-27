package block

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func TestBlockFinalizeAndVerify(t *testing.T) {
	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := transaction.NewTransaction(node, "QmCID", "Test File")
	tx.SignWithIdentity(node)

	block := NewBlock(1, []byte("prevhash"), []*transaction.Transaction{tx})
	block.Finalize(node)

	valid, err := block.Verify(signer, node.PublicKey)
	if err != nil || !valid {
		t.Fatal("Block verification failed")
	}
}

func TestBlockTamperFails(t *testing.T) {
	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := transaction.NewTransaction(node, "QmCID", "Test File")
	tx.SignWithIdentity(node)

	block := NewBlock(1, []byte("prevhash"), []*transaction.Transaction{tx})
	block.Finalize(node)

	block.Transactions[0].Metadata = "HACKED"

	valid, _ := block.Verify(signer, node.PublicKey)

	if valid {
		t.Fatal("Tampered block should fail")
	}
}