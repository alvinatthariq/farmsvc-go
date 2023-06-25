package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MySQL MySQLConfig `mapstructure:"mysql"`
	Redis RedisConfig `mapstructure:"redis"`
	Port  string      `mapstructure:"port"`
}

type MySQLConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
}

type RedisConfig struct {
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
