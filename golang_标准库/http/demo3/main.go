package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	URL := "https://www.baidu.com"

	req, _ := http.NewRequest(http.MethodGet, URL, nil)

	cli := http.Client{}
	response, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var payload interface{}
	err = json.NewDecoder(response.Body).Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response.Body:", response.Body)
	fmt.Println("payload for baidu:", payload)
}
