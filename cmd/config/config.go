package config

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Config - full config
type Config struct {
	Glacli Metadata
}

// Metadata - main config structure
type Metadata struct {
	Default string
	Servers map[string]Server
}

// Server - server config
type Server struct {
	Host, Token string
}

var conf struct {
	Full  Config
	Def   string
	InUse Server
}

// Get - get config from file
func Get(file string) (Server, error) {
	viper.SetConfigType("yaml")
	if file != "" {
		// Use config file from the flag.
		viper.SetConfigFile(file)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return conf.InUse, err
		}

		// Search config in home directory with name ".glacli" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".glacli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
		return conf.InUse, err
	}

	//
	err := viper.Unmarshal(&conf.Full)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
		return conf.InUse, err
	}

	conf.Def = conf.Full.Glacli.Default
	conf.InUse = conf.Full.Glacli.Servers[conf.Def]

	return conf.InUse, err
}
