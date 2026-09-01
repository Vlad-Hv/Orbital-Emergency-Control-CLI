package storage

func TakeResourse(storage *map[string]int, inventory *map[string]int, resourse string, amount int) {
	(*storage)[resourse] -= amount
	(*inventory)[resourse] += amount
}

func StorageCheck(stationStorage *map[string]int, resourse string) {
	if (*stationStorage)[resourse] == 0 {
		delete((*stationStorage), resourse)
	}
}
