package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) float64, float64) {
	baseTemperature := 90.0

	var adjustTemperature = func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}

	return adjustTemperature, baseTemperature
}

func main() {
	var adjustTemperature, baseTemperature = createTemperatureAdjuster()
	fmt.Println(baseTemperature)
	fmt.Println(adjustTemperature(1.5))

}
