package main

import (
	"fmt"
	"slices"

	"github.com/kristiansroberts/dvd-game-launcher/internal/actions"
	"github.com/kristiansroberts/dvd-game-launcher/internal/scan"
)

const DEBUG = true

func main() {
	games, err := scan.DiscoverEXEs()
	var gameslist []string
	if err != nil {
		fmt.Println("Error discovering EXEs:", err)
		fmt.Println("Ensure that the program is run from the correct directory structure.")
		return
	}
	for _, game := range games {
		fmt.Println("Found game:", game.Name, "at", game.ExePath)
		gameslist = append(gameslist, game.Name)
	}

	fmt.Println("Please type the name of the game you want to launch: ")
	var selectedGame string
	fmt.Scanln(&selectedGame)

	if !slices.Contains(gameslist, selectedGame) {
		fmt.Println("Game not found in the list. Please check the name and try again.")
		return
	}

	for _, game := range games {
		if game.Name == selectedGame {
			err = actions.LaunchEXE(game.ExePath)
			if err != nil {
				fmt.Println("Error launching the game:", err)
				return
			}
			break
		}
	}
}
