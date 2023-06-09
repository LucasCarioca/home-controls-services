package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LucasCarioca/home-controls-services/pkg/config"
	"github.com/LucasCarioca/home-controls-services/pkg/datasource"
	"github.com/LucasCarioca/home-controls-services/pkg/server"
)

func getEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		envFlag := flag.String("e", "dev", "")
		flag.Usage = func() {
			fmt.Println("Usage: server -e {mode}")
			os.Exit(1)
		}
		flag.Parse()
		env = *envFlag
	}
	return env
}

func main() {
	config.Init(getEnv())
	datasource.Init(config.GetConfig())
	server.Init(config.GetConfig())
}
