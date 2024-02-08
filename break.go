package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan Ke-", i)
	}

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println("Perulangan Ke-", i)

	}
}
