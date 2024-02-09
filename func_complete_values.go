package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Andrew"
	middleName = "Wiliam"
	lastName = "Napitupulu"

	return firstName, middleName, lastName
}

func main() {
	A, B, C := getCompleteName()
	fmt.Println(A, B, C)
}
