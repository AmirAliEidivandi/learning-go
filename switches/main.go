// switch is used to easily check multiple conditions
// Alternative to many if-else if blocks
// Switches execute from top-to-bottem
// optionally have a "default" action
package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func calculate(num int) int {
	return 3 + num
}

func main() {
	// basic example:
	x := 5

	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("other:", x)
	}

	url := "example.com"
	switch url {
	case "example.com":
		fmt.Println("test")
	case "google.com":
		fmt.Println("live")
	default:
		fmt.Println("dev")
	}

	// conditional cases
	switch result := calculate(5); {
	case result > 10:
		fmt.Println(">10")
	case result == 6:
		fmt.Println("==6")
	case result < 10:
		fmt.Println("<10")
	}

	// case list
	switch x {
	case 1, 2, 3:
		fmt.Println("1 or 2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}

	// fallthrough
	// fallthrough will continue checking the next case
	// letter := "hello amirali"
	// switch letter {
	// case ' ':
	// case 'a', 'b', 'm', 'r':
	// 	fmt.Println("a vowel")
	// 	fallthrough
	// case 'A', 'B', 'M', 'R':
	// 	fmt.Println("Vowel are great")
	// default:
	// 	fmt.Println("It's something else")
	// }

	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("Moderately priced item")
	default:
		fmt.Println("expensive item")
	}

	ticket := Economy

	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("First class seating")
	default:
		fmt.Println("other seating")
	}
}
