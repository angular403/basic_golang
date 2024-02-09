package main

import "fmt"

func getFullName() (string, string) {
	return "andrew", "wiliam"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
