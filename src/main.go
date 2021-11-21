package main

import (
	"fmt"

	game "./game"
)

func main() {
	fmt.Println("INIT_MAIN")
	//game.Client()
	var s string
	fmt.Scanln(&s)
	if s == "1" {
		game.Server()

	} else {
		game.Client()
	}
}
