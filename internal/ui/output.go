package ui

import "fmt"

func PrintAuth() {
	fmt.Println("---Authorization---")
}

func printAuthInput() {
	fmt.Println("Please, entry login and password(int):")
}

func InvalidAuth() {
	fmt.Println("incorrect login or password\n ")
}

func Access() {
	fmt.Println("Access applyed!")
}
