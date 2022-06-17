package email

import (
	"dabc/config"
	"gopkg.in/gomail.v2"
	"strconv"
)

var D *gomail.Dialer

func InitMail()  {
	port, _ := strconv.Atoi(config.C.Email.Port)
	d := gomail.NewDialer(config.C.Email.Host, port, config.C.Email.User, config.C.Email.PWD)
	D = d
}

func SendEmail(to []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From": {m.FormatAddress(config.C.Email.User, "DABC_Robot")},
		"To": to,
		"Subject": {subject},
	})
	m.SetBody("text/html", body)
	err := D.DialAndSend(m)
	return err
}
