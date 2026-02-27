package ledger

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func TestFullChainWithDilithium(t *testing.T) {

	signer, err := crypto.NewDilithiumSigner()
	if err != nil {
		t.Fatal("Failed to initialize Dilithium signer")
	}

	node, err := identity.NewNodeIdentity("validator-pqc", signer)
	if err != nil {
		t.Fatal("Failed to create identity")
	}

	// Genesis block
	txGenesis := transaction.NewTransaction(node, "genesis", "pqc")
	txGenesis.SignWithIdentity(node)

	genesis := block.NewBlock(0, []byte("genesis"), []*transaction.Transaction{txGenesis})
	genesis.Finalize(node)

	ledger := NewLedger(genesis)

	// Block 1
	tx1 := transaction.NewTransaction(node, "data1", "meta1")
	tx1.SignWithIdentity(node)

	block1 := block.NewBlock(1, genesis.Hash, []*transaction.Transaction{tx1})
	block1.Finalize(node)

	err = ledger.AddBlock(block1, signer, node.PublicKey)
	if err != nil {
		t.Fatal("Failed to add PQC block")
	}

	err = ledger.ValidateChain(signer, node.PublicKey)
	if err != nil {
		t.Fatal("PQC chain validation failed")
	}
}