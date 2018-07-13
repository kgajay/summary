package main

import (
	"flag"
	"fmt"
	"os"
	"logger"
	"config"
	"server"
	"dao"
	"model"
)

func main() {
	// Load config based on the server mode
	env := flag.String("e", "DEV", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	logger.Init(*env)
	config.Init(*env)
	dao.Init()
	model.Migrate()
	server.Init(*env)
}
