package simulation

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
)

func TestSyntheticTransactionGeneration(t *testing.T) {

	signer := &crypto.Ed25519Signer{}

	node, err := identity.NewNodeIdentity("validator-1", signer)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := GenerateSyntheticTransaction(node)
	if err != nil {
		t.Fatal(err)
	}

	if tx == nil {
		t.Fatal("Generated transaction is nil")
	}
}

func TestBulkGeneration(t *testing.T) {

	signer := &crypto.Ed25519Signer{}

	node, err := identity.NewNodeIdentity("validator-1", signer)
	if err != nil {
		t.Fatal(err)
	}

	txs, err := GenerateBulkTransactions(node, 1000)
	if err != nil {
		t.Fatal(err)
	}

	if len(txs) != 1000 {
		t.Fatal("Bulk generation failed")
	}
}