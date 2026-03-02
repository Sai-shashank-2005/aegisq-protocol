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

// Constructor
func NewDilithiumSigner() (*DilithiumSigner, error) {

	name := C.CString("ML-DSA-44")
	defer C.free(unsafe.Pointer(name))

	alg := C.OQS_SIG_new(name)
	if alg == nil {
		return nil, errors.New("failed to initialize Dilithium2")
	}

	return &DilithiumSigner{alg: alg}, nil
}

// GenerateKeyPair generates public and private keys
func (d *DilithiumSigner) GenerateKeyPair() ([]byte, []byte, error) {

	if d.alg == nil {
		return nil, nil, errors.New("Dilithium signer not initialized")
	}

	pub := C.malloc(C.size_t(d.alg.length_public_key))
	priv := C.malloc(C.size_t(d.alg.length_secret_key))

	if pub == nil || priv == nil {
		return nil, nil, errors.New("memory allocation failed")
	}

	defer C.free(pub)
	defer C.free(priv)

	res := C.OQS_SIG_keypair(
		d.alg,
		(*C.uint8_t)(pub),
		(*C.uint8_t)(priv),
	)

	if res != C.OQS_SUCCESS {
		return nil, nil, errors.New("keypair generation failed")
	}

	publicKey := C.GoBytes(pub, C.int(d.alg.length_public_key))
	privateKey := C.GoBytes(priv, C.int(d.alg.length_secret_key))

	return publicKey, privateKey, nil
}

// Sign signs a message using Dilithium
func (d *DilithiumSigner) Sign(privateKey []byte, message []byte) ([]byte, error) {

	if d.alg == nil {
		return nil, errors.New("Dilithium signer not initialized")
	}

	if len(privateKey) == 0 || len(message) == 0 {
		return nil, errors.New("invalid input to Sign")
	}

	sig := C.malloc(C.size_t(d.alg.length_signature))
	if sig == nil {
		return nil, errors.New("memory allocation failed")
	}
	defer C.free(sig)

	var sigLen C.size_t

	res := C.OQS_SIG_sign(
		d.alg,
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
	return signature, nil
}

// Verify verifies a Dilithium signature
func (d *DilithiumSigner) Verify(publicKey []byte, message []byte, signature []byte) bool {

	if d.alg == nil {
		return false
	}

	if len(publicKey) == 0 || len(message) == 0 || len(signature) == 0 {
		return false
	}

	res := C.OQS_SIG_verify(
		d.alg,
		(*C.uint8_t)(unsafe.Pointer(&message[0])),
		C.size_t(len(message)),
		(*C.uint8_t)(unsafe.Pointer(&signature[0])),
		C.size_t(len(signature)),
		(*C.uint8_t)(unsafe.Pointer(&publicKey[0])),
	)

	return res == C.OQS_SUCCESS
}

// Algorithm returns algorithm identifier
func (d *DilithiumSigner) Algorithm() string {
	return "dilithium2"
}

// Close frees underlying C memory (CRITICAL for long-running systems)
func (d *DilithiumSigner) Close() {
	if d.alg != nil {
		C.OQS_SIG_free(d.alg)
		d.alg = nil
	}
}
