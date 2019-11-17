package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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
	cron := &Cron{
		Second:  true,
		Minute:  false,
		Hour:    false,
		Day:     false,
		Week:    false,
		Every:   1,
		At:      "",
		Actions: nil,
	}

	go cron.StartCron()

	// os.Exit(0)

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

	err := StartServer(sv, mux)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server started on port: %s", viper.Get("server.port"))
}
