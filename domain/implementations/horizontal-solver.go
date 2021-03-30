package implementations

import (
	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/types"
)

// this implements ISequenceSolver
type HorizontalSolver struct {
	AdnSequence       *types.AdnSequence
	SequenceValidator interfaces.ISequenceValidator
}

func (s *HorizontalSolver) GetSecuences(sequencesChannel chan int) int {
	var sequences int = 0
	for _, row := range *s.AdnSequence {
		if s.SequenceValidator.ValidateSequencesForEachNitroBase(row) {
			sequences++
			if sequences > 1 {
				break
			}
		}
	}
	sequencesChannel <- sequences
	return sequences
}
