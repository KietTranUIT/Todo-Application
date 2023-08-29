package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"user-service/conf"
	"user-service/internal/core/dto"
	"user-service/internal/core/port/service"
)

type MailService struct {
	username string
	url      string
	mail     smtp.Auth
}

const (
	subject = "Subject: Todo-App verify email!\n"
	mime    = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func NewMailService(config conf.Config_MailService) service.MailService {
	return MailService{
		username: config.From,
		url:      config.URL(),
		mail:     smtp.PlainAuth("", config.From, config.GetPassword(), config.Host),
	}
}

func (m MailService) CreateMail(email string, code string) dto.MailData {
	return dto.MailData{
		From:    m.username,
		To:      []string{email},
		Subject: "Verify Email",
		Code:    code,
	}
}

func (m MailService) SendMail(mailData dto.MailData) error {
	buffer := new(bytes.Buffer)
	if tmpl, err := template.ParseFiles("view/verify_email.html"); err != nil {
		return err
	} else {
		tmpl.Execute(buffer, mailData)
	}

	to := fmt.Sprintf("To: %s\r\n", mailData.To[0])
	body := buffer.String()

	message := to + subject + mime + "\r\n" + body

	err := smtp.SendMail(m.url, m.mail, mailData.From, mailData.To, []byte(message))

	if err != nil {
		return err
	}
	return nil
}
