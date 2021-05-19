package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before run")
	code := m.Run()
	fmt.Println("after run")
	os.Exit(code)
}

func TestAdd(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			res := Add(i, j)
			fmt.Println("result :", res)
			if res != (i + j) {
				t.Fatal("result is error")
			}
		}
	}
}
