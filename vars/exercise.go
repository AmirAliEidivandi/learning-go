package main

import "fmt"

func main() {
	var myFavoriteColor string = "red"
	myAge, yearsOfBirth := 21, 2002
	var (
		// firstInitial = 'A' // 65
		// lastInitial  = 'I' // 73
		firstInitial = "A" // A
		lastInitial  = "I" // I
	)
	var calculateAge int
	calculateAge = 365 * myAge
	fmt.Println("my favorite color:", myFavoriteColor)
	fmt.Println("My age:", myAge, "years of birth:", yearsOfBirth)
	fmt.Println("Initials=", firstInitial, lastInitial)
	fmt.Println(calculateAge)
}
