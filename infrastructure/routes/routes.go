package routes

import (
	"github.com/gin-gonic/gin"
	"test.com/mutant-checker/infrastructure/controllers/health"
	"test.com/mutant-checker/infrastructure/controllers/mutant"
	"test.com/mutant-checker/infrastructure/controllers/stat"
	"test.com/mutant-checker/infrastructure/middlewares"
)

func SetRoutes(routes *gin.Engine) {
	routes.GET("/health", health.CheckHealth)

	apiV1 := routes.Group("/v1")
	{
		apiV1.GET("/stats",
			stat.GetStat,
		)
		apiV1.POST("/mutant",
			middlewares.ValidateRequest,
			mutant.IsMutant,
		)
	}
}
