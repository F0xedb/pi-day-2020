package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type piIndex struct {
	Index       int
	Input       string
	Surrounding string
}

type piResult struct {
	Index       int
	Surrounding string
}

var piDigits string

// speed in bytes
var speed = 100000

func main() {
	http.HandleFunc("/search/", SearchServer)
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	var filename = "pi.txt"
	if len(os.Args[1:]) > 0 {
		filename = os.Args[1]
	}
	loadFile(filename)
	http.ListenAndServe(":8080", nil)
}

// streamline type checks
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadFile(filename string) {
	file, err := os.Open(filename)
	check(err)
	buffer := make([]byte, speed)
	piDigits = ""
	for {
		size, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		check(err)
		piDigits += string(buffer[:size])
	}
	fmt.Println("Pi Digits loaded in!")

}

// Get the index in pi of a given string
func getIndex(index string) piResult {
	var position = strings.Index(piDigits[2:], index) + 1
	if position <= 0 {
		position = -1
	}
	return piResult{position, piDigits[position+1 : position+11]}
}

// Webserver endpoints are beeing served here

// EchoServer default / endpoint echoing to parameters of the URL back
func EchoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Echo: %s", r.URL.Path[1:])
}

// SearchServer This function is served on /search/<num> it retrieves the index found in pi.
// SearchServer With a resolution of 1 Billion
func SearchServer(w http.ResponseWriter, r *http.Request) {
	var splittedPath = strings.Split(r.URL.Path, "/")
	var index = splittedPath[len(splittedPath)-1]
	var result = getIndex(index)
	piResult := piIndex{result.Index, index, result.Surrounding}
	js, err := json.Marshal(piResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
