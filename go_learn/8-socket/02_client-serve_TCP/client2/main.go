package main

import (
	"fmt"
	// "io/ioutil"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args)
		os.Exit(1)
	}
	// 1.从输入端上获取一个tcp4的地址
	service := os.Args[1]
	// 2.获取一个 tcpAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 3.利用获得的tcpAddr建立一个tcp连接
	tcpConn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)

	// 4.现在可以使用获得的conn来进行数据交互
	// 发起一个请求
	_, err = tcpConn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result := make([]byte, 256)
	// 读取数据，并将数据放入result中
	_, err = tcpConn.Read(result)
	checkError(err)
	// 将数据打印出来
	fmt.Println(string(result))
	os.Exit(0)
}