package clear

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

var ErrUnsupportedPlatform = errors.New("unsupported platform")

// Clear clears console.
func Clear() error {
	commands := map[string]string{
		"linux":   "clear",
		"windows": "cls",
	}

	commandStr, isSupported := commands[runtime.GOOS]
	if !isSupported {
		return ErrUnsupportedPlatform
	}

	cmd := exec.Command(commandStr)
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	return err
}
