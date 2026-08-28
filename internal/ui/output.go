package ui

import "fmt"

func PrintAuth() {
	fmt.Println("---Authorization---")
}

func printAuthInputLog() {
	fmt.Println("\nPlease, enter login:")
}

func printAuthInputPassw() {
	fmt.Println("Well, enter password(int):")
}

func InvalidAuth() {
	fmt.Println("\nincorrect login or password\n ")
}

func Access() {
	fmt.Println("Access applyed!")
}
