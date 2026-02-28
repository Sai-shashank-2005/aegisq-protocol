package consensus

import (
	"errors"
	"sync"
)

/*
Layer 9 — BFT Voting Engine

Implements:

✔ PREPARE phase
✔ COMMIT phase
✔ 2f+1 quorum logic
✔ Double-vote prevention
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

	validatorSet *ValidatorSet
}

func NewVotePool(vs *ValidatorSet) *VotePool {
	return &VotePool{
		votes:        make(map[string]map[int]map[VoteType]map[string]bool),
		validatorSet: vs,
	}
}

/*
AddVote registers a vote if valid.

Enforces:
✔ Validator must be authorized
✔ No double voting
*/
func (vp *VotePool) AddVote(v Vote) error {
	vp.mu.Lock()
	defer vp.mu.Unlock()

	if _, exists := vp.validatorSet.GetValidator(v.ValidatorID); !exists {
		return errors.New("unauthorized validator")
	}

	if _, ok := vp.votes[v.BlockHash]; !ok {
		vp.votes[v.BlockHash] = make(map[int]map[VoteType]map[string]bool)
	}

	if _, ok := vp.votes[v.BlockHash][v.View]; !ok {
		vp.votes[v.BlockHash][v.View] = make(map[VoteType]map[string]bool)
	}

	if _, ok := vp.votes[v.BlockHash][v.View][v.Type]; !ok {
		vp.votes[v.BlockHash][v.View][v.Type] = make(map[string]bool)
	}

	if vp.votes[v.BlockHash][v.View][v.Type][v.ValidatorID] {
		return errors.New("double vote detected")
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