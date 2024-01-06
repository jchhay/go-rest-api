package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadEnv() {

	err := godotenv.Load(`../.env`)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

var config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type ServerConfiguration struct {
	Port        string
	Secret      string
	Mode        string
	UseDatabase bool
}

func SetupConfig() {
	// Reading Configuration File and Environment Variables
	// Environment Variable settings are prioritized over File

	viper.SetConfigFile("../config/config.yml")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("server.mode", "MODE")
	viper.BindEnv("server.usedatabase", "USE_DATABASE")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

}

func SetupConfigFile(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	config = configuration
}

func GetConfig() *Configuration {
	return config
}
