package simulation

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

type SyntheticFileMetadata struct {
	OwnerID     string
	CID         string
	EncryptedKE []byte
	EncryptedKS []byte
	FileHash    []byte
	Timestamp   int64
}

func GenerateRandomBytes(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func GenerateFakeCID() string {
	random := GenerateRandomBytes(32)
	return "Qm" + hex.EncodeToString(random)
}

func GenerateSyntheticMetadata(ownerID string) SyntheticFileMetadata {

	fileContent := GenerateRandomBytes(2048)
	hash := sha256.Sum256(fileContent)

	return SyntheticFileMetadata{
		OwnerID:     ownerID,
		CID:         GenerateFakeCID(),
		EncryptedKE: GenerateRandomBytes(1024),
		EncryptedKS: GenerateRandomBytes(1024),
		FileHash:    hash[:],
		Timestamp:   time.Now().Unix(),
	}
}

// Deterministic binary encoding
func encodeMetadataBinary(m SyntheticFileMetadata) ([]byte, error) {

	buf := new(bytes.Buffer)

	writeString := func(s string) error {
		if err := binary.Write(buf, binary.LittleEndian, int32(len(s))); err != nil {
			return err
		}
		buf.WriteString(s)
		return nil
	}

	writeBytes := func(b []byte) error {
		if err := binary.Write(buf, binary.LittleEndian, int32(len(b))); err != nil {
			return err
		}
		buf.Write(b)
		return nil
	}

	if err := writeString(m.OwnerID); err != nil {
		return nil, err
	}
	if err := writeString(m.CID); err != nil {
		return nil, err
	}
	if err := writeBytes(m.EncryptedKE); err != nil {
		return nil, err
	}
	if err := writeBytes(m.EncryptedKS); err != nil {
		return nil, err
	}
	if err := writeBytes(m.FileHash); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, m.Timestamp); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

/*
Single transaction generation (kept for compatibility)
*/
func GenerateSyntheticTransaction(
	node *identity.NodeIdentity,
) (*transaction.Transaction, error) {

	metadata := GenerateSyntheticMetadata(node.NodeID)

	payloadBytes, err := encodeMetadataBinary(metadata)
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

/*
🔥 BATCH VERSION
This is where CGO collapse happens.
*/
func GenerateBulkTransactions(
	node *identity.NodeIdentity,
	count int,
) ([]*transaction.Transaction, error) {

	txs := make([]*transaction.Transaction, count)

	// Step 1 — Create unsigned transactions
	for i := 0; i < count; i++ {

		metadata := GenerateSyntheticMetadata(node.NodeID)

		payloadBytes, err := encodeMetadataBinary(metadata)
		if err != nil {
			return nil, err
		}

		tx := transaction.NewTransaction(
			node,
			"STORE_FILE",
			string(payloadBytes),
		)

		txs[i] = tx
	}

	// Step 2 — Collect hashes
	hashes := make([][]byte, count)
	privateKeys := make([][]byte, count)

	for i, tx := range txs {
		hash, err := tx.Hash()
		if err != nil {
			return nil, err
		}
		hashes[i] = hash
		privateKeys[i] = node.PrivateKey
	}

	// Step 3 — Dilithium batch path
	if batchSigner, ok := node.Signer.(*crypto.DilithiumSigner); ok {

		sigs, err := batchSigner.BatchSign(privateKeys, hashes)
		if err != nil {
			return nil, err
		}

		for i := range txs {
			txs[i].SetSignature(sigs[i])
		}

	} else {
		// ECDSA fallback
		for i := range txs {
			sig, err := node.Sign(hashes[i])
			if err != nil {
				return nil, err
			}
			txs[i].SetSignature(sig)
		}
	}

	return txs, nil
}
