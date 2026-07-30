package main

import (
	"fmt"

	"github.com/kristiansroberts/dvd-game-launcher/internal/scan"
)

func main() {
	games, err := scan.DiscoverEXEs()
	if err != nil {
		fmt.Println("Error discovering EXEs:", err)
		fmt.Println("Ensure that the program is run from the correct directory structure.")
		return
	}
	for _, game := range games {
		fmt.Println("Found game:", game.Name, "at", game.ExePath)
	}

}
