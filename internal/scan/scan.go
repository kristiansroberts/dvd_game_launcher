package scan

import (
	"errors"
	"fmt"
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

// helper to determine what type an extra subdirectory is based on its directory name or file type
func DetermineExtraItemType(dir, file string) config.ExtraType {
	folder := strings.ToLower(dir)
	// check if the directory name matches any known extra types
	switch folder {
	case "art", "artwork", "images", "pictures":
		return config.ExtraTypeArt
	case "ost", "soundtrack", "music", "sounds", "audio":
		return config.ExtraTypeOST
	case "manual", "manuals":
		return config.ExtraTypeManual
	case "video", "videos", "trailers":
		return config.ExtraTypeVideo
	default:
		// if not, check the file extension to determine the type
		ext := strings.ToLower(filepath.Ext(file))
		switch ext {
		case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff":
			return config.ExtraTypeArt
		case ".mp3", ".flac", ".wav", ".ogg":
			return config.ExtraTypeOST
		case ".pdf", ".doc", ".docx", ".txt":
			return config.ExtraTypeManual
		case ".mp4", ".avi", ".mkv":
			return config.ExtraTypeVideo
		default:
			return config.ExtraTypeUnknown
		}
	}
}

func DiscoverExtraItems(extraDir string) ([]config.ExtraItem, error) {
	var extraItems []config.ExtraItem
	err := filepath.Walk(extraDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("filepath walk: %w", err)
		}
		if !info.IsDir() {
			dir := filepath.Base(filepath.Dir(path))
			extraType := DetermineExtraItemType(dir, info.Name())
			if extraType != config.ExtraTypeUnknown {
				extraItems = append(extraItems, config.ExtraItem{
					Name:      info.Name(),
					ExtraType: extraType,
					FilePath:  path,
				})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return extraItems, nil
}
