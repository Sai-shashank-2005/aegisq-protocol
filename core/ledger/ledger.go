package ledger

import (
	"errors"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
)

/*
Ledger represents the blockchain state container.

It maintains:
- Ordered list of blocks
- Validator authorization set (Layer 6 governance enforcement)

Security Guarantees:
✔ Sequential block ordering
✔ Hash chain integrity
✔ Cryptographic block verification
✔ Validator authorization enforcement
*/
type Ledger struct {
	Blocks       []*block.Block
	ValidatorSet *consensus.ValidatorSet
}

/*
NewLedger initializes a new ledger instance.

- Requires a genesis block
- Requires a ValidatorSet (authorization layer)
*/
func NewLedger(genesis *block.Block, vs *consensus.ValidatorSet) *Ledger {
	return &Ledger{
		Blocks:       []*block.Block{genesis},
		ValidatorSet: vs,
	}
}

/*
GetLastBlock returns the most recent block in the chain.
*/
func (l *Ledger) GetLastBlock() *block.Block {
	return l.Blocks[len(l.Blocks)-1]
}

/*
AddBlock performs full validation before appending a block.

Validation Steps:

1️⃣ Enforce sequential index
2️⃣ Enforce previous hash linkage
3️⃣ Verify block cryptographic integrity
4️⃣ Enforce validator authorization (Layer 6)
5️⃣ Prevent duplicate block insertion

Only if ALL checks pass → block is appended.
*/
func (l *Ledger) AddBlock(
	b *block.Block,
	signer crypto.Signer,
	validatorPubKey []byte,
) error {

	last := l.GetLastBlock()

	// 1️⃣ Enforce sequential index
	if b.Index != last.Index+1 {
		return errors.New("invalid block index")
	}

	// 2️⃣ Enforce previous hash linkage
	if string(b.PreviousHash) != string(last.Hash) {
		return errors.New("invalid previous hash linkage")
	}

	// 3️⃣ Verify block cryptographic integrity
	valid, err := b.Verify(signer, validatorPubKey)
	if err != nil || !valid {
		return errors.New("block verification failed")
	}

	// 4️⃣ Enforce validator authorization (NEW LAYER 6)
	// Ensure the validator is registered and public key matches
	if !l.ValidatorSet.IsAuthorized(b.Validator, validatorPubKey) {
		return errors.New("validator not authorized")
	}

	// 5️⃣ Prevent duplicate block hash insertion
	for _, existing := range l.Blocks {
		if string(existing.Hash) == string(b.Hash) {
			return errors.New("duplicate block detected")
		}
	}

	// All validations passed → append block
	l.Blocks = append(l.Blocks, b)

	return nil
}

/*
ValidateChain performs a full chain traversal and validation.

Checks:

✔ Index continuity
✔ Previous hash linkage
✔ Cryptographic block validity
✔ Validator authorization for each block

Used for:
- Node sync
- Startup verification
- Chain audit
*/
func (l *Ledger) ValidateChain(
	signer crypto.Signer,
) error {

	for i := 1; i < len(l.Blocks); i++ {

		current := l.Blocks[i]
		prev := l.Blocks[i-1]

		// Check index continuity
		if current.Index != prev.Index+1 {
			return errors.New("chain index broken")
		}

		// Check previous hash linkage
		if string(current.PreviousHash) != string(prev.Hash) {
			return errors.New("chain previous hash broken")
		}

		// Retrieve registered validator public key
		validatorKey, exists := l.ValidatorSet.GetValidator(current.Validator)
		if !exists {
			return errors.New("block signed by unknown validator")
		}

		// Verify cryptographic integrity using registered key
		valid, err := current.Verify(signer, validatorKey)
		if err != nil || !valid {
			return errors.New("block verification failed")
		}
	}

	return nil
}