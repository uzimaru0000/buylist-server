package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   *Server
	Firebase *Firebase
	APIKey   *APIKey
}

type Server struct {
	Port string `default:":5000"`
}

type Firebase struct {
	AcountKey string
}

type APIKey struct {
	YahooAPIKey string
}

var config Config

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	log.Print(os.Getenv("SERVER_PORT"))
	err = envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &config
}
