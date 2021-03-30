package interfaces

type IPositionModifier interface {
	InitValues()
	SetDiagonalLength()
	SetNumberOfDiagonals(value int)
	SetHalfOfIteration()
	SetPositions()
	SetIteration(iteration int)
	GetLengthOfMatrix() int
	GetDiagonalLength() int
	GetRow() int
	GetColum() int
	GetNumberOfDiagonals() int
	GetIsReverse() bool
}
