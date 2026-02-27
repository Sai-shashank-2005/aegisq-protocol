package identity

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
)

// NodeIdentity represents a validator or participant in the network.
type NodeIdentity struct {
	NodeID     string
	PublicKey  []byte
	privateKey []byte
	signer     crypto.Signer
}

// NewNodeIdentity generates a new node identity using the provided signer.
func NewNodeIdentity(nodeID string, signer crypto.Signer) (*NodeIdentity, error) {
	if nodeID == "" {
		return nil, errors.New("nodeID cannot be empty")
	}

	pub, priv, err := signer.GenerateKeyPair()
	if err != nil {
		return nil, err
	}

	return &NodeIdentity{
		NodeID:     nodeID,
		PublicKey:  pub,
		privateKey: priv,
		signer:     signer,
	}, nil
}

// Sign signs arbitrary message data.
func (n *NodeIdentity) Sign(message []byte) ([]byte, error) {
	return n.signer.Sign(n.privateKey, message)
}

// Verify verifies signature using node's public key.
func (n *NodeIdentity) Verify(message []byte, signature []byte) bool {
	return n.signer.Verify(n.PublicKey, message, signature)
}

// PublicKeyBase64 returns public key encoded as base64.
func (n *NodeIdentity) PublicKeyBase64() string {
	return base64.StdEncoding.EncodeToString(n.PublicKey)
}

// String implements formatted identity output.
func (n *NodeIdentity) String() string {
	return fmt.Sprintf(
		"NodeID: %s\nPublicKey(Base64): %s\n",
		n.NodeID,
		n.PublicKeyBase64(),
	)
}