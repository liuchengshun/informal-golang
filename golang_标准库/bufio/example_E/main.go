package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir := os.Args[1]
	listAll(dir, 0)
}

func listAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			for tempHier := curHier; tempHier > 0; tempHier-- {
				fmt.Println("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tempHier := curHier; tempHier > 0; tempHier-- {
				fmt.Println("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
