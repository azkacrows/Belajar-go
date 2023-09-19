package main

import "fmt"

func main() {
	var ayam map[string]int
	ayam = map[string]int{}
	ayam["bakar"] = 100
	ayam["goreng"] = 50
	fmt.Println("bakar", ayam["bakar"])
	fmt.Println("goreng", ayam["goreng"])
	// fmt.Println("madu", ayam["madu"]) output : 0
}
