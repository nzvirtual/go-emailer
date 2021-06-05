package main

import (
	"net/http"
	"strings"

	"github.com/dhawton/log4g"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(appenv string) *Server {
	server := Server{}

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(Logger)
	engine.Use(func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		log4g.Category("middleware/auth").Debug("Authorization header: " + auth)
		authPieces := strings.Split(auth, " ")
		if authPieces[1] != ApiKey {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			c.Abort()
			return
		}
	})
	server.engine = engine

	SetupRoutes(engine)

	return &server
}
