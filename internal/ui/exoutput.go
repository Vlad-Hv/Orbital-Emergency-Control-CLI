package ui

import "fmt"

func PrintAuth() {
	fmt.Println("---Authorization---")
}

func InvalidAuth() {
	fmt.Println("incorrect login or password\n ")
}

func Access() {
	fmt.Println("Access applyed!")
}

func FirstWarning() {
	fmt.Println(" --------------------------------")
	fmt.Println("|                                |")
	fmt.Println("|           🚨WARNING🚨          |")
	fmt.Println("|      Critical reactor state    |")
	fmt.Println("|                                |")
	fmt.Println(" --------------------------------")
}
