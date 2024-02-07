package main

import "fmt"

func main() {
	name := "andrew"
	switch name {
	case "andrew":
		fmt.Println("Hello Andrew")
	case "wiliam":
		fmt.Println("Hello Wiliam")
	default:
		fmt.Println("Hello, Boleh Kenalan ? ")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Lumayan Panjang")
	}

	name = "andrew napitupulu"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlau panjang")
	case length > 5:
		fmt.Println("nama Lumayan panjang")
	default:
		fmt.Println("nama Sudah Benar")

	}
}
