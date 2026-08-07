package ui

import (
	"fmt"
	"slices"

	"github.com/kristiansroberts/dvd-game-launcher/internal/actions"
	"github.com/kristiansroberts/dvd-game-launcher/internal/scan"
)

func CliDisplayGamesAndExtras() {
	games, err := scan.DiscoverGames()
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
	extraFolder, err := scan.DiscoverExtras()
	if err == nil {
		extras, err := scan.DiscoverExtraItems(extraFolder.Dir)
		if err != nil {
			fmt.Println("Error discovering extra items:", err)
		}
		for _, extra := range extras {
			fmt.Println("Found extra:", extra.Name, "of type", extra.ExtraType.String(), "at", extra.FilePath)
			gameslist = append(gameslist, extra.Name)
		}
	} else {
		fmt.Println("Error discovering extras:", err)
	}

	fmt.Println("Please type the name of the game or extra you want to launch: ")
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
