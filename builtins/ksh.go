package builtins

import (
	"fmt"
	"os"
	"os/exec"
)

func Ksh() error {
	// Execute the ksh command.
	cmd := exec.Command("ksh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ksh: %v", err)
	}

	return nil
}
