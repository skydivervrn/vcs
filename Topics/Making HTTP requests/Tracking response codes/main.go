package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const url = "http://127.0.0.1:8080/api"

func main() {
	// Write your code here
	for {
		if getMust() {
			break
		}
		continue
	}
}

func getMust() bool {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() // don't forget to close body
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
		return true
	}
	return false
}

// DO NOT MODIFY the contents of the init() or must() functions!
func init() {
	const timeout = 5 * time.Second

	var input []struct {
		Code int    `json:"code"`
		Body string `json:"body"`
	}

	must(json.NewDecoder(os.Stdin).Decode(&input))

	go func() {
		var count int
		http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
			response := input[count%len(input)]
			w.WriteHeader(response.Code)
			if _, err := w.Write([]byte(response.Body)); err != nil {
				panic(err)
			}
			count++
		})
		must(http.ListenAndServe(":8080", nil))
	}()

	time.Sleep(timeout)
}

// DO NOT delete or modify the must() function; it is a helper to check for errors!
func must(err error) {
	if err != nil {
		panic(err)
	}
}
