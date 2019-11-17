package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
)

type Server struct {
	Address  string
	Port     string
	Watcher  *fsnotify.Watcher
	Csv      *os.File
	shutdown chan struct{}
}

func NewServer(cfg Config) (server *Server, err error) {
	server = new(Server)
	server.Address = cfg.Address
	server.Port = cfg.Port
	server.Watcher, err = fsnotify.NewWatcher()

	server.Csv, err = os.Open("eng_tur.csv")
	if err != nil {
		return
	}

	go server.watchFileUpdates()

	err = server.Watcher.Add(server.Csv.Name())
	if err != nil {
		return
	}

	return
}

func (sv *Server) StartServer() error {
	// @TODO move mux under NewServer.
	mux := http.NewServeMux()
	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		fmt.Println("Development Environment")
	}
	mux.HandleFunc("/words", sv.TranslationWords)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", sv.Address, sv.Port), secure(mux))
	if err != nil {
		return err
	}
	close(sv.shutdown)

	if err := sv.Watcher.Close(); err != nil {
		return err
	}

	if err := sv.Csv.Close(); err != nil {
		return err
	}

	return nil
}

func secure(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://postwoman.io")
		h.ServeHTTP(w, r)
	})
}

func (sv *Server) watchFileUpdates() {
	for {
		select {
		case event, ok := <-sv.Watcher.Events:
			if !ok {
				return
			}
			log.Printf("Event: %s", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				// Close and reopen csv file after every update on file.

				err := sv.Csv.Close()
				if err != nil {
					log.Printf("Error: %s", err.Error())
					continue
				}

				sv.Csv, err = os.Open("eng_tur.csv")
				if err != nil {
					log.Printf("Error: %s", err.Error())
					continue
				}
			}
		case err, ok := <-sv.Watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Error: %s", err.Error())
		case <-sv.shutdown:
			return
		}
	}
}
