package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	mux := http.NewServeMux()
	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		fmt.Println("Development Environment")
	}

	defaultConfig := Server{
		Addres: viper.GetString("server.address"),
		Port:   viper.GetString("server.port"),
	}

	sv := NewServer(defaultConfig)

	mux.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World")
	})

	err := StartServer(*sv, mux)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server started on port: %s", viper.Get("server.port"))
}
