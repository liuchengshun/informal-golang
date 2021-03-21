package main

import (
	"errors"
	"fmt"
	// "net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	arith := new(Arith)
	// 注册函数
	rpc.Register(arith)
	// 建立http连接
	rpc.HandleHTTP()

	err := http.ListenAndServe(":8080", nil)
	// l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
	// go http.Serve(l, nil)
}