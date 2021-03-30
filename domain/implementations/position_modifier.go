package implementations

import (
	"test.com/mutant-checker/domain/bussiness_conditions"
)

type PositionModifier struct {
	Row               int
	Column            int
	Iteration         int
	LengthOfMatrix    int
	HalfOfIteration   int
	IsReverse         bool
	DiagonalLength    int
	NumberOfDiagonals int
}

func (pm *PositionModifier) InitValues() {
	pm.DiagonalLength = bussiness_conditions.MaxSequence()
	diagonalsIgnored := bussiness_conditions.MaxSequence() - 1
	pm.Row = (pm.LengthOfMatrix - diagonalsIgnored) - 1
	if pm.IsReverse {
		pm.Column = pm.LengthOfMatrix - 1
	} else {
		pm.Column = 0
	}
}

func (pm *PositionModifier) SetDiagonalLength() {
	if pm.HalfOfIteration > pm.Iteration {
		pm.DiagonalLength++
	} else {
		pm.DiagonalLength--
	}
}

func (pm *PositionModifier) SetPositions() {
	if pm.IsReverse {
		if pm.HalfOfIteration > pm.Iteration {
			pm.Row--
		} else {
			pm.Column--
		}
	} else {
		if pm.HalfOfIteration > pm.Iteration {
			pm.Row--
		} else {
			pm.Column++
		}
	}
}

func (pm *PositionModifier) SetNumberOfDiagonals(value int) {
	pm.NumberOfDiagonals = value
}

func (pm *PositionModifier) SetIteration(iteration int) {
	pm.Iteration = iteration
}

func (pm *PositionModifier) SetHalfOfIteration() {
	pm.HalfOfIteration = pm.NumberOfDiagonals / 2
}

func (pm *PositionModifier) GetLengthOfMatrix() int {
	return pm.LengthOfMatrix
}

func (pm *PositionModifier) GetDiagonalLength() int {
	return pm.DiagonalLength
}

func (pm *PositionModifier) GetRow() int {
	return pm.Row
}

func (pm *PositionModifier) GetColum() int {
	return pm.Column
}

func (pm *PositionModifier) GetNumberOfDiagonals() int {
	return pm.NumberOfDiagonals
}

func (pm *PositionModifier) GetIsReverse() bool {
	return pm.IsReverse
}
