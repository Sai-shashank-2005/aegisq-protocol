package consensus

import "testing"

func TestFinalityFlow(t *testing.T) {

	vs := NewValidatorSet()
	vs.AddValidator("v1", []byte("k1"))
	vs.AddValidator("v2", []byte("k2"))
	vs.AddValidator("v3", []byte("k3"))
	vs.AddValidator("v4", []byte("k4"))

	vp := NewVotePool(vs)
	fe := NewFinalityEngine(vp)

	hash := "blockA"
	height := 1
	view := 0

	// Prepare votes
	vp.AddVote(Vote{"v1", hash, view, Prepare})
	vp.AddVote(Vote{"v2", hash, view, Prepare})
	vp.AddVote(Vote{"v3", hash, view, Prepare})

	if !fe.TryPrepare(height, hash, view) {
		t.Fatal("Prepare should succeed")
	}

	// Commit votes
	vp.AddVote(Vote{"v1", hash, view, Commit})
	vp.AddVote(Vote{"v2", hash, view, Commit})
	vp.AddVote(Vote{"v3", hash, view, Commit})

	if err := fe.TryCommit(height, hash, view); err != nil {
		t.Fatal("Commit should succeed:", err)
	}

	if !fe.IsFinalized(height, hash) {
		t.Fatal("Block should be finalized")
	}
}

func TestForkPrevention(t *testing.T) {

	vs := NewValidatorSet()
	vs.AddValidator("v1", []byte("k1"))
	vs.AddValidator("v2", []byte("k2"))
	vs.AddValidator("v3", []byte("k3"))
	vs.AddValidator("v4", []byte("k4"))

	vp := NewVotePool(vs)
	fe := NewFinalityEngine(vp)

	hash1 := "blockA"
	hash2 := "blockB"
	height := 1
	view := 0

	// Finalize blockA
	vp.AddVote(Vote{"v1", hash1, view, Prepare})
	vp.AddVote(Vote{"v2", hash1, view, Prepare})
	vp.AddVote(Vote{"v3", hash1, view, Prepare})
	fe.TryPrepare(height, hash1, view)

	vp.AddVote(Vote{"v1", hash1, view, Commit})
	vp.AddVote(Vote{"v2", hash1, view, Commit})
	vp.AddVote(Vote{"v3", hash1, view, Commit})
	fe.TryCommit(height, hash1, view)

	// Try finalizing different block at same height
	err := fe.TryCommit(height, hash2, view)
	if err == nil {
		t.Fatal("Fork should not be allowed")
	}
}