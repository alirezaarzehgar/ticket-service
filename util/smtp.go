package util

import (
	"net/smtp"

	"github.com/alirezaarzehgar/ticketservice/config"
)

var (
	SmtpAuth smtp.Auth
	SmtpConf config.SmtpConf
)

type SmtpContent struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func InitMail(c config.SmtpConf) {
	SmtpConf = c
	SmtpAuth = smtp.PlainAuth("", SmtpConf.FromAddress, SmtpConf.Password, SmtpConf.Host)
}
