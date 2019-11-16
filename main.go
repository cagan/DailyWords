package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	_ "github.com/jasonlvhit/gocron"
)

type Translate struct {
	word string
	provision string
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
	startCron()
	os.Exit(0)
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

	csvFile, err := os.Open("words.csv")
	if err != nil {
		panic("Can not open csv file")
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	var words[] Translate

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		words = append(words, Translate{
			word:     line[0] ,
			provision: line[1],
		})
	}

	fmt.Println(words[0])

	err = StartServer(*sv, mux)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server started on port: %s", viper.Get("server.port"))
}

func startCron() {
	gocron.Every(1).Second().From(gocron.NextTick()).Do(cagan)
	<- gocron.Start()
}

func cagan() {
	fmt.Println("Cagan is the king")
}