package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("unhandled operation")
}

func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2))

	mul := Operation(Multiply)
	fmt.Println(mul.calculate(10, 3))

	div := Operation(Divide)
	fmt.Println(div.calculate(5, 5))

	sub := Operation(Subtract)
	fmt.Println(sub.calculate(50, 2))
}
