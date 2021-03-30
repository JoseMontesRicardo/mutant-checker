package implementations

import (
	"testing"

	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/types"
)

func TestVGetSecuences(t *testing.T) {
	dna := types.AdnSequence{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	testChann := make(chan int)
	var solver interfaces.ISequenceSolver = &VerticalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &SequenceValidator{},
	}
	expected := 1 // get sequence
	go solver.GetSecuences(testChann)
	toCheck := <-testChann
	if toCheck != expected {
		t.Errorf("GetSecuences Vertical failed, got: %d, want: %d.", toCheck, expected)
	}
}

func TestGetColumnAsRow(t *testing.T) {
	dna := types.AdnSequence{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	solver := &VerticalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &SequenceValidator{},
	}
	expected := "ACTACT" // get sequence
	toCheck := solver.GetColumnAsRow(0)
	if toCheck != expected {
		t.Errorf("GetSecuences GetColumnAsRow failed")
	}
}
