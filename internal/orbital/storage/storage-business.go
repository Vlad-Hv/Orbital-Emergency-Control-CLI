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

func FixZone(zoneStuff *map[string]int, inventory *map[string]int) {
	for resourse := range *zoneStuff {
		(*inventory)[resourse] -= (*zoneStuff)[resourse]

		if (*inventory)[resourse] == 0 {
			delete((*inventory), resourse)
		}

		delete((*zoneStuff), resourse)
	}
}

func AccessCard(step int, inventory *map[string]int) {
	if step >= 4 && step <= 7 {
		(*inventory)["accessCard"] = 1
	}
}
