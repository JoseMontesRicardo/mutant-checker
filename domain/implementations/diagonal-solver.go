package implementations

import (
	"test.com/mutant-checker/domain/bussiness_conditions"
	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/types"
)

// this implements ISequenceSolver
type DiagonalSolver struct {
	AdnSequence       *types.AdnSequence
	SequenceValidator interfaces.ISequenceValidator
	PositionModifier  interfaces.IPositionModifier
}

func (s *DiagonalSolver) GetSecuences(sequencesChannel chan int) int {
	sequences := s.GetDiagonalSecuences(false) + s.GetDiagonalSecuences(true)
	sequencesChannel <- sequences
	return sequences
}

func (s *DiagonalSolver) GetDiagonalSecuences(isReverse bool) int {
	var sequences int = 0
	diagonal := ""
	s.PositionModifier = &PositionModifier{
		LengthOfMatrix: len(*s.AdnSequence),
		IsReverse:      isReverse,
	}
	s.PositionModifier.InitValues()
	s.PositionModifier.SetNumberOfDiagonals(s.GetNumberOfDiagonal())
	s.PositionModifier.SetHalfOfIteration()
	for i := 0; i < s.PositionModifier.GetNumberOfDiagonals(); i++ {
		s.PositionModifier.SetIteration(i)
		diagonal = s.GetOneDiagonalSequence()
		if s.SequenceValidator.ValidateSequencesForEachNitroBase(diagonal) {
			sequences++
			if sequences > 1 {
				break
			}
		}
		s.PositionModifier.SetDiagonalLength()
		s.PositionModifier.SetPositions()
	}
	return sequences
}

/**
Calcula el numero de diagonales de una matriz NxN con mas de 4 secuancias de ADN
formula utilizada (x-3) + [(x-3) - 1] donde x el el tamaño de la matriz.
*/
func (s *DiagonalSolver) GetNumberOfDiagonal() int {
	x := s.PositionModifier.GetLengthOfMatrix()
	diagonalsIgnored := bussiness_conditions.MaxSequence() - 1
	return (x - diagonalsIgnored) + ((x - diagonalsIgnored) - 1)
}

func (s *DiagonalSolver) GetOneDiagonalSequence() string {
	diagonal := ""
	for i := 0; i < s.PositionModifier.GetDiagonalLength(); i++ {
		var row int
		var column int
		if s.PositionModifier.GetIsReverse() {
			row = s.PositionModifier.GetRow() + i
			column = s.PositionModifier.GetColum() - i
		} else {
			row = s.PositionModifier.GetRow() + i
			column = s.PositionModifier.GetColum() + i
		}
		diagonal += string((*s.AdnSequence)[row][column])
	}
	return diagonal
}
