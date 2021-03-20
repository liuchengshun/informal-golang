package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"os"
	"io"
	"net/http"
	"strings"
	"strconv"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	// 如果没有调用ParseForm方法,下面无法获取表单的数据
	fmt.Println("r.Form:", r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("r.Form['url_long']:",r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello, astaxie")
}

func login(w http.ResponseWriter, r *http.Request) {
	// 获取请求的方法
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		// 解析模板
		t, _ := template.ParseFiles("./login.html")
		// 渲染模板
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("./upload.html")
		t.Execute(w, token)
	} else {
		// 允许传输的最大文件为32M
		r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在upload目录
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Println("ServeHTTP error :", err)
	}
}
