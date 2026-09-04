package authorization

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getLoginData() (string, int, error) {
	err := godotenv.Load("auth-data.env")

	if err != nil {
		return "", 0, err
	}
	login := os.Getenv("LOGIN")
	password := os.Getenv("PASSWORD")
	passwordInt, err := modifyData(password)

	return login, passwordInt, err
}

func modifyData(password string) (int, error) {
	passwordInt, err := strconv.Atoi(password)

	return passwordInt, err
}
