package main

import "fmt"

func main() {
	slice := []string{"Hello", "world", "!"}

	// for i := 0; i < 10; i++ {
	// 	slice[0][i]
	// }

	for i, element := range slice {
		fmt.Println(i, element, ":")

		for _, ch := range element {
			fmt.Printf("    %q\n", ch)
		}
	}
}