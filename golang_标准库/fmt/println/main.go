package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func main() {
	p := &Person{Name: "liuchengshun", Age: 12, Sex: 2}
	fmt.Println("p:", p)
	fmt.Printf("%#v", p)
	// fmt.Printf("%L", p)
}

// func (this *Person) String() string {
// 	buffer := bytes.NewBufferString("this is")
// 	buffer.WriteString(this.Name + ",")
// 	if this.Sex == 0 {
// 		buffer.WriteString("He ")
// 	} else {
// 		buffer.WriteString("She ")
// 	}

// 	buffer.WriteString("is ")
// 	buffer.WriteString(strconv.Itoa(this.Age))
// 	buffer.WriteString(" years old.")
// 	return buffer.String()
// }

// func (this *Person) Format(f fmt.State, c rune) {
// 	if c == 'L' {
// 		f.Write([]byte(this.String()))
// 		f.Write([]byte(" Person has three fields."))
// 	} else {
// 		// 没有此句，会导致 fmt.Printf("%s", p) 啥也不输出
// 		f.Write([]byte(fmt.Sprintln(this.String())))
// 	}
// }
