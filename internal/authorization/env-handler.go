package authorization

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getLoginData() (string, int) {
	godotenv.Load("auth-data.env")

	login := os.Getenv("LOGIN")
	password := os.Getenv("PASSWORD")
	passwordInt := modifyData(password)

	return login, passwordInt
}

func modifyData(password string) int {
	passwordInt, _ := strconv.Atoi(password)

	return passwordInt
}
