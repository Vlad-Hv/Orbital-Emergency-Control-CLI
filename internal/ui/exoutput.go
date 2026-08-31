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
	fmt.Println("\nResourses:")
	for resourse, amount := range storage {
		fmt.Println(resourse, ":", amount)
	}
}

func PrintInv(inv map[string]int) {
	invMessage()
	if len(inv) == 0 {
		fmt.Println("inventory is empty")
		return
	}

	for resourse, amount := range inv {
		fmt.Println(resourse, ":", amount)
	}
}

func InvalidRoom() {
	fmt.Println("invalid room id")
}
