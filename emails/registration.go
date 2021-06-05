package emails

import (
	"bytes"
	"html/template"
	"net/http"

	log "github.com/dhawton/log4g"
	"github.com/gin-gonic/gin"
	"github.com/nzvirtual/go-emailer/lib/emailer"
)

type SendRegistrationDTO struct {
	To              string `json:"to" form:"to" binding:"required"`
	Name            string `json:"name" form:"name" binding:"required"`
	TempPassword    string `json:"temppassword" form:"temppassword" binding:"required"`
	VerificationURL string `json:"verificationURL" form:"verificationURL" binding:"required"`
}

func SendRegistration(c *gin.Context) {
	data := SendRegistrationDTO{}
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Category("send/registration").Error("Invalid email response received " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	var body bytes.Buffer

	tmpl := template.Must(template.ParseFiles("templates/registration.html"))
	if err := tmpl.Execute(&body, data); err != nil {
		log.Category("send/registration").Error("Failed to execute template: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	go emailer.SendEmail(data.To, "Welcome to NZVirtual!", body.String())
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
