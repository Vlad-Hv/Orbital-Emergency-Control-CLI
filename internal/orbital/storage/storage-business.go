package storage

func TakeResourse(storage *map[string]int, inventory *map[string]int, resourse string, amount int) {
	(*storage)[resourse] -= amount
	(*inventory)[resourse] += amount
}
