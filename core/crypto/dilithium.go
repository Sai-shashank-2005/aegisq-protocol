package crypto

/*
#cgo LDFLAGS: -loqs
#include <oqs/oqs.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type DilithiumSigner struct {
	alg *C.OQS_SIG
}

func NewDilithiumSigner() (*DilithiumSigner, error) {
	alg := C.OQS_SIG_new(C.CString("Dilithium2"))
	if alg == nil {
		return nil, errors.New("failed to initialize Dilithium")
	}
	return &DilithiumSigner{alg: alg}, nil
}

func (d *DilithiumSigner) GenerateKeyPair() ([]byte, []byte, error) {
	pub := C.malloc(C.size_t(d.alg.length_public_key))
	priv := C.malloc(C.size_t(d.alg.length_secret_key))

	res := C.OQS_SIG_keypair(d.alg,
		(*C.uint8_t)(pub),
		(*C.uint8_t)(priv),
	)

	if res != C.OQS_SUCCESS {
		return nil, nil, errors.New("keypair generation failed")
	}

	publicKey := C.GoBytes(pub, C.int(d.alg.length_public_key))
	privateKey := C.GoBytes(priv, C.int(d.alg.length_secret_key))

	C.free(pub)
	C.free(priv)

	return publicKey, privateKey, nil
}

func (d *DilithiumSigner) Sign(privateKey []byte, message []byte) ([]byte, error) {

	sig := C.malloc(C.size_t(d.alg.length_signature))
	var sigLen C.size_t

	res := C.OQS_SIG_sign(d.alg,
		(*C.uint8_t)(sig),
		&sigLen,
		(*C.uint8_t)(unsafe.Pointer(&message[0])),
		C.size_t(len(message)),
		(*C.uint8_t)(unsafe.Pointer(&privateKey[0])),
	)

	if res != C.OQS_SUCCESS {
		return nil, errors.New("sign failed")
	}

	signature := C.GoBytes(sig, C.int(sigLen))
	C.free(sig)

	return signature, nil
}

func (d *DilithiumSigner) Verify(publicKey []byte, message []byte, signature []byte) bool {

	res := C.OQS_SIG_verify(d.alg,
		(*C.uint8_t)(unsafe.Pointer(&message[0])),
		C.size_t(len(message)),
		(*C.uint8_t)(unsafe.Pointer(&signature[0])),
		C.size_t(len(signature)),
		(*C.uint8_t)(unsafe.Pointer(&publicKey[0])),
	)

	return res == C.OQS_SUCCESS
}

func (d *DilithiumSigner) Algorithm() string {
	return "dilithium2"
}