package interfaces

type ISequenceValidator interface {
	ValidateSequencesForEachNitroBase(row string) bool
	ValidateSequenceInRow(row string, nitrogenousBase string) bool
	GetEqualSequence(nitrogenousBase string) string
}
