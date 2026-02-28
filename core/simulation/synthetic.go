package simulation

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

type SyntheticFileMetadata struct {
	OwnerID     string `json:"owner_id"`
	CID         string `json:"cid"`
	EncryptedKE []byte `json:"encrypted_ke"`
	EncryptedKS []byte `json:"encrypted_ks"`
	FileHash    []byte `json:"file_hash"`
	Timestamp   int64  `json:"timestamp"`
}

// GenerateRandomBytes returns securely generated random bytes.
func GenerateRandomBytes(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

// GenerateFakeCID simulates an IPFS CID-like string.
func GenerateFakeCID() string {
	random := GenerateRandomBytes(32)
	return "Qm" + hex.EncodeToString(random)
}

// GenerateSyntheticMetadata builds realistic metadata.
func GenerateSyntheticMetadata(ownerID string) SyntheticFileMetadata {

	fileContent := GenerateRandomBytes(2048) // simulate 2KB file
	hash := sha256.Sum256(fileContent)

	return SyntheticFileMetadata{
		OwnerID:     ownerID,
		CID:         GenerateFakeCID(),
		EncryptedKE: GenerateRandomBytes(1024), // simulate Kyber ciphertext
		EncryptedKS: GenerateRandomBytes(1024),
		FileHash:    hash[:],
		Timestamp:   time.Now().Unix(),
	}
}

// GenerateSyntheticTransaction converts metadata into a signed transaction.
func GenerateSyntheticTransaction(
	node *identity.NodeIdentity,
) (*transaction.Transaction, error) {

	// IMPORTANT: Correct field name
	metadata := GenerateSyntheticMetadata(node.NodeID)

	payloadBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	tx := transaction.NewTransaction(
		node,
		"STORE_FILE",
		string(payloadBytes),
	)

	if err := tx.SignWithIdentity(node); err != nil {
		return nil, err
	}

	return tx, nil
}

// GenerateBulkTransactions creates multiple synthetic transactions.
func GenerateBulkTransactions(
	node *identity.NodeIdentity,
	count int,
) ([]*transaction.Transaction, error) {

	var txs []*transaction.Transaction

	for i := 0; i < count; i++ {
		tx, err := GenerateSyntheticTransaction(node)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}

	return txs, nil
}