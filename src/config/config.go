/*
Package config contains all the configurations for the service
*/
package config

import (
	"github.com/spf13/viper"
	"log"
	"logger"
	"strings"
	"encoding/json"
)

var config *AppConfig

// Create private data struct to hold config options.
type ServerInfo struct {
	Host string
	Host_2 string
	Port string
	NestedList1 []int8
	NestedList2 []float32
	NestedList3	[]string
	NestedList4 []float32
}

type AppConfig struct {
	Age     int8       `yaml:"age"`
	Hacker  bool       `yaml:"hacker"`
	Name    string     `yaml:"name"`
	Hobbies []string   `yaml:"hobbies"`
	Nums	[]int8     `yaml:"nums"`
	Server  ServerInfo `yaml:server`
}

// Init is an exported method that takes the environment, starts the viper (external lib),
// and returns the configuration struct.
func Init(env string) {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.AddConfigPath("src/config/")
	v.SetConfigName("config")

	err = v.MergeInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatal("Error on parsing configuration file. Error: " + err.Error())
	}
	conf := &AppConfig{}
	err  = v.Unmarshal(conf)

	if err != nil {
		// log.Fatal("unable to decode into config struct, Error: " + err.Error())
		panic(err.Error())
	}


	config = conf
	cfgJson, err := json.Marshal(config)
	logger.Log.Infof("env: %s, config: %s", env, cfgJson)

}

// GetConfig function to expose the config object
func GetConfig() *AppConfig {
	return config
}
