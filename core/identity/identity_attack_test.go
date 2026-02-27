package identity

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
)

func TestAttack_WrongPublicKey(t *testing.T) {
	signer := &crypto.Ed25519Signer{}

	node1, _ := NewNodeIdentity("node1", signer)
	node2, _ := NewNodeIdentity("node2", signer)

	msg := []byte("secure message")
	sig, _ := node1.Sign(msg)

	if signer.Verify(node2.PublicKey, msg, sig) {
		t.Fatal("Layer1 failed: signature verified with wrong public key")
	}
}
