package main

import (
	"OStation/internal/authorization"
	"OStation/internal/ui"
	"fmt"
)

func main() {
	err := authorization.Auth()

	if err != nil {
		fmt.Println(err)
		return
	}

	ui.Access()
}
