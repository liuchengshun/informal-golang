package main

import (
	"fmt"
	"os"
	"net"
	"time"
	"strings"
	"strconv"
)

func main() {
	service := ":1200"
	// 获取tcpAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 开启监听模式
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// 开启服务
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()

	for {
		read_len, err := conn.Read(request)
		if err != nil {
			continue
		}
		if read_len == 0 {
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		request = make([]byte, 128) // clear last read content
	}
}