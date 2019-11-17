package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	reader *csv.Reader
)

type Translate struct {
	Word      string `json:"word"`
	Provision string `json:"provision"`
}

type TranslationOptions struct {
	Limit int `json:"limit"`
	FromLang string `json:"from_lang"`
	ToLang string `json:"to_lang"`
}

func init() {
	csvFile, err := os.Open("eng_tur.csv")
	if err != nil {
		panic("Can not open csv file")
	}
	reader = csv.NewReader(bufio.NewReader(csvFile))
}

func TranslationWords(w http.ResponseWriter, r *http.Request) {
	queryVal := r.URL.Query()

	fmt.Println(queryVal["limit"][0])
	opts := TranslationOptions{}

	if queryVal["limit"] != nil {
		opts.Limit, _ = strconv.Atoi(queryVal["limit"][0])
	}

	if queryVal["from"] != nil {
		opts.FromLang = queryVal["from"][0]
	}

	if queryVal["to"] != nil {
		opts.ToLang = queryVal["to"][0]
	}

	var words[] Translate

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		words = append(words, Translate{
			Word:      line[0] ,
			Provision: line[1],
		})
	}

	js, err := json.Marshal(words)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
