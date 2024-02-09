package main

import "fmt"

type Filter func(string) string

func spamChatWithFilter(name string, filter Filter) {
	filteredname := filter(name)
	fmt.Println("Hello", filteredname)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	} else {
		return name
	}
}

func main() {
	spamChatWithFilter("Andrew", spamFilter)

	filter := spamFilter
	spamChatWithFilter("Anjing", filter)
}
