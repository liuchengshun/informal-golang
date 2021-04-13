package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/liuchengshun/imformal-form/go_http_template/email2/config"
)

type Request struct {
	to       []string
	subject  string
	body     string
	from     string
	password string
	host     string
}

func NewRequest(from string, to []string, subject, body, password, host string) *Request {
	return &Request{
		to:       to,
		subject:  subject,
		body:     body,
		from:     from,
		password: password,
		host:     host,
	}
}

func main() {
	// Msg := &Message{
	// 	ActiveURL: "http://127.0.0.1:5677/api/v1/auth/register?code=a401d864-1556-4842-a41c-877b3c5b7c80",
	// 	UserMsg: UserMsg{
	// 		UserName: "ZhangSan",
	// 		Phone:    "123456",
	// 		Email:    "132@qq.com",
	// 		Company:  "tongfangyouyun",
	// 		CheckURL: "http://127.0.0.1:5677/api/v1/auth/register?code=a401d864-1556-4842-a41c-877b3c5b7c80",
	// 	},
	// 	SuccessURL: "https://www.tfcloud.com/",
	// 	FailURL:    "https://www.tfcloud.com/",
	// 	Captcha:    "7d8F",
	// }
	// NewRequest(from string, to []string, subject, body, password, host string)

	config.Load("config.example.json")

	r := NewRequest("1136089132@qq.com", []string{"1351169665@qq.com"}, "同方有云",
		" ", "rwkiovjiuqsihbhf", "smtp.qq.com")
	verityURL := "http://127.0.0.1:5677/api/v1/auth/register?code=a401d864-1556-4842-a41c-877b3c5b7c80"

	if str, err := RenderEmailRegisterVerifyTemplate(verityURL); err == nil {
		r.body = str
		r.SendEmail()
	}
	// if str, err := RenderEmailRegisterReviewTemplate("ZhangSan", "123456", "132@qq.com", "tongfangyouyun", verityURL); err == nil {
	// 	r.body = str
	// 	r.SendEmail()
	// }
}

func (r *Request) Auth() smtp.Auth {
	return smtp.PlainAuth("", r.from, r.password, r.host)
}
func (r *Request) SendEmail() (bool, error) {
	subject := "Subject: " + r.subject + "!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n\n"
	msg := []byte(subject + mime + r.body)
	auth := r.Auth()

	if err := smtp.SendMail(r.host+":25", auth, r.from, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func renderEmailTemplate(name string, data interface{}) (string, error) {
	t, err := template.ParseFiles("./templates/email.tmpl")
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	d := struct {
		Data    interface{}
		Logo    string
		Company string
	}{
		Data:    data,
		Logo:    config.GetString("ui.about.logo"),
		Company: config.GetString("ui.about.company_name"),
	}
	fmt.Println("LogoURL:", d.Logo)
	if err := t.ExecuteTemplate(buf, name, d); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderEmailCaptchaTemplate(code string) (string, error) {
	data := struct {
		Code string
	}{
		Code: code,
	}
	return renderEmailTemplate("captcha", data)
}

func RenderEmailRegisterOKTemplate(loginURL string) (string, error) {
	data := struct {
		LoginURL string
	}{
		LoginURL: loginURL,
	}
	return renderEmailTemplate("register-ok", data)
}

func RenderEmailRegisterFailedTemplate(registerURL string) (string, error) {
	data := struct {
		RegisterURL string
	}{
		RegisterURL: registerURL,
	}
	return renderEmailTemplate("register-failed", data)
}

func RenderEmailRegisterReviewTemplate(name, phone, email, company, reviewURL string) (string, error) {
	data := struct {
		Name      string
		Phone     string
		Email     string
		Company   string
		ReviewURL string
	}{
		Name:      name,
		Phone:     phone,
		Email:     email,
		Company:   company,
		ReviewURL: reviewURL,
	}
	return renderEmailTemplate("register-review", data)
}

func RenderEmailRegisterVerifyTemplate(verifyURL string) (string, error) {
	data := struct {
		VerifyURL string
	}{
		VerifyURL: verifyURL,
	}
	return renderEmailTemplate("register-verify", data)
}
