package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string
	Breed string
	Age   int
}

func main() {
	// DO NOT modify the code block below:
	var name, breed string
	var age int
	fmt.Scanln(&name, &breed, &age)

	dog := Dog{Name: name, Breed: breed, Age: age}

	dogJSON, err := json.Marshal(dog)
	if err != nil {
		log.Fatal(err)
	}

	// How can you print `dogJSON` in a humanly-readable format?
	fmt.Println(string(dogJSON))
}
