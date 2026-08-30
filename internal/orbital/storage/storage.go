package storage

func CreateStorage() map[string]int {
	return map[string]int{
		"metal":   13,
		"fuel":    5,
		"wire":    6,
		"medical": 3,
	}
}

func CreateInventory() map[string]int {
	mapInventory := make(map[string]int)
	return mapInventory
}
