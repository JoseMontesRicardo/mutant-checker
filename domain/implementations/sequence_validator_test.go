package implementations

import (
	"testing"
)

func TestGetEqualSequence(t *testing.T) {
	expected := "AAAA"
	nitrogenousBase := "A"
	sequenceValidator := SequenceValidator{}
	toCheck := sequenceValidator.GetEqualSequence(nitrogenousBase)
	if toCheck != expected {
		t.Errorf("GetEqualSequence failed")
	}
}

func TestValidateSequenceInRow(t *testing.T) {
	expected := true
	nitrogenousBase := "A"
	row := "TAAAAT"
	sequenceValidator := SequenceValidator{}
	toCheck := sequenceValidator.ValidateSequenceInRow(row, nitrogenousBase)
	if toCheck != expected {
		t.Errorf("ValidateSequenceInRow failed")
	}
}

func TestValidateSequencesForEachNitroBase(t *testing.T) {
	expected := true
	row := "TAAAAT"
	sequenceValidator := SequenceValidator{}
	toCheck := sequenceValidator.ValidateSequencesForEachNitroBase(row)
	if toCheck != expected {
		t.Errorf("ValidateSequencesForEachNitroBase failed")
	}
}
