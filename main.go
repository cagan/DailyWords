package main

import (
	"fmt"
	"log"

	_ "github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Address string
	Port    string
}

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

	defaultConfig := Config{
		Address: viper.GetString("server.address"),
		Port:    viper.GetString("server.port"),
	}

	sv, err := NewServer(defaultConfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server started on port: %s", viper.Get("server.port"))
	if err := sv.StartServer(); err != nil {
		panic(err)
	}

}
