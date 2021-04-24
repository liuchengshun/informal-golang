package main

import (
	"fmt"
)

func main() {
	var (
        name string
        age  int
    )
    n, _ := fmt.Sscan("polaris 28", &name, &age)
    // 可以将"polaris 28"中的空格换成"\n"试试
    // n, _ := fmt.Sscan("polaris\n28", &name, &age)
    fmt.Println(n, name, age)
}
