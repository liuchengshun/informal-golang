package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	service := ":7777"
	// 1.获取一个tcp4
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 2.监听端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	
	for {
		// 获取从服务器端来的请求
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		// 将当地的时间返回到客户端
		conn.Write([]byte(daytime))
		conn.Close()
	}
}