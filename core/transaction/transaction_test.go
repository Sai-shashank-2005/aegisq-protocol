package transaction

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
)

func TestTransactionSignVerify(t *testing.T) {
	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := NewTransaction(node, "QmCID123", "Test File")
	tx.SignWithIdentity(node)

	valid, err := tx.Verify(signer)
	if err != nil || !valid {
		t.Fatal("Transaction verification failed")
	}
}

func TestTransactionTamperFails(t *testing.T) {
	signer := &crypto.Ed25519Signer{}
	node, _ := identity.NewNodeIdentity("validator-1", signer)

	tx := NewTransaction(node, "QmCID123", "Test File")
	tx.SignWithIdentity(node)

	tx.Metadata = "HACKED"

	valid, _ := tx.Verify(signer)
	if valid {
		t.Fatal("Tampered transaction should fail")
	}
}