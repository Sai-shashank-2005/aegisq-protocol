package crypto

type Signer interface {
	GenerateKeyPair() ([]byte, []byte, error)
	Sign(privateKey []byte, message []byte) ([]byte, error)
	Verify(publicKey []byte, message []byte, signature []byte) bool
	Algorithm() string
}