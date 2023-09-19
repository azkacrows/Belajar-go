package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	person1 := person{
		name: "jamal",
		age:  19,
	}
	fmt.Println("nama:", person1.name)
	fmt.Println("umur:", person1.age)
}
