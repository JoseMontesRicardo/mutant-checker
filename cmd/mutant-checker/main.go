package main

import (
	"github.com/gin-gonic/gin"
	"test.com/mutant-checker/infrastructure/db"
	"test.com/mutant-checker/infrastructure/server"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	db.Initialize()
	server.StartServer()
}
