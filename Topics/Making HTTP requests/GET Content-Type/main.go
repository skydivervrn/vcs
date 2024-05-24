package main

import (
	"fmt"
	"log"
	"net/http"
)

// Write the code to finish the implementation of `getContentType()` below:
func getContentType(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Access the `Header` field of `response` and pass "Content-Type" to `Get` below:
	return response.Header.Get("Content-Type")
}

// DO NOT modify the contents of the main() or must() functions!
func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"message": "Hello, JSON!"}`)
	})

	http.HandleFunc("/wiki", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Welcome to the JetBrains Academy Wiki!</h1>")
	})

	go func() {
		must(http.ListenAndServe(":8080", nil))
	}()

	var url string
	fmt.Scan(&url)
	fmt.Println(getContentType(url))
}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
