package scheduler

import (
	"errors"
	"sort"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
)

type RoundRobinScheduler struct {
	orderedValidators []string
}

func NewRoundRobinScheduler(vs *consensus.ValidatorSet) *RoundRobinScheduler {

	// Extract NodeIDs deterministically
	ids := vs.GetValidatorIDs()

	sort.Strings(ids)

	return &RoundRobinScheduler{
		orderedValidators: ids,
	}
}

// GetLeader returns expected leader for block height
func (r *RoundRobinScheduler) GetLeader(blockIndex int) (string, error) {

	if len(r.orderedValidators) == 0 {
		return "", errors.New("no validators registered")
	}

	pos := blockIndex % len(r.orderedValidators)

	return r.orderedValidators[pos], nil
}