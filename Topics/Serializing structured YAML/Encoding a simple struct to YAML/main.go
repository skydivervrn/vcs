package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"log"
)

// DO NOT delete or modify the Car struct implementation:
type Car struct {
	Make  string `yaml:"make"`
	Model string `yaml:"model"`
	Year  int    `yaml:"year"`
	Color string `yaml:"color,omitempty"`
}

func main() {
	// DO NOT delete! - This code block takes as an input the values for the 'car' struct:
	var _make, model, color string
	var year int
	fmt.Scanln(&_make, &model, &year, &color)

	car := Car{
		Make:  _make,
		Model: model,
		Year:  year,
		Color: color,
	}

	// Write the code below to properly encode the 'car' struct into the 'carYAML' variable:
	data, err := yaml.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}

	// Print the serialized 'carYAML' below; remember to cast it as a 'string'!
	fmt.Println(string(data))
}
