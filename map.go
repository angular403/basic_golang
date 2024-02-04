package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "andrew",
		"age":     "30",
		"address": "jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	book := make(map[string]string)

	book["title"] = "Golang"
	book["author"] = "andrew ganteng"
	book["wrong"] = "map salah"

	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)

}
