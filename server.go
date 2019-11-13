package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {
	Addres string 
	Port string
}

func init() {
	if viper.GetBool(`debug`) {

	}
}

func NewServer(sv Server) *Server {
	server := new(Server)
	server.Addres = sv.Addres
	server.Port = sv.Port
	return server
}

func StartServer(sv Server, mux *http.ServeMux) error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", sv.Addres, sv.Port), mux)

	if err != nil {
		return err
	}

	return nil
}