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
	case strLinux, "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	case strWindows:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	}

	return err
}
