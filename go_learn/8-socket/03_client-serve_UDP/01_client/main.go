package main

import (
	"fmt"
	"os"
	"net"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	// 发送数据
	_, err = conn.Write([]byte("anything"))
	checkError(err)

	// 定义一个[]byte类型的数据，用来装返回的数据
	buf := make([]byte, 1028)
	n, err := conn.Read(buf)
	checkError(err)
	fmt.Println(string(buf[:n]))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}