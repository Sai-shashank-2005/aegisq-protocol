package crypto

import (
	"crypto/sha3"
	"crypto/rand"
)

type PQCSigner struct{}

func (p *PQCSigner) GenerateKeyPair() ([]byte, []byte, error) {
	priv := make([]byte, 64)
	pub := make([]byte, 64)
	rand.Read(priv)
	rand.Read(pub)
	return pub, priv, nil
}

func (p *PQCSigner) Sign(privateKey []byte, message []byte) ([]byte, error) {
	hash := sha3.Sum256(append(privateKey, message...))
	return hash[:], nil
}

func (p *PQCSigner) Verify(publicKey []byte, message []byte, signature []byte) bool {
	hash := sha3.Sum256(append(publicKey, message...))
	return string(hash[:]) == string(signature)
}

func (p *PQCSigner) Algorithm() string {
	return "pqc-mock"
}