package crypto

import "testing"

func BenchmarkDilithiumKeyGen(b *testing.B) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	for i := 0; i < b.N; i++ {
		signer.GenerateKeyPair()
	}
}

func BenchmarkDilithiumSign(b *testing.B) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	_, priv, _ := signer.GenerateKeyPair()
	msg := []byte("Benchmark Message")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		signer.Sign(priv, msg)
	}
}

func BenchmarkDilithiumVerify(b *testing.B) {
	signer, _ := NewDilithiumSigner()
	defer signer.Close()

	pub, priv, _ := signer.GenerateKeyPair()
	msg := []byte("Benchmark Message")
	sig, _ := signer.Sign(priv, msg)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		signer.Verify(pub, msg, sig)
	}
}