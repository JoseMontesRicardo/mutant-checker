package implementations

import (
	"strings"

	"test.com/mutant-checker/domain/bussiness_conditions"
	"test.com/mutant-checker/domain/constants"
)

type SequenceValidator struct{}

func (sv *SequenceValidator) ValidateSequencesForEachNitroBase(row string) bool {
	return sv.ValidateSequenceInRow(row, constants.NITROGENOUS_BASE_A) ||
		sv.ValidateSequenceInRow(row, constants.NITROGENOUS_BASE_C) ||
		sv.ValidateSequenceInRow(row, constants.NITROGENOUS_BASE_G) ||
		sv.ValidateSequenceInRow(row, constants.NITROGENOUS_BASE_T)
}

func (sv *SequenceValidator) ValidateSequenceInRow(row string, nitrogenousBase string) bool {
	return strings.Contains(row, sv.GetEqualSequence(nitrogenousBase))
}

func (sv *SequenceValidator) GetEqualSequence(nitrogenousBase string) string {
	return strings.Repeat(nitrogenousBase, bussiness_conditions.MaxSequence())
}
