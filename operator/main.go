package main

import "fmt"

func main() {
	var total = 3 + 3
	fmt.Println(total)

	a := 5
	a = 10
	fmt.Println(a)

	var firstName = "Amirali"
	_ = firstName

	lastName := "Eidivandi"
	_ = lastName

	var (
		firstName2 = 'A'
		lastName2  = 'E'
	)

	var hh = 5
	fmt.Println(hh)

	fmt.Println(firstName2, lastName2)

	b := 10
	b = b + 10
	fmt.Println(b)

	fmt.Println(b > a)  // true
	fmt.Println(b < a)  // false
	fmt.Println(b >= a) // true
	fmt.Println(b <= a) // false
	fmt.Println(b == a) // false
	fmt.Println(b != a) // true

	// login operator
	// || = Or - && = And - ! = Not
	// Always returns a boolean (true or false)
}
