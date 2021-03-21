package main

import (
	"fmt"
	"os"
	"time"
	"net"
)

func main() {
	service := ":1200"
	// 获取udp连接
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	// 开启监听
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn *net.UDPConn) {
	// 定义一个[]byte类型的数据，用来装请求体
	buf := make([]byte, 1024)
	_, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
