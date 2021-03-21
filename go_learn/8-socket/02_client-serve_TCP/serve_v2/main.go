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
	service := ":1200"
	// 获取tcpAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	// 监听tcp
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
		// go 实现了多并发
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}