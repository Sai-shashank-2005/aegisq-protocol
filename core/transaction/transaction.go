package transaction

import (
	"bytes"
	"encoding/binary"
	"errors"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
)

type Transaction struct {
	SenderID  string
	PublicKey []byte
	Algorithm string
	DataHash  string
	Metadata  string
	Timestamp int64
	Signature []byte
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

	buf := new(bytes.Buffer)

	// SenderID
	if err := binary.Write(buf, binary.LittleEndian, int32(len(tx.SenderID))); err != nil {
		return nil, err
	}
	buf.WriteString(tx.SenderID)

	// PublicKey
	if err := binary.Write(buf, binary.LittleEndian, int32(len(tx.PublicKey))); err != nil {
		return nil, err
	}
	buf.Write(tx.PublicKey)

	// Algorithm
	if err := binary.Write(buf, binary.LittleEndian, int32(len(tx.Algorithm))); err != nil {
		return nil, err
	}
	buf.WriteString(tx.Algorithm)

	// DataHash
	if err := binary.Write(buf, binary.LittleEndian, int32(len(tx.DataHash))); err != nil {
		return nil, err
	}
	buf.WriteString(tx.DataHash)

	// Metadata
	if err := binary.Write(buf, binary.LittleEndian, int32(len(tx.Metadata))); err != nil {
		return nil, err
	}
	buf.WriteString(tx.Metadata)

	// Timestamp
	if err := binary.Write(buf, binary.LittleEndian, tx.Timestamp); err != nil {
		return nil, err
	}

	return crypto.Hash(buf.Bytes()), nil
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

// 🔥 REQUIRED for batch signing
func (tx *Transaction) SetSignature(sig []byte) {
	tx.Signature = sig
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
