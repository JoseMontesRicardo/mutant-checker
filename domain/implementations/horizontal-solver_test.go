package implementations

import (
	"testing"

	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/types"
)

func TestHGetSecuences(t *testing.T) {
	dna := types.AdnSequence{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	testChann := make(chan int)
	var solver interfaces.ISequenceSolver = &HorizontalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &SequenceValidator{},
	}
	expected := 1 // get sequence
	go solver.GetSecuences(testChann)
	toCheck := <-testChann
	if toCheck != expected {
		t.Errorf("GetSecuences Horizontal failed, got: %d, want: %d.", toCheck, expected)
	}
}
