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

	ids := vs.GetValidatorIDs()
	sort.Strings(ids)

	return &RoundRobinScheduler{
		orderedValidators: ids,
	}
}

func (r *RoundRobinScheduler) GetLeader(height int, view int) (string, error) {

	if len(r.orderedValidators) == 0 {
		return "", errors.New("no validators registered")
	}

	pos := (height + view) % len(r.orderedValidators)

	return r.orderedValidators[pos], nil
}