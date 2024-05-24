package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	name := "user"
	response, err := http.Post("http://127.0.0.1:8080/user", "text/plain", bytes.NewBufferString(name))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

// DO NOT MODIFY the contents of the init() or must() functions!
func init() {
	const timeout = 5 * time.Second

	var input string
	fmt.Scan(&input)

	go func() {
		http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
			all, err := io.ReadAll(r.Body)
			must(err)
			if r.Method != http.MethodPost || len(all) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if _, err = w.Write([]byte(input)); err != nil {
				panic(err)
			}
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
