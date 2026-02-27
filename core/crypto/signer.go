package crypto

// Signer defines the cryptographic signature interface.
// This abstraction allows replacing Ed25519 with Dilithium seamlessly.
type Signer interface {
	GenerateKeyPair() (publicKey []byte, privateKey []byte, err error)
	Sign(privateKey []byte, message []byte) ([]byte, error)
	Verify(publicKey []byte, message []byte, signature []byte) bool
}