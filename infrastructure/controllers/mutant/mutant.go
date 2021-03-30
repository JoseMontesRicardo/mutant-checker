package mutant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"test.com/mutant-checker/domain/implementations"
	"test.com/mutant-checker/domain/interfaces"
	"test.com/mutant-checker/domain/models"
	"test.com/mutant-checker/domain/repositories"
	"test.com/mutant-checker/domain/types"
	"test.com/mutant-checker/domain/use_cases"
	"test.com/mutant-checker/infrastructure/repositories_impl"
)

func IsMutant(c *gin.Context) {
	var body types.MutantBody
	var repo repositories.MutantRepo = &repositories_impl.MutantRepo
	response := types.MutantResponse{Status: http.StatusOK}
	c.ShouldBindBodyWith(&body, binding.JSON)
	var verticalSolver interfaces.ISequenceSolver = &implementations.VerticalSolver{
		AdnSequence:       &body.Dna,
		SequenceValidator: &implementations.SequenceValidator{},
	}
	var horizontalSolver interfaces.ISequenceSolver = &implementations.HorizontalSolver{
		AdnSequence:       &body.Dna,
		SequenceValidator: &implementations.SequenceValidator{},
	}
	var diagonalSolver interfaces.ISequenceSolver = &implementations.DiagonalSolver{
		AdnSequence:       &body.Dna,
		SequenceValidator: &implementations.SequenceValidator{},
		PositionModifier:  &implementations.PositionModifier{},
	}
	var mutantChecker interfaces.IMutantChecker = &use_cases.MutantChecker{
		HorizontalSolver: horizontalSolver,
		VerticalSolver:   verticalSolver,
		DiagonalSolver:   diagonalSolver,
	}
	isMutant := mutantChecker.IsMutant(body.Dna)
	mutant := models.Mutant{IsMutant: isMutant}
	repo.Insert(mutant)
	if isMutant {
		c.JSON(http.StatusOK, response)
	} else {
		response.Status = http.StatusForbidden
		c.JSON(http.StatusForbidden, response)
	}
}
