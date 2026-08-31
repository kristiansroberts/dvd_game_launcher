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
	icon, err := fyne.LoadResourceFromPath("./icon.ico")
	if err == nil {
		a.SetIcon(icon)
	}
	w := a.NewWindow("Dvd Game Launcher")

	_, err = showMainMenu(w)
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

// shows the extras menu with buttons for each extra item, includes a back button to return to the main menu
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
	for _, extraItem := range extraItems {
		// construct a "sub-folder" in the app for the extra item based on its type, organizing it regardless of file structure
		extraType := scan.DetermineExtraItemType(filepath.Base(filepath.Dir(extraItem.FilePath)), extraItem.Name)
		label := extraType.String()
		btn := widget.NewButton(label, func() {
			if err := showExtrasTypeMenu(w, extraType); err != nil {
				widget.NewLabel(fmt.Sprintf("Error launching extra type menu: %v", err))
			}
		})
		buttons = append(buttons, btn)
	}

	w.SetContent(container.NewVBox(buttons...))
	return nil
}

func showExtrasTypeMenu(w fyne.Window, extraType config.ExtraType) error {
	// Placeholder implementation for showing extras of a specific type
	w.SetContent(widget.NewLabel(fmt.Sprintf("Showing extras of type: %s", extraType.String())))
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
