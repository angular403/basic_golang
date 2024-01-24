package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Andrew"
	names[1] = "Wiliam"
	names[2] = "Napitupulu"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{90, 80, 99,
		200, 300, 400}
	fmt.Println(len(values))

}
