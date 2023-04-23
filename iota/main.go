package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	// switch d {
	// case North:
	// 	return fmt.Sprintf("North")
	// case East:
	// 	return fmt.Sprintf("East")
	// case South:
	// 	return fmt.Sprintf("hey South")
	// case West:
	// 	return fmt.Sprintf("West")
	// default:
	// 	return "other direction"
	// }

	return []string{"North", "East", "South", "West"}[d]
}

func main() {
	north := North
	fmt.Println(north)
}
