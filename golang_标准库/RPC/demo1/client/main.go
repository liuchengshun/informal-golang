package main

import (
	"fmt"
	"net/rpc"
)

type Parmas struct {
	Arg1 int
	Arg2 int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8989")
	if err != nil {
		panic(err)
	}

	var req float64
	req = 3.4333
	var resp float64

	err = client.Call("MathUtil.CalculateCircleArea", req, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp:", resp)

	args := Parmas{
		Arg1: 4,
		Arg2: 6,
	}
	var resp2 int
	if err := client.Call("MathUtil.Add", args, &resp2); err != nil {
		panic(err)
	}
	fmt.Println("the two together:", resp2)

}
