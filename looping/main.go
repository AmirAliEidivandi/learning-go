package main

import "fmt"

func main() {
	// for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// foe: while
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("sum is -5:", sum)
	}
}
