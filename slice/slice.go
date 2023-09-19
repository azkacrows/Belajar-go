package main

import "fmt"

func main() {
	var fruitA = []string{"mangga", "jambu", "pisang", "kelopo"} //slice
	var fruitB = [2]string{"anggur", "alpukat"}                  //array
	var fruitC = [...]string{"pepaya", "kelopo"}                 //array
	var fruitD = make([]string, 3, 5)                            //array
	fmt.Println(fruitA, fruitB, fruitC, fruitD)
}
