package main

import (
	"fmt"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
)

func main() {
	signer := &crypto.Ed25519Signer{}

	node, err := identity.NewNodeIdentity("validator-1", signer)
	if err != nil {
		panic(err)
	}

	fmt.Println("Node Created:")
	fmt.Println(node)

	message := []byte("AegisQ Protocol Genesis")

	signature, err := node.Sign(message)
	if err != nil {
		panic(err)
	}

	valid := node.Verify(message, signature)

	fmt.Println("Signature Valid:", valid)
}