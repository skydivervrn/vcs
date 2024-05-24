package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
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

	dogYAML, err := yaml.Marshal(dog)
	if err != nil {
		log.Fatal(err)
	}

	// How can you print `dogYAML` in a humanly-readable format?
	fmt.Println(string(dogYAML))
}
