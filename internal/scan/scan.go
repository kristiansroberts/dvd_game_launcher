package scan

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/kristiansroberts/dvd-game-launcher/internal/config"
)

func collectionRoot() (string, error) {
	exePath, err := os.Executable()
	// retrieve the path on the current executable, hinges on the matching the correct dvd file structure, so we need to check for errors here
	if err != nil {
		return "", err
	}

	return filepath.Dir(exePath), nil
}

func findEXEs(dir string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.EqualFold(filepath.Ext(f.Name()), ".exe") {
			return filepath.Join(dir, f.Name()), nil
		}
	}
	return "", nil
}

func DiscoverGames() ([]config.Game, error) {
	root, err := collectionRoot()
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var games []config.Game
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		gameDir := filepath.Join(root, e.Name())
		exePath, err := findEXEs(gameDir)
		if err != nil {
			return nil, err
		}
		if exePath != "" {
			games = append(games, config.Game{
				Name:    e.Name(),
				Dir:     gameDir,
				ExePath: exePath,
			})
		}
	}
	return games, nil
}

// search for an extras directory in the root of the collection and create an object
func DiscoverExtras() (config.Extras, error) {
	root, err := collectionRoot()
	if err != nil {
		return config.Extras{}, err
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return config.Extras{}, err
	}

	var extra config.Extras
	for _, e := range entries {
		if e.IsDir() && strings.EqualFold(e.Name(), "extras") {
			extraDir := filepath.Join(root, e.Name())
			extra = config.Extras{
				Name: e.Name(),
				Dir:  extraDir,
			}
			return extra, nil
		}

	}
	return config.Extras{}, errors.New("No extras directory found")
}

// helper to determine what type an extra subdirectory is based on its name
// func determineExtraItemType(dir string) string {
// 	// Check for specific subdirectories to determine the type of extra
// 	subDirs := map[string]string{
// 		strings.ToLower("art"):    "Art",
// 		strings.ToLower("ost"):    "OST",
// 		strings.ToLower("manual"): "Manual",
// 		strings.ToLower("video"):  "Video",
// 	}
// 	for subDir, extraType := range subDirs {
// 		if _, err := os.Stat(filepath.Join(dir, subDir)); err == nil {
// 			return extraType
// 		}
// 	}
// 	return "Unknown"
// }
