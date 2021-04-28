package main

import (
	"fmt"
	"net/url"

	"github.com/go-redis/redis/v8"
)

const rdURL string = "redis://liuliu@127.0.0.1:6379?password=abc123&username=lcs&sentinel_master=razor"

type Options redis.UniversalOptions

func main() {
	o, err := ParseURL(rdURL)
	fmt.Println("options:", o)
	fmt.Println("error:", err)
}

func ParseURL(rdURL string) (*Options, error) {
	u, err := url.Parse(rdURL)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "redis" {
		return nil, fmt.Errorf("the redis url is error, want redis but %s", u.Scheme)
	}

	o := &Options{}

	if p, ok := u.User.Password(); ok {
		fmt.Println("password:", p)
	}

	user := u.User.Username()
	fmt.Println("username:", user)

	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("the redis url of query is error")
	}

	for k := range values {
		switch k {
		case "password":
			o.Password = values.Get(k)
		case "sentinel_master":
			o.MasterName = values.Get(k)
		case "sentinel_pass":
			o.SentinelPassword = values.Get(k)
		}
	}

	return o, nil
}
