package main

import "fmt"

func main() {

	a := "Anderson"

	switch "Value" {
	case "Bob":
		fmt.Println("Hey Bob")
	case "Maria":
		fmt.Println("Hey Maria")
	default:
		fmt.Println("Hey Bob")
	}
}
