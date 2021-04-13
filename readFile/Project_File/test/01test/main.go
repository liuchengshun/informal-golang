package main

import (
	"flag"
	"fmt"
	"os"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	if *verbose {
		fmt.Println("verbose is:", *verbose)
	}
	fmt.Println("verbose:", *verbose)

	os.Stdin.Read(make([]byte, 1))

	fmt.Fprintf(os.Stderr, "du1:%v\n", "err")

	fmt.Println("End")
}
