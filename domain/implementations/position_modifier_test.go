package implementations

import (
	"testing"
)

func TestInitValues(t *testing.T) {
	positionModifier := PositionModifier{}
	positionModifier.InitValues()
	if positionModifier.Column != 0 {
		t.Errorf("InitValues failed, got: %d, want: %d.", positionModifier.Column, 0)
	}
}

func TestSetDiagonalLength(t *testing.T) {
	expected := 1 // DiagonalLength + 1
	positionModifier := PositionModifier{
		Iteration:       1,
		HalfOfIteration: 2,
		DiagonalLength:  0,
	}
	positionModifier.SetDiagonalLength()
	if positionModifier.DiagonalLength != expected {
		t.Errorf("SetDiagonalLength SetDiagonalLengthIncrement failed, got: %d, want: %d.", positionModifier.DiagonalLength, expected)
	}
}

func TestSetPositions(t *testing.T) {
	expectedRow := 0    // DiagonalLength + 1
	expectedColumn := 1 // DiagonalLength + 1
	positionModifier := PositionModifier{
		Iteration:       1,
		HalfOfIteration: 2,
		IsReverse:       true,
		Row:             1,
		Column:          1,
	}
	positionModifier.SetPositions()
	if positionModifier.Row != expectedRow {
		t.Errorf("SetPositions failed, got: %d, want: %d.", positionModifier.Row, expectedRow)
	}
	if positionModifier.Column != expectedColumn {
		t.Errorf("SetPositions failed, got: %d, want: %d.", positionModifier.Column, expectedColumn)
	}
}

func TestSetNumberOfDiagonals(t *testing.T) {
	expected := 4
	positionModifier := PositionModifier{}
	positionModifier.SetNumberOfDiagonals(4)
	if positionModifier.NumberOfDiagonals != expected {
		t.Errorf("TestSetNumberOfDiagonals failed, got: %d, want: %d.", positionModifier.NumberOfDiagonals, expected)
	}
}

func TestSetIteration(t *testing.T) {
	expected := 4
	positionModifier := PositionModifier{}
	positionModifier.SetIteration(4)
	if positionModifier.Iteration != expected {
		t.Errorf("TestSetNumberOfDiagonals failed, got: %d, want: %d.", positionModifier.Iteration, expected)
	}
}

func TestSetHalfOfIteration(t *testing.T) {
	expected := 2
	positionModifier := PositionModifier{
		NumberOfDiagonals: 5,
	}
	positionModifier.SetHalfOfIteration()
	if positionModifier.HalfOfIteration != expected {
		t.Errorf("SetHalfOfIteration failed, got: %d, want: %d.", positionModifier.HalfOfIteration, expected)
	}
}

func TestGetDiagonalLength(t *testing.T) {
	expected := 5
	positionModifier := PositionModifier{
		DiagonalLength: 5,
	}
	toCheck := positionModifier.GetDiagonalLength()
	if toCheck != expected {
		t.Errorf("GetDiagonalLength failed, got: %d, want: %d.", toCheck, expected)
	}
}

func TestGetRow(t *testing.T) {
	expected := 5
	positionModifier := PositionModifier{
		Row: 5,
	}
	toCheck := positionModifier.GetRow()
	if toCheck != expected {
		t.Errorf("GetRow failed, got: %d, want: %d.", toCheck, expected)
	}
}

func TestGetColum(t *testing.T) {
	expected := 5
	positionModifier := PositionModifier{
		Column: 5,
	}
	toCheck := positionModifier.GetColum()
	if toCheck != expected {
		t.Errorf("GetColum failed, got: %d, want: %d.", toCheck, expected)
	}
}

func TestGetNumberOfDiagonals(t *testing.T) {
	expected := 5
	positionModifier := PositionModifier{
		NumberOfDiagonals: 5,
	}
	toCheck := positionModifier.GetNumberOfDiagonals()
	if toCheck != expected {
		t.Errorf("GetNumberOfDiagonals failed, got: %d, want: %d.", toCheck, expected)
	}
}

func TestGetIsReverse(t *testing.T) {
	expected := true
	positionModifier := PositionModifier{
		IsReverse: true,
	}
	toCheck := positionModifier.GetIsReverse()
	if toCheck != expected {
		t.Errorf("GetIsReverse failed")
	}
}
