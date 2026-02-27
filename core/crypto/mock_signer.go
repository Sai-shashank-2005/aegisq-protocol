package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

type MockSigner struct{}

func (m *MockSigner) GenerateKeyPair() ([]byte, []byte, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	return pub, priv, err
}

func (m *MockSigner) Sign(privateKey []byte, message []byte) ([]byte, error) {
	priv := ed25519.PrivateKey(privateKey)
	signature := ed25519.Sign(priv, message)
	return signature, nil
}

func (m *MockSigner) Verify(publicKey []byte, message []byte, signature []byte) bool {
	pub := ed25519.PublicKey(publicKey)
	return ed25519.Verify(pub, message, signature)
}