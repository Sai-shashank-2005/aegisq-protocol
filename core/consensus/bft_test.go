package consensus

import "testing"

func setupValidators() *ValidatorSet {
	vs := NewValidatorSet()

	vs.AddValidator("v1", []byte("k1"))
	vs.AddValidator("v2", []byte("k2"))
	vs.AddValidator("v3", []byte("k3"))
	vs.AddValidator("v4", []byte("k4"))

	return vs
}

func TestPrepareQuorum(t *testing.T) {
	vs := setupValidators()
	vp := NewVotePool(vs)

	hash := "block123"

	vp.AddVote(Vote{"v1", hash, 0, Prepare})
	vp.AddVote(Vote{"v2", hash, 0, Prepare})
	vp.AddVote(Vote{"v3", hash, 0, Prepare})

	if !vp.HasQuorum(hash, 0, Prepare) {
		t.Fatal("Should reach prepare quorum")
	}
}

func TestCommitQuorum(t *testing.T) {
	vs := setupValidators()
	vp := NewVotePool(vs)

	hash := "block123"

	vp.AddVote(Vote{"v1", hash, 0, Commit})
	vp.AddVote(Vote{"v2", hash, 0, Commit})
	vp.AddVote(Vote{"v3", hash, 0, Commit})

	if !vp.HasQuorum(hash, 0, Commit) {
		t.Fatal("Should reach commit quorum")
	}
}

func TestDoubleVoteRejected(t *testing.T) {
	vs := setupValidators()
	vp := NewVotePool(vs)

	hash := "block123"

	err := vp.AddVote(Vote{"v1", hash, 0, Prepare})
	if err != nil {
		t.Fatal(err)
	}

	err = vp.AddVote(Vote{"v1", hash, 0, Prepare})
	if err == nil {
		t.Fatal("Double vote should be rejected")
	}
}

func TestUnauthorizedValidatorRejected(t *testing.T) {
	vs := setupValidators()
	vp := NewVotePool(vs)

	err := vp.AddVote(Vote{"evil", "block123", 0, Prepare})
	if err == nil {
		t.Fatal("Unauthorized validator should be rejected")
	}
}