package builtins

import (
	"fmt"
	"os"
	"os/exec"
)

func Shell(args ...string) error {
	// Check if there are any arguments provided to the sh command.
	if len(args) == 0 {
		return fmt.Errorf("sh: missing argument")
	}

	// Execute the sh command with the provided arguments.
	cmd := exec.Command("sh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("sh: %v", err)
	}

	return nil
}
