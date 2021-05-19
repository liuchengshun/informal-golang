package main

import (
	"math"
	"net/http"
	"net/rpc"
)

type MathUtil struct{}

type Parmas struct {
	Arg1 int
	Arg2 int
}

func (mu *MathUtil) CalculateCircleArea(req float64, resp *float64) error {
	*resp = math.Pi * req * req
	return nil
}

func (mu *MathUtil) Add(reqs Parmas, resp *int) error {
	*resp = reqs.Arg1 + reqs.Arg2
	return nil
}

func main() {
	// 初始化指针数据类型
	mathUtil := new(MathUtil)

	// 调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err)
	}

	// 通过该函数把mathUtil中提供的服务注册到HTTP协议上，方便调用者可以利用http的方式进行传输
	rpc.HandleHTTP()

	// 在特定的端口进行监听
	// listen, err := net.Listen("tcp", ":8989")
	// if err != nil {
	// 	panic(err)
	// }
	// go http.Serve(listen, nil)
	http.ListenAndServe(":8989", nil)
}
