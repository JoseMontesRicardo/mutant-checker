package repositories

import (
	"test.com/mutant-checker/domain/models"
	"test.com/mutant-checker/domain/types"
)

type MutantRepo interface {
	Insert(doc models.Mutant) (id string)
	GetStats() types.Stats
}
