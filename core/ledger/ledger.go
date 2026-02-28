package ledger

import (
	"errors"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/scheduler"
)

/*
Ledger represents the blockchain state container.

Maintains:
✔ Ordered list of blocks
✔ Validator authorization set
✔ Deterministic leader scheduler (Layer 7)

Security Guarantees:
✔ Sequential block ordering
✔ Hash chain integrity
✔ Cryptographic block verification
✔ Validator authorization enforcement
✔ Deterministic leader enforcement
*/
type Ledger struct {
	Blocks       []*block.Block
	ValidatorSet *consensus.ValidatorSet
	Scheduler    *scheduler.RoundRobinScheduler
}

/*
NewLedger initializes ledger with:

- Genesis block
- Validator set
- Round-robin scheduler
*/
func NewLedger(genesis *block.Block, vs *consensus.ValidatorSet) *Ledger {

	s := scheduler.NewRoundRobinScheduler(vs)

	return &Ledger{
		Blocks:       []*block.Block{genesis},
		ValidatorSet: vs,
		Scheduler:    s,
	}
}

/*
GetLastBlock returns the most recent block.
*/
func (l *Ledger) GetLastBlock() *block.Block {
	return l.Blocks[len(l.Blocks)-1]
}

/*
AddBlock performs strict validation before appending.

Validation Steps:

1️⃣ Enforce sequential index
2️⃣ Enforce previous hash linkage
3️⃣ Enforce deterministic leader schedule (Layer 7)
4️⃣ Enforce validator authorization (Layer 6)
5️⃣ Verify block cryptographic integrity
6️⃣ Prevent duplicate block insertion
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

	// 3️⃣ Enforce deterministic leader schedule
	expectedLeader, err := l.Scheduler.GetLeader(b.Index)
	if err != nil {
		return err
	}

	if b.Validator != expectedLeader {
		return errors.New("block produced by wrong scheduled validator")
	}

	// 4️⃣ Enforce validator authorization
	if !l.ValidatorSet.IsAuthorized(b.Validator, validatorPubKey) {
		return errors.New("validator not authorized")
	}

	// 5️⃣ Verify cryptographic integrity
	valid, err := b.Verify(signer, validatorPubKey)
	if err != nil || !valid {
		return errors.New("block verification failed")
	}

	// 6️⃣ Prevent duplicate block hash insertion
	for _, existing := range l.Blocks {
		if string(existing.Hash) == string(b.Hash) {
			return errors.New("duplicate block detected")
		}
	}

	// Append block
	l.Blocks = append(l.Blocks, b)

	return nil
}

/*
ValidateChain performs full-chain validation.

Checks:

✔ Index continuity
✔ Previous hash linkage
✔ Deterministic leader correctness
✔ Validator authorization
✔ Cryptographic integrity
*/
func (l *Ledger) ValidateChain(
	signer crypto.Signer,
) error {

	for i := 1; i < len(l.Blocks); i++ {

		current := l.Blocks[i]
		prev := l.Blocks[i-1]

		// Index continuity
		if current.Index != prev.Index+1 {
			return errors.New("chain index broken")
		}

		// Previous hash linkage
		if string(current.PreviousHash) != string(prev.Hash) {
			return errors.New("chain previous hash broken")
		}

		// Deterministic leader enforcement
		expectedLeader, err := l.Scheduler.GetLeader(current.Index)
		if err != nil {
			return err
		}

		if current.Validator != expectedLeader {
			return errors.New("invalid leader at block height")
		}

		// Validator must exist
		validatorKey, exists := l.ValidatorSet.GetValidator(current.Validator)
		if !exists {
			return errors.New("block signed by unknown validator")
		}

		// Cryptographic verification
		valid, err := current.Verify(signer, validatorKey)
		if err != nil || !valid {
			return errors.New("block verification failed")
		}
	}

	return nil
}