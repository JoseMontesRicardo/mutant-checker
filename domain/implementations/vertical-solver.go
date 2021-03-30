package implementations

import (
	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/types"
)

// this implements ISequenceSolver
type VerticalSolver struct {
	AdnSequence       *types.AdnSequence
	SequenceValidator interfaces.ISequenceValidator
}

func (s *VerticalSolver) GetSecuences(sequencesChannel chan int) int {
	var sequences int = 0
	row := ""
	for index := range *s.AdnSequence {
		row = s.GetColumnAsRow(index)
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

func (s *VerticalSolver) GetColumnAsRow(index int) string {
	column := ""
	for _, row := range *s.AdnSequence {
		column += string(row[index])
	}
	return column
}
