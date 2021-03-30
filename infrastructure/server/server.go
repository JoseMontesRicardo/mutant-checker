package server

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"test.com/mutant-checker/infrastructure/routes"
)

type Server struct {
	Engine *gin.Engine
}

func StartServer() *Server {
	server := &Server{
		Engine: gin.Default(),
	}
	server.SetConfig()
	routes.SetRoutes(server.Engine)
	server.Listen()
	return server
}

func (pServer *Server) SetConfig() {
	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(os.Getenv("ALLOW_ORIGINS"), ";")
	config.AllowCredentials = true
	config.AllowHeaders = []string{"content-type", "captcha"}
	pServer.Engine.Use(cors.New(config))
}

func (pServer *Server) Listen() {
	pServer.Engine.Run(os.Getenv("API_URL"))
}
