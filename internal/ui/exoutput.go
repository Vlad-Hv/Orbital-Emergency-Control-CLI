package ui

import (
	"fmt"
	//"OStation/internal"
)

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

func PrintStorage(storage map[string]int) {
	for resourse, amount := range storage {
		fmt.Println("\nResourses:")
		fmt.Println(resourse, ":", amount)
	}
}
