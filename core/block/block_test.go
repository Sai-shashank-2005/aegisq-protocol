package block

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func createTestTx(t *testing.T, node *identity.NodeIdentity) *transaction.Transaction {
	tx := transaction.NewTransaction(node, "payload", "data")
	if err := tx.SignWithIdentity(node); err != nil {
		t.Fatal(err)
	}
	return tx
}

func TestBlockFinalizeAndVerify(t *testing.T) {

	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := createTestTx(t, node)

	block := NewBlock(
		1,   // height
		0,   // view
		[]byte("prev_hash"),
		[]*transaction.Transaction{tx},
	)

	if err := block.Finalize(node); err != nil {
		t.Fatal(err)
	}

	valid, err := block.Verify(signer, node.PublicKey)
	if err != nil || !valid {
		t.Fatal("Block verification failed")
	}
}

func TestBlockTamperFails(t *testing.T) {

	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := createTestTx(t, node)

	block := NewBlock(
		1,
		0,
		[]byte("prev_hash"),
		[]*transaction.Transaction{tx},
	)

	if err := block.Finalize(node); err != nil {
		t.Fatal(err)
	}

	// Tamper block
	block.Hash = []byte("fake_hash")

	valid, _ := block.Verify(signer, node.PublicKey)
	if valid {
		t.Fatal("Tampered block should fail")
	}
}