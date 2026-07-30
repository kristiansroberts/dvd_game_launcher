package scan

import (
	"os"
	"path/filepath"
	"testing"
)

// simple testing to ensure the scan functions properly find the exes in each folder
func TestFindEXEs_ReturnsFirstEXE(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	game1exe := filepath.Join(tempDir, "game1.exe")
	game1settingsexe := filepath.Join(tempDir, "game1settings.exe")

	if err := os.WriteFile(game1exe, []byte("dummy content"), 0644); err != nil {
		t.Fatalf("Failed to create game1.exe: %v", err)
	}
	if err := os.WriteFile(game1settingsexe, []byte("dummy content"), 0644); err != nil {
		t.Fatalf("Failed to create game1settings.exe: %v", err)
	}

	got, err := findEXEs(tempDir)
	if err != nil {
		t.Fatalf("findEXEs returned an error: %v", err)
	}
	if got != game1exe {
		t.Errorf("findEXEs returned %q, want %q", got, game1exe)
	}
}

func TestFindEXEs_ReturnsEmptyStringWhenNoEXE(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	if err := os.WriteFile(filepath.Join(tempDir, "file.txt"), []byte("dummy content"), 0644); err != nil {
		t.Fatalf("Failed to create file.txt: %v", err)
	}

	if err := os.WriteFile(filepath.Join(tempDir, "image.png"), []byte("dummy content"), 0644); err != nil {
		t.Fatalf("Failed to create image.png: %v", err)
	}

	got, err := findEXEs(tempDir)
	if err != nil {
		t.Fatalf("findEXEs returned an error: %v", err)
	}
	if got != "" {
		t.Errorf("findEXEs returned %q, want empty string", got)
	}
}
