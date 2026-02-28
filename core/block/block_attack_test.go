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

	tx := transaction.NewTransaction(node, "payload", "data")
	tx.SignWithIdentity(node)

	block := NewBlock(
		1,
		0,
		[]byte("prev_hash"),
		[]*transaction.Transaction{tx},
	)

	block.Finalize(node)

	// Corrupt transaction signature after block finalized
	block.Transactions[0].Signature = []byte("corrupted")

	valid, _ := block.Verify(signer, node.PublicKey)

	if valid {
		t.Fatal("Attack succeeded: block should be invalid after tx mutation")
	}
}

func TestAttack_ModifyBlockHash(t *testing.T) {

	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := transaction.NewTransaction(node, "payload", "data")
	tx.SignWithIdentity(node)

	block := NewBlock(
		1,
		0,
		[]byte("prev_hash"),
		[]*transaction.Transaction{tx},
	)

	block.Finalize(node)

	block.Hash = []byte("corrupted")

	valid, _ := block.Verify(signer, node.PublicKey)

	if valid {
		t.Fatal("Corrupted block hash should fail")
	}
}