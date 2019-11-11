package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func init() {
	viper.SetConfigFile(`config_local.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	mux := http.NewServeMux()

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
