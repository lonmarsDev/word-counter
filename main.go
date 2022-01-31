package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	listenPort string = ":8081"
)

func main() {
	http.HandleFunc("/", wordHandler)
	log.Println("listen on", listenPort)

	s := &http.Server{
		Addr:           listenPort,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

type Wordlist struct {
	Word  string
	Count int
}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	filteredWord := workProcessor(reqBody)
	jsonData, err := json.Marshal(filteredWord)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func workProcessor(wordsByte []byte) []Wordlist {

	wordStorage := make(map[string]int)

	words := strings.Fields(string(wordsByte))

	for _, v := range words {
		if _, ok := wordStorage[v]; ok {
			wordStorage[v]++
			continue
		}
		wordStorage[v] = 1
	}

	wl := make([]Wordlist, 0)
	var loopCount int = 0
	for k, v := range wordStorage {
		wl = append(wl, Wordlist{k, v})
		loopCount++

	}
	sort.Slice(wl, func(i, j int) bool {
		return wl[i].Count > wl[j].Count
	})

	return wl[:10]

}
