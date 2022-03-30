// Package clear contains method used to clear the shell console.
package clear

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ErrUnsupportedPlatform is returned, when cleaning console on
// current platform is not supported.
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

	// nolint:gosec // it is intendet
	cmd := exec.Command(commandStr)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error clearing console: %w", err)
	}

	return nil
}
