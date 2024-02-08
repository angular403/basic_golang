package main

import "fmt"

func sayHelloToProgrammer(FirstName string, LastName string, age int, hobby string) {
	fmt.Println("Name = ", FirstName, LastName, "\n", "Age = ", age, "\n", "Hobby = ", hobby)
}

func main() {
	sayHelloToProgrammer("Andrew", "Wiliam", 30, "Coding")
}
