package modscan

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const civ6AppID = "289070"

func steamModFolderLocation() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine user home directory: %w", err)
	}

	var path string
	switch runtime.GOOS {
	case "windows":
		progFiles := os.Getenv("ProgramFiles(x86)")
		if progFiles == "" {
			progFiles = os.Getenv("ProgramFiles")
		}
		if progFiles != "" {
			path = filepath.Join(progFiles, "Steam", "steamapps", "workshop", "content", civ6AppID)
		} else {
			path = filepath.Join(homeDir, "Steam", "steamapps", "workshop", "content", civ6AppID)
		}
	case "darwin":
		path = filepath.Join(homeDir, "Library", "Application Support", "Steam", "steamapps", "workshop", "content", civ6AppID)
	case "linux":
		if xdgData := os.Getenv("XDG_DATA_HOME"); xdgData != "" {
			path = filepath.Join(xdgData, "Steam", "steamapps", "workshop", "content", civ6AppID)
		} else {
			path = filepath.Join(homeDir, ".local", "share", "Steam", "steamapps", "workshop", "content", civ6AppID)
		}
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return path, nil
}

func locationModFolderLocation() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine user home directory: %w", err)
	}

	var path string
	switch runtime.GOOS {
	case "windows":
		path = filepath.Join(homeDir, "Documents", "My Games", "Sid Meier's Civilization VI", "Mods")
	case "darwin":
		path = filepath.Join(homeDir, "Library", "Application Support", "Sid Meier's Civilization VI", "Mods")
	case "linux":
		path = filepath.Join(homeDir, ".local", "share", "aspyr-media", "Sid Meier's Civilization VI", "Mods")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return path, nil
}
