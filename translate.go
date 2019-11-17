package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var (
	reader *csv.Reader
)

type Translate struct {
	Word      string `json:"word"`
	Provision string `json:"provision"`
}

type TranslationOptions struct {
	Limit    int    `json:"limit"`
	FromLang string `json:"from_lang"`
	ToLang   string `json:"to_lang"`
}

// func init() {
// 	csvFile, err := os.Open("eng_tur.csv")
// 	if err != nil {
// 		panic("Can not open csv file")
// 	}
// 	reader = csv.NewReader(bufio.NewReader(csvFile))
// }

func (sv *Server) TranslationWords(w http.ResponseWriter, r *http.Request) {
	// var limit int
	// if limitStr := r.FormValue("limit"); limitStr == "" {
	// 	http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
	// 	return
	// } else {
	// 	var err error
	// 	limit, err = strconv.Atoi(limitStr)
	// 	if err != nil {
	// 		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
	// 		return
	// 	}
	// }
	//
	// from := r.FormValue("from")
	// if from == "" {
	// 	http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
	// 	return
	// }
	//
	// to := r.FormValue("to")
	// if to == "" {
	// 	http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
	// 	return
	// }
	//
	// opts := TranslationOptions{
	// 	Limit:    limit,
	// 	FromLang: from,
	// 	ToLang:   to,
	// }

	var words []Translate

	if _, err := sv.Csv.Seek(0, io.SeekStart); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reader = csv.NewReader(bufio.NewReader(sv.Csv))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if len(line) > 0 {
			words = append(words, Translate{
				Word:      line[0],
				Provision: line[1],
			})
		}
	}

	b, err := json.Marshal(words)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
