package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/topboyasante/mrkt-api/api/v1/utils"
	"github.com/topboyasante/mrkt-api/internal/config"
)

type SMTPAuth struct {
	identity string
	username string
	password string
	host     string
}

var EmailConfig = initSMTPAuth()

func initSMTPAuth() *SMTPAuth {
	auth := &SMTPAuth{
		identity: "",
		username: config.ENV.SMTPUsername,
		password: config.ENV.SMTPPassword,
		host:     config.ENV.SMTPHost,
	}

	return auth
}

func SendMailWithSMTP(authVar *SMTPAuth, fromName, subject, templatePath string, values any, to []string) error {
	var body bytes.Buffer

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		utils.Logger().Errorf("error: unable to parse template: %v", err)
		return err
	}

	err = t.Execute(&body, values)
	if err != nil {
		utils.Logger().Errorf("error: unable to execute template: %v", err)
		return err
	}

	auth := smtp.PlainAuth(
		authVar.identity,
		authVar.username,
		authVar.password,
		authVar.host,
	)

	from := fmt.Sprintf("%s <%s>", fromName, config.ENV.SMTPUsername)

	headers := fmt.Sprintf("From: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";", from)
	msg := fmt.Sprintf("Subject: %v\n%v\n\n%v", subject, headers, body.String())

	err = smtp.SendMail(
		config.ENV.SMTPAddress,
		auth,
		config.ENV.SMTPUsername,
		to,
		[]byte(msg),
	)

	if err != nil {
		utils.Logger().Errorf("error: unable to send mail: %v", err)
		return err
	}

	return nil
}
