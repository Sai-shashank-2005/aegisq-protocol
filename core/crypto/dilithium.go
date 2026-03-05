package crypto

/*
#cgo LDFLAGS: -loqs
#include <oqs/oqs.h>
#include <stdlib.h>

int batch_sign(
    OQS_SIG *alg,
    uint8_t **messages,
    size_t *msg_lens,
    uint8_t **private_keys,
    uint8_t **signatures,
    size_t *sig_lens,
    size_t count
);

int batch_sign(
    OQS_SIG *alg,
    uint8_t **messages,
    size_t *msg_lens,
    uint8_t **private_keys,
    uint8_t **signatures,
    size_t *sig_lens,
    size_t count
) {
    for (size_t i = 0; i < count; i++) {
        if (OQS_SIG_sign(
            alg,
            signatures[i],
            &sig_lens[i],
            messages[i],
            msg_lens[i],
            private_keys[i]
        ) != OQS_SUCCESS) {
            return -1;
        }
    }
    return 0;
}
*/
import "C"

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

var cgoCallCount uint64

func GetCGOCallCount() uint64 {
	return atomic.LoadUint64(&cgoCallCount)
}

func ResetCGOCallCount() {
	atomic.StoreUint64(&cgoCallCount, 0)
}

type DilithiumSigner struct {
	alg *C.OQS_SIG
}

func NewDilithiumSigner() (*DilithiumSigner, error) {

	name := C.CString("ML-DSA-44")
	defer C.free(unsafe.Pointer(name))

	atomic.AddUint64(&cgoCallCount, 1)
	alg := C.OQS_SIG_new(name)
	if alg == nil {
		return nil, errors.New("failed to initialize Dilithium")
	}

	return &DilithiumSigner{alg: alg}, nil
}

func (d *DilithiumSigner) GenerateKeyPair() ([]byte, []byte, error) {

	if d.alg == nil {
		return nil, nil, errors.New("Dilithium signer not initialized")
	}

	pubLen := int(d.alg.length_public_key)
	privLen := int(d.alg.length_secret_key)

	publicKey := make([]byte, pubLen)
	privateKey := make([]byte, privLen)

	atomic.AddUint64(&cgoCallCount, 1)
	res := C.OQS_SIG_keypair(
		d.alg,
		(*C.uint8_t)(unsafe.Pointer(&publicKey[0])),
		(*C.uint8_t)(unsafe.Pointer(&privateKey[0])),
	)

	if res != C.OQS_SUCCESS {
		return nil, nil, errors.New("keypair generation failed")
	}

	return publicKey, privateKey, nil
}

func (d *DilithiumSigner) Sign(privateKey []byte, message []byte) ([]byte, error) {

	if d.alg == nil {
		return nil, errors.New("Dilithium signer not initialized")
	}

	if len(privateKey) == 0 || len(message) == 0 {
		return nil, errors.New("invalid input to Sign")
	}

	sigMax := int(d.alg.length_signature)
	signature := make([]byte, sigMax)

	var sigLen C.size_t

	atomic.AddUint64(&cgoCallCount, 1)
	res := C.OQS_SIG_sign(
		d.alg,
		(*C.uint8_t)(unsafe.Pointer(&signature[0])),
		&sigLen,
		(*C.uint8_t)(unsafe.Pointer(&message[0])),
		C.size_t(len(message)),
		(*C.uint8_t)(unsafe.Pointer(&privateKey[0])),
	)

	if res != C.OQS_SUCCESS {
		return nil, errors.New("sign failed")
	}

	return signature[:sigLen], nil
}

func (d *DilithiumSigner) BatchSign(
	privateKeys [][]byte,
	messages [][]byte,
) ([][]byte, error) {

	if d.alg == nil {
		return nil, errors.New("Dilithium signer not initialized")
	}

	if len(privateKeys) != len(messages) {
		return nil, errors.New("batch size mismatch")
	}

	count := len(messages)
	if count == 0 {
		return nil, errors.New("empty batch")
	}

	sigMax := int(d.alg.length_signature)
	signatures := make([][]byte, count)

	ptrSize := unsafe.Sizeof(uintptr(0))

	// Allocate pointer arrays in C memory
	cMessages := C.malloc(C.size_t(count) * C.size_t(ptrSize))
	cMsgLens := C.malloc(C.size_t(count) * C.size_t(unsafe.Sizeof(C.size_t(0))))
	cPrivKeys := C.malloc(C.size_t(count) * C.size_t(ptrSize))
	cSigs := C.malloc(C.size_t(count) * C.size_t(ptrSize))
	cSigLens := C.malloc(C.size_t(count) * C.size_t(unsafe.Sizeof(C.size_t(0))))

	if cMessages == nil || cMsgLens == nil || cPrivKeys == nil || cSigs == nil || cSigLens == nil {
		return nil, errors.New("C allocation failed")
	}

	defer C.free(cMessages)
	defer C.free(cMsgLens)
	defer C.free(cPrivKeys)
	defer C.free(cSigs)
	defer C.free(cSigLens)

	msgPtrArray := (*[1 << 30]*C.uint8_t)(cMessages)[:count:count]
	msgLenArray := (*[1 << 30]C.size_t)(cMsgLens)[:count:count]
	privPtrArray := (*[1 << 30]*C.uint8_t)(cPrivKeys)[:count:count]
	sigPtrArray := (*[1 << 30]*C.uint8_t)(cSigs)[:count:count]
	sigLenArray := (*[1 << 30]C.size_t)(cSigLens)[:count:count]

	for i := 0; i < count; i++ {

		if len(privateKeys[i]) == 0 || len(messages[i]) == 0 {
			return nil, errors.New("invalid empty input in batch")
		}

		signatures[i] = make([]byte, sigMax)

		msgPtrArray[i] = (*C.uint8_t)(unsafe.Pointer(&messages[i][0]))
		msgLenArray[i] = C.size_t(len(messages[i]))
		privPtrArray[i] = (*C.uint8_t)(unsafe.Pointer(&privateKeys[i][0]))
		sigPtrArray[i] = (*C.uint8_t)(unsafe.Pointer(&signatures[i][0]))
	}

	atomic.AddUint64(&cgoCallCount, 1)

	res := C.batch_sign(
		d.alg,
		(**C.uint8_t)(cMessages),
		(*C.size_t)(cMsgLens),
		(**C.uint8_t)(cPrivKeys),
		(**C.uint8_t)(cSigs),
		(*C.size_t)(cSigLens),
		C.size_t(count),
	)

	if res != 0 {
		return nil, errors.New("batch sign failed")
	}

	for i := 0; i < count; i++ {
		signatures[i] = signatures[i][:sigLenArray[i]]
	}

	return signatures, nil
}

func (d *DilithiumSigner) Verify(publicKey []byte, message []byte, signature []byte) bool {

	if d.alg == nil {
		return false
	}

	if len(publicKey) == 0 || len(message) == 0 || len(signature) == 0 {
		return false
	}

	atomic.AddUint64(&cgoCallCount, 1)
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

func (d *DilithiumSigner) Algorithm() string {
	return "dilithium2"
}

func (d *DilithiumSigner) Close() {
	if d.alg != nil {
		atomic.AddUint64(&cgoCallCount, 1)
		C.OQS_SIG_free(d.alg)
		d.alg = nil
	}
}
