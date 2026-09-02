package ui

import (
	"OStation/internal/orbital"
	"fmt"
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

func ControlRoom() {
	fmt.Println("You are already in control room")
}

func AllZones(zones map[int]*orbital.Zone) {
	zoneMessage()
	for _, zone := range zones {
		fmt.Println("Name:", zone.Name)
		fmt.Println("ID:", zone.ID, "\n ")
	}
}

func ReactorReport(reactor orbital.Zone) {
	fmt.Print("\nReactor Condition:")
	fmt.Println(reactor.Condition)
	fmt.Println("\nStuff to fix:")
	for resourse, amount := range reactor.StuffToFix {
		fmt.Println(resourse, ":", amount)
	}
}

func ReactorStableReport(reactor orbital.Zone) {
	fmt.Print("\nReactor Condition:")
	fmt.Println(reactor.Condition)
}

func SuccesFixed(name string) {
	fmt.Println(name, " fixed successfully!")
}
