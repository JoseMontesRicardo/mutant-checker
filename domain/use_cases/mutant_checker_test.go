package use_cases

import (
	"testing"

	"test.com/mutant-checker/domain/implementations"
	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/types"
)

func TestIsMutant(t *testing.T) {
	expected := true
	dna := types.AdnSequence{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	var verticalSolver interfaces.ISequenceSolver = &implementations.VerticalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &implementations.SequenceValidator{},
	}
	var horizontalSolver interfaces.ISequenceSolver = &implementations.HorizontalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &implementations.SequenceValidator{},
	}
	var diagonalSolver interfaces.ISequenceSolver = &implementations.DiagonalSolver{
		AdnSequence:       &dna,
		SequenceValidator: &implementations.SequenceValidator{},
		PositionModifier:  &implementations.PositionModifier{},
	}
	var mutantChecker interfaces.IMutantChecker = &MutantChecker{
		HorizontalSolver: horizontalSolver,
		VerticalSolver:   verticalSolver,
		DiagonalSolver:   diagonalSolver,
	}
	toCheck := mutantChecker.IsMutant(dna)
	if toCheck != expected {
		t.Errorf("toCheck failed")
	}
}
