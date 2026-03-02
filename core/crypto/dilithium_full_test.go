package crypto

import "testing"

func getSigners(t *testing.T) map[string]Signer {
	signers := make(map[string]Signer)

	dilithium, err := NewDilithiumSigner()
	if err != nil {
		t.Fatal(err)
	}
	signers["Dilithium"] = dilithium

	ecdsa, err := NewECDSASigner()
	if err != nil {
		t.Fatal(err)
	}
	signers["ECDSA"] = ecdsa

	return signers
}

func TestSignerInitialization(t *testing.T) {
	for name, signer := range getSigners(t) {
		t.Run(name, func(t *testing.T) {
			if signer == nil {
				t.Fatal("Signer is nil")
			}
		})
	}
}

func TestKeyGeneration(t *testing.T) {
	for name, signer := range getSigners(t) {
		t.Run(name, func(t *testing.T) {

			pub, priv, err := signer.GenerateKeyPair()
			if err != nil {
				t.Fatal(err)
			}

			if len(pub) == 0 || len(priv) == 0 {
				t.Fatal("Key generation failed")
			}
		})
	}
}

func TestSignVerify(t *testing.T) {
	for name, signer := range getSigners(t) {
		t.Run(name, func(t *testing.T) {

			pub, priv, _ := signer.GenerateKeyPair()
			msg := []byte("Post-Quantum Blockchain Test")

			sig, err := signer.Sign(priv, msg)
			if err != nil {
				t.Fatal(err)
			}

			if !signer.Verify(pub, msg, sig) {
				t.Fatal("Signature verification failed")
			}
		})
	}
}

func TestMutationFails(t *testing.T) {
	for name, signer := range getSigners(t) {
		t.Run(name, func(t *testing.T) {

			pub, priv, _ := signer.GenerateKeyPair()

			msg := []byte("Original Message")
			sig, _ := signer.Sign(priv, msg)

			mutated := []byte("Tampered Message")

			if signer.Verify(pub, mutated, sig) {
				t.Fatal("Verification should fail for modified message")
			}
		})
	}
}

func TestReplayFails(t *testing.T) {
	for name, signer := range getSigners(t) {
		t.Run(name, func(t *testing.T) {

			pub1, priv1, _ := signer.GenerateKeyPair()
			pub2, _, _ := signer.GenerateKeyPair()

			msg := []byte("Replay Test")

			sig, _ := signer.Sign(priv1, msg)

			if signer.Verify(pub2, msg, sig) {
				t.Fatal("Replay succeeded with wrong public key")
			}

			if !signer.Verify(pub1, msg, sig) {
				t.Fatal("Valid signature rejected")
			}
		})
	}
}

func TestDeterministicVerification(t *testing.T) {
	for name, signer := range getSigners(t) {
		t.Run(name, func(t *testing.T) {

			pub, priv, _ := signer.GenerateKeyPair()
			msg := []byte("Consistency Test")

			sig, _ := signer.Sign(priv, msg)

			for i := 0; i < 10; i++ {
				if !signer.Verify(pub, msg, sig) {
					t.Fatal("Verification failed on repeated checks")
				}
			}
		})
	}
}

func TestSignatureCorruptionFails(t *testing.T) {
	for name, signer := range getSigners(t) {
		t.Run(name, func(t *testing.T) {

			pub, priv, _ := signer.GenerateKeyPair()
			msg := []byte("Corruption Test")

			sig, _ := signer.Sign(priv, msg)

			corrupted := make([]byte, len(sig))
			copy(corrupted, sig)
			corrupted[10] ^= 0xFF

			if signer.Verify(pub, msg, corrupted) {
				t.Fatal("Corrupted signature verified")
			}
		})
	}
}
