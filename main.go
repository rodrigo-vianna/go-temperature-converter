package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func convertKelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func convertCelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func main() {
	// Check if the user provided a temperature
	if len(os.Args) < 2 {
		fmt.Println("Please provide a temperature")
		os.Exit(1)
	}
	var temperature float64
	var toKelvin bool

	// Convert the temperature to a float64
	temperature, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Please provide a valid number")
		fmt.Println(err)
		os.Exit(1)
	}

	// Remove the first argument
	os.Args = os.Args[1:]

	// Parse the flags
	flag.BoolVar(&toKelvin, "k", false, "Convert to Kelvin")
	flag.Parse()

	// Convert the temperature
	var result float64
	var resultUnit rune
	if toKelvin {
		result = convertCelsiusToKelvin(temperature)
		resultUnit = 'K'
	} else {
		result = convertKelvinToCelsius(temperature)
		resultUnit = 'C'
	}

	// Print the result
	fmt.Printf("Temperature: %.2f %c\n", result, resultUnit)
}
