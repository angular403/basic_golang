package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func main() {
	var andrew Customer
	andrew.Name = "Andrew Wiliam"
	andrew.Address = "Jakarta"
	andrew.age = 25

	fmt.Println(andrew)

	hafsah := Customer{
		Name:    "Hafsah Betari Andhika",
		Address: "Solo",
		age:     22,
	}
	fmt.Println(hafsah)

	ara := Customer{"Ara", "Medan", 23}
	fmt.Println(ara)
}
