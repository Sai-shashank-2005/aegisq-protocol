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

	tests := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "A",
		4: "B",
		5: "C",
	}

	for index, expected := range tests {

		leader, err := s.GetLeader(index)
		if err != nil {
			t.Fatal(err)
		}

		if leader != expected {
			t.Fatalf("expected %s, got %s", expected, leader)
		}
	}
}