package crypto

import "testing"

//////////////////////////////
// DILITHIUM BENCHMARKS
//////////////////////////////

func BenchmarkDilithiumKeyGen(b *testing.B) {
	signer, err := NewDilithiumSigner()
	if err != nil {
		b.Fatal(err)
	}
	defer signer.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := signer.GenerateKeyPair()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDilithiumSign(b *testing.B) {
	signer, err := NewDilithiumSigner()
	if err != nil {
		b.Fatal(err)
	}
	defer signer.Close()

	_, priv, err := signer.GenerateKeyPair()
	if err != nil {
		b.Fatal(err)
	}

	msg := []byte("Benchmark Message")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := signer.Sign(priv, msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDilithiumVerify(b *testing.B) {
	signer, err := NewDilithiumSigner()
	if err != nil {
		b.Fatal(err)
	}
	defer signer.Close()

	pub, priv, err := signer.GenerateKeyPair()
	if err != nil {
		b.Fatal(err)
	}

	msg := []byte("Benchmark Message")

	sig, err := signer.Sign(priv, msg)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if !signer.Verify(pub, msg, sig) {
			b.Fatal("verify failed")
		}
	}
}

//////////////////////////////
// ECDSA BENCHMARKS
//////////////////////////////

func BenchmarkECDSAKeyGen(b *testing.B) {
	signer, err := NewECDSASigner()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := signer.GenerateKeyPair()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkECDSASign(b *testing.B) {
	signer, err := NewECDSASigner()
	if err != nil {
		b.Fatal(err)
	}

	_, priv, err := signer.GenerateKeyPair()
	if err != nil {
		b.Fatal(err)
	}

	msg := []byte("Benchmark Message")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := signer.Sign(priv, msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkECDSAVerify(b *testing.B) {
	signer, err := NewECDSASigner()
	if err != nil {
		b.Fatal(err)
	}

	pub, priv, err := signer.GenerateKeyPair()
	if err != nil {
		b.Fatal(err)
	}

	msg := []byte("Benchmark Message")

	sig, err := signer.Sign(priv, msg)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if !signer.Verify(pub, msg, sig) {
			b.Fatal("verify failed")
		}
	}
}
