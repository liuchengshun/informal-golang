package main

import "fmt"

type student struct {
	name string
}

func (s *student) YourName() string {
	return s.name
}

type Mgs interface {
	YourName() string
}

func main() {
	var stu1 Mgs = (*student)(nil)
	fmt.Println("stu1:", stu1)

	var num int = 56
	str := fmt.Sprintf("testing number is %v", num)
	fmt.Println("str:", str)
}
