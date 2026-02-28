package consensus

import (
	"errors"
	"sync"
)

/*
Layer 10 — Finality Engine

Enforces:

✔ Prepare → Commit transition
✔ 2f+1 commit required
✔ Single finalization per height
✔ No fork after finality
*/

type FinalityEngine struct {
	mu sync.Mutex

	votePool *VotePool

	// height -> blockHash finalized
	finalized map[int]string

	// blockHash -> prepared
	prepared map[string]bool
}

func NewFinalityEngine(vp *VotePool) *FinalityEngine {
	return &FinalityEngine{
		votePool: vp,
		finalized: make(map[int]string),
		prepared:  make(map[string]bool),
	}
}

/*
TryPrepare checks if prepare quorum reached.
*/
func (fe *FinalityEngine) TryPrepare(height int, hash string, view int) bool {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	if fe.votePool.HasQuorum(hash, view, Prepare) {
		fe.prepared[hash] = true
		return true
	}

	return false
}

/*
TryCommit checks if commit quorum reached and finalizes block.
*/
func (fe *FinalityEngine) TryCommit(height int, hash string, view int) error {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	// Must be prepared first
	if !fe.prepared[hash] {
		return errors.New("block not prepared")
	}

	if !fe.votePool.HasQuorum(hash, view, Commit) {
		return errors.New("commit quorum not reached")
	}

	// Prevent fork at same height
	if _, exists := fe.finalized[height]; exists {
		return errors.New("height already finalized")
	}

	fe.finalized[height] = hash
	return nil
}

/*
IsFinalized checks if block is finalized.
*/
func (fe *FinalityEngine) IsFinalized(height int, hash string) bool {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	finalHash, exists := fe.finalized[height]
	return exists && finalHash == hash
}