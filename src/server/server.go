package server

import (
	"config"
	"logger"
	"sync"
	"fmt"
)

var onceRest sync.Once

// Init function to initialize the service
func Init(env string) {
	onceRest.Do(func() {
		conf := config.GetConfig()
		logger.Log.Info("Initializing Rest server")
		r := NewRouter()

		fmt.Println("API's:")
		for _, elem := range r.Routes() {
			fmt.Printf("%-10s\t%-30s\t%-20s\n", elem.Method, elem.Path, elem.Name)

		}
		// This is done to Suppress Accept Incoming Network Connections warnings on OSX
		var err error
		if env == "DEV" {
			err = r.Start("localhost" + conf.Server.Port)
		} else {
			err = r.Start(conf.Server.Port)
		}

		if err != nil {
			logger.Log.Fatal("Unable to bring service up: " + err.Error())
		}
	})
}
