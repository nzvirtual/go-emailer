package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nzvirtual/go-emailer/emails"
)

func SetupRoutes(engine *gin.Engine) {
	engine.POST("/send/registration", emails.SendRegistration)
}
