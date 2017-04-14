package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config struct defines the config structure
type Config struct {
	Mongo         MongoConfig `json:"mongo"`
	Host          string      `json:"host"`
	ProxyCheckURL string      `json:"proxyCheckURL"`
}

// MongoConfig has config values for Mongo
type MongoConfig struct {
	Addr       string `json:"addr"`
	DB         string `json:"db"`
	Collection string `json:"collection"`
	Event      string `json:"event"`
}

// NewConfig parses config file and return Config struct
func NewConfig() *Config {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln("Read config file error.")
	}
	config := &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}

	return config
}
