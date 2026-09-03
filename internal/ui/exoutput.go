package ui

import (
	history "OStation/internal/eventhistory"
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
	fmt.Println(name, "fixed successfully!")
}

func CommunicationReport(communication orbital.Zone) {
	fmt.Print("\nCommunication Condition: ")
	fmt.Println(communication.Condition)
	fmt.Println("\nStuff to fix:")
	for resourse, amount := range communication.StuffToFix {
		fmt.Println(resourse, ":", amount)
	}
}

func CommunicationStable(communication orbital.Zone) {
	fmt.Print("\nCommunication Condition: ")
	fmt.Println(communication.Condition)
}

func LeftRoom() {
	fmt.Println("Come back to the controll room!")
}

func StationReport(station orbital.OrbitalStation) {
	stationReport()
	fmt.Println("CurrentZone:", station.CurrentZone.Name, station.CurrentZone.ID)
	fmt.Println("Energy:", station.Energy)
	fmt.Println("Oxygen:", station.Oxygen)
}

func SignalSucces() {
	fmt.Println("Signal Sent Successfully!")
}

func EnterFromMenu() {
	fmt.Println("Enter from Control room")
}

func LifeSupport(lifeZone orbital.Zone) {
	fmt.Println("\nLife Support Condition:", lifeZone.Condition)
	fmt.Println("\nStuff to fix:")
	for resourse, amount := range lifeZone.StuffToFix {
		fmt.Println(resourse, ":", amount)
	}
}

func LifeStableSupport(lifeZone orbital.Zone) {
	fmt.Println("\nLife Support Condition:", lifeZone.Condition)
}

func SpesToolNotification() {
	fmt.Println("You found the special tool for Life Support zone forward the Reactor")
}

func HistoryReport(history history.History) {
	missionLog()
	if len(history) == 0 {
		emptyLog()
		return
	}
	for _, message := range history {
		fmt.Println(message)
	}
}
