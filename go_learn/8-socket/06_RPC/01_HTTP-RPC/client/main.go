package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:", os.Args[0], "server")
		os.Exit(0)
	}
	// 取出Ip地址
	serverAddress := os.Args[1]

	// 获取rpc的http连接
	client, err := rpc.DialHTTP("tcp", serverAddress + ":8080")
	if err != nil {
		fmt.Println("dailing:", err)
	}
	args := Args{17, 8}
	var reply int
	// 使用服务端的连接,并使用服务端的方法
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot Quotient
	// 使用服务端的方法
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}