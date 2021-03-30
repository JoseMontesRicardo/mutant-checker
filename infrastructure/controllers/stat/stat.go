package stat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/mutant-checker/domain/repositories"
	"test.com/mutant-checker/infrastructure/repositories_impl"
)

func GetStat(c *gin.Context) {
	var repo repositories.MutantRepo = &repositories_impl.MutantRepo
	response := repo.GetStats()
	c.JSON(http.StatusOK, response)
}
