package main

import (
	"fmt"
	_ "github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
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

	cron := NewCron(Cron{
		Second:  true,
		Minute:  false,
		Hour:    false,
		Day:     false,
		Week:    false,
		Every:   22,
		At:      "",
		Actions: nil,
	})

	StartCron(*cron)

	os.Exit(0)

	mux := http.NewServeMux()
	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		fmt.Println("Development Environment")
	}

	defaultConfig := Server{
		Address: viper.GetString("server.address"),
		Port:    viper.GetString("server.port"),
	}

	sv := NewServer(defaultConfig)

	mux.HandleFunc("/words", TranslationWords)

	err := StartServer(*sv, mux)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server started on port: %s", viper.Get("server.port"))
}
