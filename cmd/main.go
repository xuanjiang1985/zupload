package main

import (
	"flag"
	"log"
	"zupload/config"
)

func main() {
	configFile := flag.String("c", "", "-c x.yaml")
	flag.Parse()

	conf, err := config.InitConfig(*configFile)
	if err != nil {
		log.Fatalf("load config err: %v", err)
	}
	_ = conf
}
