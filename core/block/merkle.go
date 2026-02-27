package block

import (
	"bytes"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
)

func ComputeMerkleRoot(hashes [][]byte) []byte {
	if len(hashes) == 0 {
		return nil
	}

	if len(hashes) == 1 {
		return hashes[0]
	}

	var nextLevel [][]byte

	for i := 0; i < len(hashes); i += 2 {
		if i+1 == len(hashes) {
			nextLevel = append(nextLevel, hashes[i])
		} else {
			combined := bytes.Join([][]byte{hashes[i], hashes[i+1]}, nil)
			nextLevel = append(nextLevel, crypto.Hash(combined))
		}
	}

	return ComputeMerkleRoot(nextLevel)
}