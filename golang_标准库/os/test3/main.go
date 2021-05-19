package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	fileName1 := "D:/HackerNEws_app/node_modules/.bin"
	fileName2 := "bb.txt"
	fmt.Println(filepath.IsAbs(fileName1)) //true
	fmt.Println(filepath.IsAbs(fileName2)) //false
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2)) // /Users/ruby/go/src/l_file/bb.txt

	fmt.Println("获取父目录：", path.Join(fileName1, ".."))

	// 创建目录
	// err := os.Mkdir("D:/WrokSpace/Learn_go/informal-golang/golang_标准库/os/test_test", 0666)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }

	// 创建文件
	// file, err := os.Create("D:/WrokSpace/Learn_go/informal-golang/golang_标准库/os/test_test/bb.txt")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }
	// defer file.Close()

	// 4.打开文件：
	// file3, err := os.Open("D:/WrokSpace/Learn_go/informal-golang/golang_标准库/os/test_test/bb.txt") //只读的
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }
	// fmt.Println(file3)
	// bs := []byte{}
	// fmt.Println("bs:", bs)
	// n, err := file3.Read(bs)
	// fmt.Println("n:", n)
	// fmt.Println("err:", err)
	// fmt.Println("bs:", string(bs))

	file, err := os.OpenFile("D:/WrokSpace/Learn_go/informal-golang/golang_标准库/os/test_test/bb.txt", os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	file.WriteString("helloworld\n")

	bs := []byte{97, 98, 99, 100} //a,b,c,d
	file.Write(bs)

}
