package main

import "fmt"

func main() {
	switch age := 0; {
	case age == 0:
		fmt.Println("newborn")
	case age >= 1 && age <= 3:
		fmt.Println("toddler")
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
}