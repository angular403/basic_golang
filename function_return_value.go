package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}
func main() {
	//result := getHello("andrew")
	//fmt.Println(result)

	fmt.Println(getHello("andrew"))
	fmt.Println(getHello("wiliam"))
}
