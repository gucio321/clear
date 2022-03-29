package clear

import (
	"os"
	"os/exec"
	"runtime"
)

// Clear clears console.
func Clear() error {
	var err error

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	}

	return err
}
