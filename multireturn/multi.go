package main

import (
	"fmt"
	"math"
)

func main() {
	area, circum := calculate(3.5)
	fmt.Println("area: ", area)
	fmt.Println("circum: ", circum)
	fmt.Println(area, circum)
}

func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}
