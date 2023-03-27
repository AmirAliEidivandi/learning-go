package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}
	printSlice("Route 1", route)

	route = append(route, "Home")
	printSlice("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[2], "visited")

	route = route[2:]
	printSlice("Remaining locations", route)

	// mySlice := []int{1, 2, 3, 4}
	// item1 := mySlice[0]
	// fmt.Println("item1:", item1) // 1

	// numbers := [...]int{1, 2, 3, 4, 5}

	// slice1 := numbers[:]
	// fmt.Println("slice1:", slice1) // [1 2 3 4]

	// slice2 := numbers[1:]
	// fmt.Println("slice2:", slice2) // [2 3 4 5]

	// slice3 := numbers[:2]
	// fmt.Println("slice3:", slice3) // [1 2]

	// slice4 := numbers[2:4]
	// fmt.Println("slice4:", slice4) // [3 4]
}
