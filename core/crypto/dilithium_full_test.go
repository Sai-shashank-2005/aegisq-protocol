package crypto

import (
	"testing"
)

func TestDilithiumInitialization(t *testing.T) {
	signer, err := NewDilithiumSigner()
	if err != nil {
		t.Fatalf("Failed to initialize Dilithium signer: %v", err)
	}
	defer signer.Close()
}

func TestDilithiumKeyGenerationSizes(t *testing.T) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	pub, priv, err := signer.GenerateKeyPair()
	if err != nil {
		t.Fatal("Key generation failed")
	}

	if len(pub) < 1000 {
		t.Fatal("Public key size too small — not Dilithium")
	}

	if len(priv) < 2000 {
		t.Fatal("Private key size too small — not Dilithium")
	}
}

func TestDilithiumSignVerify(t *testing.T) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	pub, priv, _ := signer.GenerateKeyPair()

	msg := []byte("Post-Quantum Blockchain Test")

	sig, err := signer.Sign(priv, msg)
	if err != nil {
		t.Fatal("Signing failed")
	}

	if len(sig) < 2000 {
		t.Fatal("Signature too small — not Dilithium")
	}

	valid := signer.Verify(pub, msg, sig)
	if !valid {
		t.Fatal("Signature verification failed")
	}
}

func TestDilithiumMutationFails(t *testing.T) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	pub, priv, _ := signer.GenerateKeyPair()

	msg := []byte("Original Message")

	sig, _ := signer.Sign(priv, msg)

	mutated := []byte("Tampered Message")

	if signer.Verify(pub, mutated, sig) {
		t.Fatal("Verification should fail for modified message")
	}
}

func TestDilithiumSignatureReplayFails(t *testing.T) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	pub1, priv1, _ := signer.GenerateKeyPair()
	pub2, _, _ := signer.GenerateKeyPair()

	msg := []byte("Replay Test")

	sig, _ := signer.Sign(priv1, msg)

	// Attempt verification using different public key
	if signer.Verify(pub2, msg, sig) {
		t.Fatal("Replay attack succeeded with wrong public key")
	}

	// Valid case
	if !signer.Verify(pub1, msg, sig) {
		t.Fatal("Valid signature rejected")
	}
}

func TestDilithiumDeterministicVerification(t *testing.T) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	pub, priv, _ := signer.GenerateKeyPair()
	msg := []byte("Consistency Test")

	sig, _ := signer.Sign(priv, msg)

	for i := 0; i < 10; i++ {
		if !signer.Verify(pub, msg, sig) {
			t.Fatal("Verification failed on repeated checks")
		}
	}
}

func TestDilithiumSignatureCorruptionFails(t *testing.T) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	pub, priv, _ := signer.GenerateKeyPair()
	msg := []byte("Corruption Test")

	sig, _ := signer.Sign(priv, msg)

	// Flip a byte
	corrupted := make([]byte, len(sig))
	copy(corrupted, sig)
	corrupted[10] ^= 0xFF

	if signer.Verify(pub, msg, corrupted) {
		t.Fatal("Corrupted signature verified successfully")
	}
}