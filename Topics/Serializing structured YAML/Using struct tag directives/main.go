package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"log"
)

// Add the correct YAML struct tags.
type Engine struct {
	Type  string `yaml:"engine_type"`
	Power int    `yaml:"max_horsepower"`
	Turbo bool   `yaml:"turbo"`
}

type Car struct {
	Make   string  `yaml:"make"`
	Model  string  `yaml:"model"`
	Year   int     `yaml:"year"`
	Price  float32 `yaml:"-"`
	Engine Engine  `yaml:",inline"`
}

// DO NOT change the code within the main function! - Your task is only to add the correct YAML struct tags above!
func main() {
	var carMake, carModel, carEngineType string
	var carYear, carPower int
	var carPrice float32
	var turbo bool

	fmt.Scanln(&carMake, &carModel, &carYear, &carPrice, &carEngineType, &carPower, &turbo)

	engine := Engine{
		Type:  carEngineType,
		Power: carPower,
		Turbo: turbo,
	}

	car := Car{
		Make:   carMake,
		Model:  carModel,
		Year:   carYear,
		Engine: engine,
		Price:  carPrice,
	}

	carYAML, err := yaml.Marshal(car)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(string(carYAML))
}
