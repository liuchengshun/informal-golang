package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type UserMsg struct {
	UserName string
	Phone    string
	Email    string
	Company  string
	CheckURL string
}

type Message struct {
	ActiveURL  string
	UserMsg    UserMsg
	SuccessURL string
	FailURL    string
	Captcha    string
}

func ParseAdmin(w http.ResponseWriter, r *http.Request) {
	Msg := &Message{
		ActiveURL: "https://www.tfcloud.com/",
		UserMsg: UserMsg{
			UserName: "ZhangSan",
			Phone:    "123456",
			Email:    "132@qq.com",
			Company:  "tongfangyouyun",
			CheckURL: "https://www.tfcloud.com/",
		},
		SuccessURL: "https://www.tfcloud.com/",
		FailURL:    "https://www.tfcloud.com/",
		Captcha:    "7d8F",
	}
	// parse template
	t, err := template.ParseFiles("./template/success.html", "./static/footer.html", "./static/header.html")
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}
	t.ExecuteTemplate(w, "content", nil)
	err = t.ExecuteTemplate(os.Stdout, "content", Msg)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

func main() {
	// 注册路由
	http.HandleFunc("/email", ParseAdmin)
	// 创建一个服务器
	http.ListenAndServe(":8080", nil)
}
