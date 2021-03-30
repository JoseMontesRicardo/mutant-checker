package interfaces

type ISequenceSolver interface {
	GetSecuences(sequencesChannel chan int) int
}
