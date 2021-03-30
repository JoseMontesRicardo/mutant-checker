package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/validator.v2"
	"test.com/mutant-checker/domain/constants"
	"test.com/mutant-checker/domain/types"
)

func ValidateRequest(c *gin.Context) {
	var body types.MutantBody
	response := types.MutantResponse{Status: http.StatusInternalServerError}
	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
	}
	// ValidateStruct(c, &body),
	ValidateStruct(c, &body)
	ValidateLength(c, &body)
	ValidateNitrogenBase(c, &body)
	c.Next()
}

func ValidateStruct(c *gin.Context, mutantBody *types.MutantBody) {
	err := validator.Validate(mutantBody)
	if err != nil {
		json.Marshal(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "sólo se permiten secuencias hasta 10X10")
	}
}

func ValidateLength(c *gin.Context, mutantBody *types.MutantBody) {
	var err error
	numberOfSequences := len(mutantBody.Dna)
	for _, nitrogenBase := range mutantBody.Dna {
		if numberOfSequences != len(nitrogenBase) {
			err = errors.New("la matriz debe ser simetrica (NxN)")
		}
	}
	if err != nil {
		json.Marshal(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
}

func ValidateNitrogenBase(c *gin.Context, mutantBody *types.MutantBody) {
	var err error = nil
	acceptedLeters := constants.NITROGENOUS_BASE_A + constants.NITROGENOUS_BASE_C + constants.NITROGENOUS_BASE_G + constants.NITROGENOUS_BASE_T
	for _, nitrogenBase := range mutantBody.Dna {
		for _, nitrogenBaseLeter := range strings.Split(nitrogenBase, "") {
			if !strings.Contains(acceptedLeters, nitrogenBaseLeter) {
				err = errors.New("letra de base nitrogenada inválida " + nitrogenBaseLeter)
			}
		}
	}
	if err != nil {
		json.Marshal(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
}
