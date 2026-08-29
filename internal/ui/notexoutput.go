package ui

import "fmt"

func printAuthInput() {
	fmt.Println("Please, entry login and password(int):")
}

func askMenuOption() {
	fmt.Println("Choose the option:")
}

func printMenu() {
	fmt.Println("\n----Operator Orbital Station Menu----")
	fmt.Println("1. Station status")
	fmt.Println("2. Move to another section")
	fmt.Println("4. Inventory / resources")
	fmt.Println("5. Mission log")
}

func printChooseRoom() {
	fmt.Println("Enter location ID:")
}
