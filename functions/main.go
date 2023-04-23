package main

import "fmt"

// func double(x int) int {
// 	return x * 5
// }

// func add(lhs, rhs int) int {
// 	return lhs + rhs
// }

// func greet() {
// 	fmt.Println("Hello from greet functions")
// }

// func main() {
// 	greet()
// 	dozen := double(5)
// 	fmt.Println("A dozen is:", dozen)

// 	bakersDozen := add(dozen, 1)
// 	fmt.Println("A baker's dozen is:", bakersDozen)

// 	anotherBakersDozen := add(double(6), 1)
// 	fmt.Println("another baker's dozen is:", anotherBakersDozen)
// }

////////////////////////////////////////////////////////////////////////////////////
// type Coordinate struct {
// 	X, Y int
// }

// func shiftBy(x, y int, coord *Coordinate) {
// 	coord.X += x
// 	coord.Y += y
// 	fmt.Println(coord)
// }

// func (coord *Coordinate) shiftBy(x, y int) {
// 	coord.X += x
// 	coord.Y += y
// 	fmt.Println(coord)
// }

// func main() {
// 	coord := Coordinate{5, 5}
// 	coord.shiftBy(1, 1)
// 	shiftBy(1, 1, &coord)
// }

// /////////////////////////////////////////////////////////////////////
type Coordinate struct {
	X, Y int
}

func (c Coordinate) shfitDist(other Coordinate) Coordinate {
	return Coordinate{other.X - c.X, other.Y - c.Y}
}

func main() {
	first := Coordinate{2, 2}
	second := Coordinate{1, 5}
	distance := first.shfitDist(second)
	fmt.Println(distance)
}
