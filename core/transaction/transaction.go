package transaction

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
)

type Transaction struct {
	Sender    string `json:"sender"`
	DataHash  string `json:"data_hash"`
	Metadata  string `json:"metadata"`
	Timestamp int64  `json:"timestamp"`
	Signature []byte `json:"signature"`
}

// NewTransaction creates a new unsigned transaction
func NewTransaction(sender, dataHash, metadata string) *Transaction {
	return &Transaction{
		Sender:    sender,
		DataHash:  dataHash,
		Metadata:  metadata,
		Timestamp: time.Now().Unix(),
	}
}

// computePayloadHash hashes transaction fields excluding signature
func (tx *Transaction) computePayloadHash() ([]byte, error) {
	if tx.Sender == "" || tx.DataHash == "" {
		return nil, errors.New("invalid transaction fields")
	}

	payload := struct {
		Sender    string `json:"sender"`
		DataHash  string `json:"data_hash"`
		Metadata  string `json:"metadata"`
		Timestamp int64  `json:"timestamp"`
	}{
		Sender:    tx.Sender,
		DataHash:  tx.DataHash,
		Metadata:  tx.Metadata,
		Timestamp: tx.Timestamp,
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return crypto.Hash(bytes), nil
}

// SignWithIdentity signs transaction using node identity
func (tx *Transaction) SignWithIdentity(node *identity.NodeIdentity) error {
	hash, err := tx.computePayloadHash()
	if err != nil {
		return err
	}

	signature, err := node.Sign(hash)
	if err != nil {
		return err
	}

	tx.Signature = signature
	return nil
}

// Verify verifies transaction signature using public key and signer
func (tx *Transaction) Verify(signer crypto.Signer, publicKey []byte) (bool, error) {
	hash, err := tx.computePayloadHash()
	if err != nil {
		return false, err
	}

	return signer.Verify(publicKey, hash, tx.Signature), nil
}