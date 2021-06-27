package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type person struct {
	Name string
	Age  int
	Sex  int
}

func (p *person) String() string {
	buf := bytes.NewBufferString("this is ")
	buf.WriteString(p.Name + ", ")
	if p.Sex == 0 {
		buf.WriteString("he is a man,")
	} else {
		buf.WriteString("she is a woman,")
	}

	buf.WriteString("age is " + strconv.Itoa(p.Age))
	return buf.String()
}

func main() {
	p := &person{"liuchengshun", 0, 23}
	fmt.Println(p)
}
