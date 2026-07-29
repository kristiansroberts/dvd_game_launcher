package scan

import (
	"os"
	"path/filepath"
	"strings"
)

func FindEXE(dir string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, f := range files {
		if strings.HasSuffix(strings.ToLower(f.Name()), strings.ToLower(".exe")) {
			return filepath.Join(dir, f.Name()), nil
		}
	}
	return "", nil
}
