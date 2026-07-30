package scan

import (
	"os"
	"path/filepath"
	"strings"
)

type Game struct {
	Name    string
	Dir     string
	ExePath string
}

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

func DiscoverEXEs() ([]Game, error) {
	root, err := collectionRoot()
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var games []Game
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
			games = append(games, Game{
				Name:    e.Name(),
				Dir:     gameDir,
				ExePath: exePath,
			})
		}
	}
	return games, nil
}
