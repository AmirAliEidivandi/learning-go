package main

import "fmt"

func twoTwos() (int, int) {
	return 2, 2
}

func greetings(name string) {
	fmt.Println("Hello", name)
}

func add(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func anyNumber() int {
	return 5
}

func hiThere() string {
	return "Hi There"
}

func main() {
	greetings("AmirAli")
	fmt.Println(hiThere())
	a, b := twoTwos()
	answer := add(anyNumber(), a, b)
	fmt.Println(answer)
}
