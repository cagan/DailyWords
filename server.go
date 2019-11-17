package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	Address string
	Port    string
}

func NewServer(sv Server) *Server {
	server := new(Server)
	server.Address = sv.Address
	server.Port = sv.Port
	return server
}

func StartServer(sv Server, mux *http.ServeMux) error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", sv.Address, sv.Port), mux)

	if err != nil {
		return err
	}

	return nil
}