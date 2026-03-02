package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type ECDSASigner struct{}

func NewECDSASigner() (*ECDSASigner, error) {
	return &ECDSASigner{}, nil
}

func (e *ECDSASigner) GenerateKeyPair() ([]byte, []byte, error) {

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// 65-byte uncompressed public key
	pub := elliptic.Marshal(elliptic.P256(), priv.X, priv.Y)

	// 32-byte fixed private key
	privBytes := make([]byte, 32)
	dBytes := priv.D.Bytes()
	copy(privBytes[32-len(dBytes):], dBytes)

	return pub, privBytes, nil
}

func (e *ECDSASigner) Sign(privateKey []byte, message []byte) ([]byte, error) {

	curve := elliptic.P256()

	d := new(big.Int).SetBytes(privateKey)

	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = d
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(privateKey)

	hash := sha256.Sum256(message)

	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return nil, err
	}

	// FIXED 64-byte signature
	sig := make([]byte, 64)

	rb := r.Bytes()
	sb := s.Bytes()

	copy(sig[32-len(rb):32], rb)
	copy(sig[64-len(sb):], sb)

	return sig, nil
}

func (e *ECDSASigner) Verify(publicKey, message, signature []byte) bool {

	curve := elliptic.P256()

	x, y := elliptic.Unmarshal(curve, publicKey)
	if x == nil {
		return false
	}

	pub := ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	hash := sha256.Sum256(message)

	if len(signature) != 64 {
		return false
	}

	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:])

	return ecdsa.Verify(&pub, hash[:], r, s)
}

func (e *ECDSASigner) Algorithm() string {
	return "ECDSA_P256"
}
