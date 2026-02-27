package crypto

import (
	"golang.org/x/crypto/sha3"
)

// Hash computes SHA3-256 hash of input data.
func Hash(data []byte) []byte {
	h := sha3.New256()
	h.Write(data)
	return h.Sum(nil)
}