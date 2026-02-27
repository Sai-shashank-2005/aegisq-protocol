package transaction

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
)

type Transaction struct {
	SenderID  string `json:"sender_id"`
	PublicKey []byte `json:"public_key"`
	Algorithm string `json:"algorithm"`
	DataHash  string `json:"data_hash"`
	Metadata  string `json:"metadata"`
	Timestamp int64  `json:"timestamp"`
	Signature []byte `json:"signature"`
}

func NewTransaction(node *identity.NodeIdentity, dataHash, metadata string) *Transaction {
	return &Transaction{
		SenderID:  node.NodeID,
		PublicKey: node.PublicKey,
		Algorithm: node.Algorithm(),
		DataHash:  dataHash,
		Metadata:  metadata,
		Timestamp: time.Now().Unix(),
	}
}

func (tx *Transaction) computePayloadHash() ([]byte, error) {
	if tx.SenderID == "" || tx.DataHash == "" {
		return nil, errors.New("invalid transaction fields")
	}

	payload := struct {
		SenderID  string `json:"sender_id"`
		PublicKey []byte `json:"public_key"`
		Algorithm string `json:"algorithm"`
		DataHash  string `json:"data_hash"`
		Metadata  string `json:"metadata"`
		Timestamp int64  `json:"timestamp"`
	}{
		SenderID:  tx.SenderID,
		PublicKey: tx.PublicKey,
		Algorithm: tx.Algorithm,
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

func (tx *Transaction) Hash() ([]byte, error) {
	return tx.computePayloadHash()
}

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

func (tx *Transaction) Verify(signer crypto.Signer) (bool, error) {
	if tx.Algorithm != signer.Algorithm() {
		return false, errors.New("algorithm mismatch")
	}

	hash, err := tx.computePayloadHash()
	if err != nil {
		return false, err
	}

	return signer.Verify(tx.PublicKey, hash, tx.Signature), nil
}