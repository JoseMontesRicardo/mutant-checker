package use_cases

import (
	"test.com/mutant-checker/domain/bussiness_conditions"
	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/types"
)

type MutantChecker struct {
	HorizontalSolver interfaces.ISequenceSolver
	VerticalSolver   interfaces.ISequenceSolver
	DiagonalSolver   interfaces.ISequenceSolver
}

func (mck *MutantChecker) IsMutant(dna types.AdnSequence) bool {
	totalSequencesFound := 0
	sequencesChannel := make(chan int)
	go mck.HorizontalSolver.GetSecuences(sequencesChannel)
	go mck.VerticalSolver.GetSecuences(sequencesChannel)
	go mck.DiagonalSolver.GetSecuences(sequencesChannel)
	totalSequencesFound += <-sequencesChannel
	totalSequencesFound += <-sequencesChannel
	totalSequencesFound += <-sequencesChannel
	return totalSequencesFound > bussiness_conditions.MinimumSequencesForMutant()
}
