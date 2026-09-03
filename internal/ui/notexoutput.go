package ui

import (
	"OStation/internal/orbital"
	"fmt"
)

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
	fmt.Println("3. Check hull zones")
	fmt.Println("4. Inventory / resources")
	fmt.Println("5. Mission log")
	fmt.Println("6. Send emergy signal SOS")
	fmt.Println("7. Check inventory")
}

func accessCard() {
	fmt.Println("\nYou found access card in the controll room")
}

func printChooseRoom() {
	fmt.Println("Enter location ID:")
}

func printStorageMenu() {
	fmt.Println("\n---Storage menu---")
	fmt.Println("1. Check all resourses")
	fmt.Println("2. Get some resourses")
	fmt.Println("3. Come back control room")
	fmt.Print("Choose one option:")
}

func askStorageTake() {
	fmt.Println("Enter what stuff and his amount you wanna get:")
}

func invMessage() {
	fmt.Println("\n====Inventory====")
}

func zoneMessage() {
	fmt.Println("===Zone list===")
}

func reactorMenu(reactor orbital.Zone) {
	fmt.Println("---Reactor Menu---")
	if reactor.Condition == "unstable" {
		fmt.Println("1. Reactor condition report\n2. Try to fix it\n3. Back to control room")
	} else {
		fmt.Println("1. Reactor condition report\n3. Back to control room")
	}
}

func communicationMenu(communication orbital.Zone) {
	fmt.Println("---Communication Menu---")
	if communication.Condition == "unstable" {
		fmt.Println("1 Communication condition report\n2. Try to fix it\n3. Back to control room")
	} else {
		fmt.Println("1. Communication condition \n3. Back to control room")
	}
}

func stationReport() {
	fmt.Println("===Station Condition===")
}

func lifeSupportMenu(lifeSup orbital.Zone) {
	fmt.Println("---Life Support---")

	if lifeSup.Condition == "unstable" {
		fmt.Println("1. Life Support condition report\n2. Try to fix it\n3. Back to control room")
	} else {
		fmt.Println("1.  Communication condition report\n3. Back to control room")
	}
}

func missionLog() {
	fmt.Println("---Mission Log---")
}

func emptyLog() {
	fmt.Println("history is empty")
}
