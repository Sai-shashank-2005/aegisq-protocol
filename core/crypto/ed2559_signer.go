package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

// Ed25519Signer is a temporary development signer.
// This will be replaced with Dilithium in production.
type Ed25519Signer struct{}

func (e *Ed25519Signer) GenerateKeyPair() ([]byte, []byte, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	return pub, priv, err
}

func (e *Ed25519Signer) Sign(privateKey []byte, message []byte) ([]byte, error) {
	priv := ed25519.PrivateKey(privateKey)
	signature := ed25519.Sign(priv, message)
	return signature, nil
}

func (e *Ed25519Signer) Verify(publicKey []byte, message []byte, signature []byte) bool {
	pub := ed25519.PublicKey(publicKey)
	return ed25519.Verify(pub, message, signature)
}