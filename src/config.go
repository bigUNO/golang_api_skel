package main

import (
  "os"
	//"fmt"
  "log"

  "github.com/BurntSushi/toml"
)

type tomlConfig struct {
  Title        string
  MongoDBHosts string
  AuthDatabase string
  AuthUserName string
  AuthPassword string
  IsDrop       bool
}


// Reads info from config file
func ReadConfig() tomlConfig {
	//var configfile = flags.Configfile
  var configfile = "../config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config tomlConfig
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Loading config file: ", config.Title)
	return config
}
