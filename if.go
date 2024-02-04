package main

import "fmt"

func main() {
	name := "andrew"
	if name == "andrew" {
		fmt.Println("Hello Andrew")
	} else if name == "raden" {
		fmt.Println("Hello Raden")
	} else {
		fmt.Println("Boleh Kenalan")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

}
