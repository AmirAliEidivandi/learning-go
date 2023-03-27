package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office", cleaned: true},
		{name: "Home"},
		{name: "WC", cleaned: true},
		{name: "Warhouse"},
	}

	checkCleanliness(rooms)

	rooms[1].cleaned = true
	rooms[3].cleaned = true
	checkCleanliness(rooms)

	// myArray := [5]int{1, 2, 3, 4, 5}
	// var myArray [3]int
	// myArray := [3]int{201, 325, 4322}
	// myArray := [...]int{10, 3, 32, 8127, 4329278}
	myArray := [...]string{"hi", "bye"}
	fmt.Println(myArray)

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}
