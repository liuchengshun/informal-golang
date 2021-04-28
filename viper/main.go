package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	server := viper.GetString(`server.host`)
	fmt.Println("server:", server)
}
