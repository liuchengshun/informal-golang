package main

import (
	"fmt"
	"log"
	"os"
)

func Print(funcname string, err error, valueReturn ...interface{}) {
	fmt.Println("=======================")
	log.Println("os->function name: ", funcname)
	for k, v := range valueReturn {
		log.Printf("return values: \n")
		log.Printf("value[%v] = %v\n", k, v)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=======================")
}

func main() {
	// Hostname返回内核提供的主机名。
	name, err := os.Hostname()
	Print("Hostnam", err, name)

	// Getpagesize返回底层的系统内存页的尺寸。
	size := os.Getpagesize()
	Print("Getpagesize", nil, size)

	// Environ返回表示环境变量的格式为"key=value"的字符串的切片拷贝。
	// strs := os.Environ()
	// Print("Environ", nil, strs)

	id := os.Getuid()
	Print("Getuid", nil, id)

	id = os.Geteuid()
	Print("Geteuid", nil, id)
}
