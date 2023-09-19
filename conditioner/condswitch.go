package main

import "fmt"

func main(){
	var point = 10
	switch point {
	case 10:
		fmt.Println("lulus dengan nilai sempurna")
	case 5, 6, 7:
		fmt.Println("lulus")
	case 4:
		fmt.Println("hampir lulus")
	default:
		fmt.Println("tidak lulus")
	}
}