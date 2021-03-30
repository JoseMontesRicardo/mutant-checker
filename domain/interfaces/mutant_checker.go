package interfaces

import "test.com/mutant-checker/domain/types"

type IMutantChecker interface {
	IsMutant(dna types.AdnSequence) bool
}
