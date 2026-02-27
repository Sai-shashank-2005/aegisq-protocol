package ledger

import (
	"errors"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
)

type Ledger struct {
	Blocks []*block.Block
}

// Create new ledger with genesis block
func NewLedger(genesis *block.Block) *Ledger {
	return &Ledger{
		Blocks: []*block.Block{genesis},
	}
}

// GetLastBlock returns latest block
func (l *Ledger) GetLastBlock() *block.Block {
	return l.Blocks[len(l.Blocks)-1]
}

// AddBlock validates and appends block
func (l *Ledger) AddBlock(b *block.Block, signer crypto.Signer, validatorPubKey []byte) error {

	last := l.GetLastBlock()

	// 1️⃣ Enforce sequential index
	if b.Index != last.Index+1 {
		return errors.New("invalid block index")
	}

	// 2️⃣ Enforce previous hash linkage
	if string(b.PreviousHash) != string(last.Hash) {
		return errors.New("invalid previous hash linkage")
	}

	// 3️⃣ Verify block integrity
	valid, err := b.Verify(signer, validatorPubKey)
	if err != nil || !valid {
		return errors.New("block verification failed")
	}

	// 4️⃣ Prevent duplicate hash
	for _, existing := range l.Blocks {
		if string(existing.Hash) == string(b.Hash) {
			return errors.New("duplicate block detected")
		}
	}

	l.Blocks = append(l.Blocks, b)
	return nil
}

// ValidateChain performs full chain validation
func (l *Ledger) ValidateChain(signer crypto.Signer, validatorPubKey []byte) error {

	for i := 1; i < len(l.Blocks); i++ {

		current := l.Blocks[i]
		prev := l.Blocks[i-1]

		// Check index
		if current.Index != prev.Index+1 {
			return errors.New("chain index broken")
		}

		// Check previous hash linkage
		if string(current.PreviousHash) != string(prev.Hash) {
			return errors.New("chain previous hash broken")
		}

		// Verify block cryptographically
		valid, err := current.Verify(signer, validatorPubKey)
		if err != nil || !valid {
			return errors.New("block verification failed")
		}
	}

	return nil
}