package crypto

import (
	"testing"
)

func TestHashDeterministic(t *testing.T) {
	data := []byte("AegisQ")

	h1 := Hash(data)
	h2 := Hash(data)

	if string(h1) != string(h2) {
		t.Fatal("Hash function is not deterministic")
	}
}

func TestHashChangesOnInputChange(t *testing.T) {
	h1 := Hash([]byte("A"))
	h2 := Hash([]byte("B"))

	if string(h1) == string(h2) {
		t.Fatal("Hash should differ for different inputs")
	}
}