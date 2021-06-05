package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/common-nighthawk/go-figure"
	log "github.com/dhawton/log4g"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nzvirtual/go-emailer/lib/emailer"
)

var ApiKey string

func main() {
	intro := figure.NewFigure("NZV API", "", false).Slicify()
	for i := 0; i < len(intro); i++ {
		log.Category("main").Info(intro[i])
	}

	log.Category("main").Info("Starting NZV Email Service")
	log.Category("main").Info("Checking for .env, loading if exists")
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Category("main").Fatal("Error loading .env file " + err.Error())
		}
	}

	appenv := Getenv("APP_ENV", "dev")
	log.Category("main").Info(fmt.Sprintf("APP_ENV=%s", appenv))

	if appenv == "production" {
		log.SetLogLevel(log.INFO)
		log.Category("main").Info("Setting gin to Release Mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.SetLogLevel(log.DEBUG)
	}

	log.Category("main").Info("Configuring Emailer")
	port, err := strconv.Atoi(Getenv("SMTP_PORT", "25"))
	if err != nil {
		log.Category("main").Info("Could not convert port to int, setting 25")
		port = 25
	}
	enc, err := strconv.Atoi(Getenv("SMTP_TLS", "0"))
	if err != nil {
		log.Category("main").Info("Unable to convert SMTP_TLS to int, setting false")
		enc = 0
	}
	emailer.SetOptions(emailer.Options{
		Hostname:   Getenv("SMTP_HOSTNAME", "localhost"),
		Port:       port,
		Username:   Getenv("SMTP_USERNAME", "noreply@nzvirtual.org"),
		Password:   Getenv("SMTP_PASSWORD", "password"),
		Encryption: enc != 0,
	})

	log.Category("main").Info("Looking up API Key")
	ApiKey = os.Getenv("API_KEY")

	log.Category("main").Info("Configuring gin server")
	server := NewServer(appenv)

	server.engine.Run(fmt.Sprintf(":%s", Getenv("PORT", "3000")))
}

func Getenv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}
