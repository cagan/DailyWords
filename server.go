package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Server struct {
	Addres string 
	Port string
}

var defaultConfig Server

func init() {
	if viper.GetBool(`debug`) {

	}
}

func NewServer(server Server) *Server {
	server := new(Server)
	return server
}

func StartServer(sv Server, mux *http.ServeMux) error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", sv.Addres, sv.Port), mux)

	if err != nil {
		return err
	}

	return nil
}