package main

import "fmt"

type Temperature float64

// nolint: gomnd // DO NOT remove this comment!
func (t Temperature) CelsiusToFahrenheit() Temperature {
	return t*9/5 + 32
}

// nolint: gomnd // DO NOT remove this comment!
func (t Temperature) FahrenheitToKelvin() Temperature {
	return t + 273.15
}

func (t Temperature) AddTemperature(other Temperature) Temperature {
	return t + other
}

func main() {
	var t1 Temperature
	fmt.Scan(&t1)

	t2 := t1.AddTemperature(t1)
	t3 := t2.CelsiusToFahrenheit()
	t4 := t3.FahrenheitToKelvin()
	fmt.Println(t4)
}
