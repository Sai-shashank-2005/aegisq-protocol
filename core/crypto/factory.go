package crypto

func NewDefaultSigner() (Signer, error) {
    return NewDilithiumSigner()
}