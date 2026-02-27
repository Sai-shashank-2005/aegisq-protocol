package identity

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
)

func TestIdentitySignVerify(t *testing.T) {
	signer := &crypto.Ed25519Signer{}

	node, err := NewNodeIdentity("validator-1", signer)
	if err != nil {
		t.Fatal(err)
	}

	message := []byte("Test message")

	signature, err := node.Sign(message)
	if err != nil {
		t.Fatal(err)
	}

	if !node.Verify(message, signature) {
		t.Fatal("Signature verification failed")
	}
}

func TestSignatureFailsOnModifiedMessage(t *testing.T) {
	signer := &crypto.Ed25519Signer{}

	node, _ := NewNodeIdentity("validator-1", signer)

	message := []byte("Original message")
	signature, _ := node.Sign(message)

	modified := []byte("Tampered message")

	if node.Verify(modified, signature) {
		t.Fatal("Signature should fail for modified message")
	}
}