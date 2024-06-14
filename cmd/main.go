package main

import (
	"flag"
	"log"
	"rate-limiter/api/server"
	"rate-limiter/config"
)

func main() {

	var (
		configEnv = flag.String("config", "", "config file (default is config.yaml)")
	)
	flag.Parse()

	if *configEnv == "" {

		if err := config.FileExist("config/config.yaml"); err == nil {
			*configEnv = "config/config.yaml"
		} else {
			log.Fatal("parameter -config is required. Use -config=config/config.yaml for example")
		}
	}
	if err := config.Load(*configEnv); err != nil {
		log.Fatalf("Failed get config [%s]", err)
	}

	log.Println(" [*] Start service")
	server.Execute()
}
