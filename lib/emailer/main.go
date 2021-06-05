package emailer

import (
	"time"

	log "github.com/dhawton/log4g"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Options struct {
	Hostname   string
	Port       int
	Username   string
	Password   string
	Encryption bool
}

var opt Options

func SetOptions(options Options) {
	opt = options
}

func SendEmail(to string, subject string, body string) {
	server := mail.NewSMTPClient()
	server.Host = opt.Hostname
	server.Port = opt.Port
	server.Username = opt.Username
	server.Password = opt.Password
	if opt.Encryption {
		server.Encryption = mail.EncryptionSTARTTLS
	}
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	smtpClient, err := server.Connect()
	if err != nil {
		log.Category("emailer").Fatal("Could not connect to SMTP " + err.Error())
		panic("")
	}

	email := mail.NewMSG()
	email.SetFrom("NZ Virtual <noreply@nzvirtual.org>").AddTo(to).SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	if email.Error != nil {
		log.Category("emailer").Fatal("Error building email " + err.Error())
		panic("")
	}

	err = email.Send(smtpClient)
	if err != nil {
		log.Category("emailer").Fatal("Error sending email " + err.Error())
		panic("")
	} else {
		log.Category("emailer").Debug("Email sent to " + to)
	}
}

func IfThenElse(condition bool, t interface{}, f interface{}) interface{} {
	if condition {
		return t
	}
	return f
}
