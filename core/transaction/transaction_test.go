package transaction

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
)

func TestTransactionSignVerify(t *testing.T) {
	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := NewTransaction("validator-1", "QmCID123", "Test File")

	err := tx.SignWithIdentity(node)
	if err != nil {
		t.Fatal(err)
	}

	valid, err := tx.Verify(signer, node.PublicKey)
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatal("Transaction signature invalid")
	}
}

func TestTransactionTamperFails(t *testing.T) {
	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := NewTransaction("validator-1", "QmCID123", "Test File")
	tx.SignWithIdentity(node)

	// Tamper metadata
	tx.Metadata = "Hacked File"

	valid, _ := tx.Verify(signer, node.PublicKey)
	if valid {
		t.Fatal("Tampered transaction should fail verification")
	}
}