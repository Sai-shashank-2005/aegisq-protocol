package crypto

import (
	"crypto/sha256"
)

// Hash computes SHA3-256 hash of input data.
func Hash(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}
