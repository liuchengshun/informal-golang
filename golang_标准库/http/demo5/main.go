package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Set("regsiter_id", "c8af79ec-34bf-4c0c-b331-4c23b5925656")
	v.Add("name", "Lisi")

	u := url.URL{
		Scheme:   "https",
		Host:     "127.0.0.1:8080",
		Path:     "/hello",
		RawQuery: v.Encode(),
	}

	fmt.Println("u.string:", u.String())
}
