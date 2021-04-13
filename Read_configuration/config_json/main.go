package main

import (
	"flag"
	"log"

	"github.com/Unknwon/goconfig"
)

var cfgFile string

func init() {
	flag.StringVar(&cfgFile, "config", "config.yaml", "read configuration")
}

func main() {
	cfg, err := goconfig.LoadConfigFile(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	host, _ := cfg.GetValue("server", "host")
	log.Println("host:", host)
}
