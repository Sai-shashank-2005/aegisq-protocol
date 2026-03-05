package simulation

import (
	"crypto/rand"
	"fmt"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

// GenerateSyntheticDataset generates N realistic storage transactions.
// Each transaction simulates hashing a random file-like input.
func GenerateSyntheticDataset(
	count int,
	node *identity.NodeIdentity,
) ([]*transaction.Transaction, error) {

	var txs []*transaction.Transaction

	for i := 0; i < count; i++ {

		// Simulate random file data (1KB)
		rawData := make([]byte, 1024)
		_, err := rand.Read(rawData)
		if err != nil {
			return nil, err
		}

		// Hash raw input (this is what blockchain stores)
		dataHash := crypto.Hash(rawData)

		tx := transaction.NewTransaction(
			node,
			fmt.Sprintf("%x", dataHash),
			fmt.Sprintf("Synthetic File Upload #%d", i),
		)

		err = tx.SignWithIdentity(node)
		if err != nil {
			return nil, err
		}

		txs = append(txs, tx)
	}

	return txs, nil
}
