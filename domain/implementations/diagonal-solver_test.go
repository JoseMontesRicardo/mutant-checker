package implementations

import (
	"testing"

	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/types"
)

func TestGetSecuences(t *testing.T) {
	dna := types.AdnSequence{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	testChann := make(chan int)
	var solver interfaces.ISequenceSolver = &DiagonalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &SequenceValidator{},
		PositionModifier:  &PositionModifier{},
	}
	expected := 1 // get sequence
	go solver.GetSecuences(testChann)
	toCheck := <-testChann
	if toCheck != expected {
		t.Errorf("GetSecuences failed, got: %d, want: %d.", toCheck, expected)
	}
}

func TestGetDiagonalSecuences(t *testing.T) {
	dna := types.AdnSequence{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	solver := &DiagonalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &SequenceValidator{},
		PositionModifier:  &PositionModifier{},
	}
	expected1 := 1 // get sequence
	expected2 := 0 // does not get sequence
	toCheck1 := solver.GetDiagonalSecuences(false)
	toCheck2 := solver.GetDiagonalSecuences(true)
	if toCheck1 != expected1 {
		t.Errorf("GetDiagonalSecuences failed, got: %d, want: %d.", toCheck1, expected1)
	}
	if toCheck2 != expected2 {
		t.Errorf("GetDiagonalSecuences failed, got: %d, want: %d.", toCheck2, expected2)
	}
}

func TestGetNumberOfDiagonal(t *testing.T) {
	dna := types.AdnSequence{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	solver := &DiagonalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &SequenceValidator{},
		PositionModifier: &PositionModifier{
			LengthOfMatrix: 10,
		},
	}
	expected := 13
	toCheck := solver.GetNumberOfDiagonal()
	if toCheck != expected {
		t.Errorf("GetDiagonalSecuences failed, got: %d, want: %d.", toCheck, expected)
	}
}

func TestGetOneDiagonalSequence(t *testing.T) {
	dna := types.AdnSequence{ // test with 10X10 matrix
		"ATGCGATTTT",
		"CAGTGCGTCA",
		"TTATGTGTCA",
		"AGAAGGGTCA",
		"CCCCTAGTCA",
		"TCACTCGTCA",
		"TCACTCGTCA",
		"TCACTCGTCA",
		"TCACTCGTCA",
		"TCACTCGTCA",
	}
	solver := &DiagonalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &SequenceValidator{},
		PositionModifier: &PositionModifier{
			IsReverse:      false,
			LengthOfMatrix: 10,
			DiagonalLength: 10,
			Row:            0,
			Column:         0,
		},
	}
	expected := "AAAATCGTCA"
	toCheck := solver.GetOneDiagonalSequence()
	if toCheck != expected {
		t.Errorf("toCheck failed")
	}
}
