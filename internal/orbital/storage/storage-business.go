package storage

func TakeResource(storage map[string]int, inventory map[string]int, resource string, amount int) {
	storage[resource] -= amount
	inventory[resource] += amount
}

func StorageCheck(stationStorage map[string]int, resource string) {
	if stationStorage[resource] == 0 {
		delete(stationStorage, resource)
	}
}

func FixZone(requiredStuff map[string]int, inventory map[string]int) {
	for resource := range requiredStuff {
		inventory[resource] -= requiredStuff[resource]

		if inventory[resource] == 0 {
			delete(inventory, resource)
		}

		delete(requiredStuff, resource)
	}
}

func AccessCard(step int, inventory map[string]int) {
	if step >= 4 && step <= 7 {
		inventory["accessCard"] = 1
	}
}

func SpecTool(inventory map[string]int) {
	inventory["tool"] += 1
}
