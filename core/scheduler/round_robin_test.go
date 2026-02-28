package scheduler

import (
	"testing"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
)

func TestRoundRobinRotation(t *testing.T) {

	vs := consensus.NewValidatorSet()

	vs.AddValidator("A", []byte("a"))
	vs.AddValidator("B", []byte("b"))
	vs.AddValidator("C", []byte("c"))

	s := NewRoundRobinScheduler(vs)

	tests := []struct {
		height   int
		view     int
		expected string
	}{
		{0, 0, "A"},
		{1, 0, "B"},
		{2, 0, "C"},
		{3, 0, "A"},
		{1, 1, "C"}, // failover
		{1, 2, "A"}, // next failover
	}

	for _, test := range tests {

		leader, err := s.GetLeader(test.height, test.view)
		if err != nil {
			t.Fatal(err)
		}

		if leader != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, leader)
		}
	}
}