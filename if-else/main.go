package main

import (
	"fmt"
	"strings"
)

func characters(s string) string {
	return strings.ToUpper(s)
}

func role(s string) string {
	if s == "Admin" {
		return "wow"
	}
	return s
}

func average(a, b, c int) float32 {
	return float32(a + b + c) / 3
}

func main() {
	// if syntax
	// if condition {
	// do something
	// }

	// if esle syntax
	// if condition {
	// 	// do something
	// } else {
	// do something
	// }

	firstName := "AmirAli"

	if firstName == "AmirAli" {
		fmt.Println("Yes :)")
	} else {
		fmt.Println("No :(")
	}

	if characters("amirali") == "AMIRALI" {
		fmt.Println(characters("amirali"), "=> AMIRALI")
	} else {
		fmt.Println(characters("amiarli"), "No upper :(")
	}

	a := 5

	if a != 5 {
		fmt.Println("a != 5")
	} else if a == 4 {
		fmt.Println("a == 4")
	} else {
		a = 10
		fmt.Println(a)
	}

	num1 := 2
	num2 := 5

	if num1 == 2 || num2 == 3 {
		fmt.Println("Or operator", num1, num2)
	} else if num1 == 2 && num2 == 5 {
		fmt.Println("And Operator")
	}

	if myRole := role("Admin"); myRole == "Admin" {
		fmt.Println("Welcome to dashboard :)")
	} else if myRole == "User" {
		fmt.Println("Welcome to our site")
	}

	var hello int = 123
	if hello == 12 {
		fmt.Println("this is hello world :)")
	} else if hello == 123 {
		fmt.Println("this is a hi", hello)
	} else {
		fmt.Println(role("user"))
	}

	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("quiz1 scored higher than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 scored higher than quiz1")
	} else {
		fmt.Println("quiz 1 & quize 2 have the smae score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("acceptable grades")
	} else {
		fmt.Println("instructor did a bad job!")
	}
}