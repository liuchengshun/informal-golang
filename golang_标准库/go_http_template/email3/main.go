package main

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/go-playground/validator/v10"
)

type emailService struct {
	from string
}

type EmailService interface {
	SendCaptcha(message, code, email string) error
}

func main() {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="iso-8859-15">
		<title>MMOGA POWER</title>
	</head>
	<body>
		GO 发送邮件，官方连包都帮我们写好了，真是贴心啊！！！
	</body>
	</html>`

	eamilSvc := NewEmailService()
	if err := eamilSvc.SendCaptcha(body, "123456", "1136089132@qq.com"); err != nil {
		log.Fatal(err)
	}
}

func NewEmailService() EmailService {
	return &emailService{
		from: "lcs_shun@qq.com",
	}
}

func (svc *emailService) SendCaptcha(message, code, email string) error {
	validate := validator.New()
	if err := validate.Var(email, "email"); err != nil {
		return fmt.Errorf("invalid email address, %v", err)
	}

	auth := svc.auth(email)
	msg := svc.getCaptchaMessage(email, message)

	err := smtp.SendMail(
		"smtp.qq.com:25",
		auth,
		svc.from,
		[]string{email},
		[]byte(msg),
	)

	return err
}

func (svc *emailService) auth(email string) smtp.Auth {
	return smtp.PlainAuth("", svc.from, "ndvqyipzfqrsgihg", "smtp.qq.com")
}

// func (svc *emailService) auth(email string) smtp.Auth {
// 	return smtp.PlainAuth("", svc.from, "rwkiovjiuqsihbhf", "smtp.qq.com")
// }

func (svc *emailService) getCaptchaMessage(email, message string) string {
	subject := "邮箱验证码"
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := "From: " + svc.from + "\r\nTo: " + email +
		"\r\nSubject: " + subject + "\r\n" + contentType + "\r\n" + message
	return msg
}
