package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("rata-rata: %2f", avg)
	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, nunumbers := range numbers {
		total += nunumbers
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}
