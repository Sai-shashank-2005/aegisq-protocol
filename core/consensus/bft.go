package consensus

import (
	"errors"
	"sync"
)

/*
Layer 9 — Hardened BFT Voting Engine

Implements:

✔ PREPARE phase
✔ COMMIT phase
✔ 2f+1 quorum logic
✔ Double-vote prevention
✔ Equivocation prevention (strict)
✔ View-aware voting
*/

type VoteType int

const (
	Prepare VoteType = iota
	Commit
)

type Vote struct {
	ValidatorID string
	BlockHash   string
	View        int
	Type        VoteType
}

type VotePool struct {
	mu sync.Mutex

	// blockHash -> view -> voteType -> validatorID -> bool
	votes map[string]map[int]map[VoteType]map[string]bool

	// NEW: strict equivocation tracking
	// view -> voteType -> validatorID -> blockHash
	seenVotes map[int]map[VoteType]map[string]string

	validatorSet *ValidatorSet
}

func NewVotePool(vs *ValidatorSet) *VotePool {
	return &VotePool{
		votes:        make(map[string]map[int]map[VoteType]map[string]bool),
		seenVotes:    make(map[int]map[VoteType]map[string]string),
		validatorSet: vs,
	}
}

/*
AddVote registers a vote if valid.

Enforces:
✔ Validator must be authorized
✔ No double voting (same block)
✔ No equivocation (different block same view)
*/
func (vp *VotePool) AddVote(v Vote) error {
	vp.mu.Lock()
	defer vp.mu.Unlock()

	// 1️⃣ Authorization check
	if _, exists := vp.validatorSet.GetValidator(v.ValidatorID); !exists {
		return errors.New("unauthorized validator")
	}

	// 2️⃣ Initialize seenVotes structure
	if _, ok := vp.seenVotes[v.View]; !ok {
		vp.seenVotes[v.View] = make(map[VoteType]map[string]string)
	}

	if _, ok := vp.seenVotes[v.View][v.Type]; !ok {
		vp.seenVotes[v.View][v.Type] = make(map[string]string)
	}

	// 3️⃣ Equivocation prevention
	if existingHash, voted := vp.seenVotes[v.View][v.Type][v.ValidatorID]; voted {
		if existingHash == v.BlockHash {
			return errors.New("double vote detected")
		}
		return errors.New("equivocation detected")
	}

	// Record seen vote globally
	vp.seenVotes[v.View][v.Type][v.ValidatorID] = v.BlockHash

	// 4️⃣ Store vote per block
	if _, ok := vp.votes[v.BlockHash]; !ok {
		vp.votes[v.BlockHash] = make(map[int]map[VoteType]map[string]bool)
	}

	if _, ok := vp.votes[v.BlockHash][v.View]; !ok {
		vp.votes[v.BlockHash][v.View] = make(map[VoteType]map[string]bool)
	}

	if _, ok := vp.votes[v.BlockHash][v.View][v.Type]; !ok {
		vp.votes[v.BlockHash][v.View][v.Type] = make(map[string]bool)
	}

	vp.votes[v.BlockHash][v.View][v.Type][v.ValidatorID] = true

	return nil
}

/*
HasQuorum checks if 2f+1 votes exist.
*/
func (vp *VotePool) HasQuorum(blockHash string, view int, voteType VoteType) bool {
	vp.mu.Lock()
	defer vp.mu.Unlock()

	n := vp.validatorSet.Count()
	if n == 0 {
		return false
	}

	f := (n - 1) / 3
	quorum := 2*f + 1

	if _, ok := vp.votes[blockHash]; !ok {
		return false
	}

	if _, ok := vp.votes[blockHash][view]; !ok {
		return false
	}

	count := len(vp.votes[blockHash][view][voteType])
	return count >= quorum
}