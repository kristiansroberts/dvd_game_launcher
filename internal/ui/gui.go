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
	"github.com/kristiansroberts/dvd-game-launcher/internal/config"
	"github.com/kristiansroberts/dvd-game-launcher/internal/scan"
)

func Gui() {
	a := app.New()
	w := a.NewWindow("Dvd Game Launcher")

	_, err := showMainMenu(w)
	if err != nil {
		return
	}

	// w.SetContent(container.NewVBox(buttons...))
	w.ShowAndRun()
}

// makes buttons for each discovered game, right now rendered with text. function separated for better menu navigation
func showMainMenu(w fyne.Window) ([]fyne.CanvasObject, error) {
	games, err := scan.DiscoverGames()
	if err != nil {
		w.SetContent(widget.NewLabel("Error discovering games: " + err.Error()))
		w.ShowAndRun()
		return nil, err
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

	extras, err := scan.DiscoverExtras()
	if err != nil {
		w.SetContent(widget.NewLabel("Error discovering extras: " + err.Error()))
		w.ShowAndRun()
		return nil, err
	}
	if extras.Name != "" {
		extrasBtn := widget.NewButton("Extras", func() {
			if err := showExtrasMenu(w, extras); err != nil {
				widget.NewLabel(fmt.Sprintf("Error launching extra: %v", err))
			}
		})
		buttons = append(buttons, extrasBtn)
	}

	w.SetContent(container.NewVBox(buttons...))
	return buttons, nil
}

func showExtrasMenu(w fyne.Window, e config.Extras) error {
	var buttons []fyne.CanvasObject

	backButton := widget.NewButton("Back", func() {
		showMainMenu(w)
	})
	buttons = append(buttons, backButton)

	extraItems, err := scan.DiscoverExtraItems(e.Dir)
	if err != nil {
		w.SetContent(widget.NewLabel("Error discovering extra items: " + err.Error()))
		w.ShowAndRun()
		return err
	}
	for _, item := range extraItems {
		extraItem := item
		label := displayNameFromExe(extraItem.FilePath)
		btn := widget.NewButton(label, func() {
			if err := actions.LaunchEXE(extraItem.FilePath); err != nil {
				widget.NewLabel(fmt.Sprintf("Error launching extra item: %v", err))
			}
		})
		buttons = append(buttons, btn)
	}

	w.SetContent(container.NewVBox(buttons...))
	return nil
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
