package ui

import (
	"fmt"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kristiansroberts/dvd-game-launcher/internal/actions"
	"github.com/kristiansroberts/dvd-game-launcher/internal/scan"
)

func Gui() {
	a := app.New()
	w := a.NewWindow("Dvd Game Launcher")

	games, err := scan.DiscoverGames()
	if err != nil {
		w.SetContent(widget.NewLabel("Error discovering games: " + err.Error()))
		w.ShowAndRun()
		return
	}

	var buttons []fyne.CanvasObject
	for _, g := range games {
		game := g
		label := displayNameFromExe(game.ExePath)
		btn := widget.NewButton(label, func() {
			if err := actions.LaunchEXE(game.ExePath); err != nil {
				widget.NewLabel(fmt.Sprintf("Error launching game: %v", err))
			}
		})
		buttons = append(buttons, btn)
	}

	w.SetContent(container.NewVBox(buttons...))
	w.ShowAndRun()
}

// helper function to generate a display name from an executable path, remove the file extension and capitalize the first letter
func displayNameFromExe(exePath string) string {
	base := filepath.Base(exePath)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	if name == "" {
		return name
	}
	return strings.ToUpper(name[:1]) + name[1:]
}
