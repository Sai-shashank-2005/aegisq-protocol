package identity

import (
	"encoding/base64"
	"fmt"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
)

type NodeIdentity struct {
	NodeID     string
	PublicKey  []byte
	PrivateKey []byte
	Signer     crypto.Signer
}

func NewNodeIdentity(nodeID string, signer crypto.Signer) (*NodeIdentity, error) {
	pub, priv, err := signer.GenerateKeyPair()
	if err != nil {
		return nil, err
	}

	return &NodeIdentity{
		NodeID:     nodeID,
		PublicKey:  pub,
		PrivateKey: priv,
		Signer:     signer,
	}, nil
}

func (n *NodeIdentity) Sign(message []byte) ([]byte, error) {
	return n.Signer.Sign(n.PrivateKey, message)
}

func (n *NodeIdentity) Verify(message []byte, signature []byte) bool {
	return n.Signer.Verify(n.PublicKey, message, signature)
}

func (n *NodeIdentity) Algorithm() string {
	return n.Signer.Algorithm()
}

func (n *NodeIdentity) PublicKeyBase64() string {
	return base64.StdEncoding.EncodeToString(n.PublicKey)
}

func (n *NodeIdentity) String() string {
	return fmt.Sprintf(
		"NodeID: %s\nPublicKey: %s\nAlgorithm: %s\n",
		n.NodeID,
		n.PublicKeyBase64(),
		n.Algorithm(),
	)
}