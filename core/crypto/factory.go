package crypto

import (
	"fmt"
	"os"
)

func NewDefaultSigner() (Signer, error) {

	switch os.Getenv("CRYPTO_ALG") {

	case "ecdsa":
		s, err := NewECDSASigner()
		fmt.Println("Using signer:", s.Algorithm())
		return s, err

	default:
		s, err := NewDilithiumSigner()
		fmt.Println("Using signer:", s.Algorithm())
		return s, err
	}
}
