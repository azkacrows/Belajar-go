package main

import "fmt"

func main() {
	// var arr [5]int
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// arr[3] = 4
	// arr[4] = 5
	// fmt.Println(arr)

	var fruits = [4]string{"mangga", "jambu", "pisang", "kelopo"}
	fmt.Println("jumlah buah", len(fruits))
	fmt.Println("isi semua buah", fruits)
	fmt.Println("isi buah ke-3", fruits[2])
}
