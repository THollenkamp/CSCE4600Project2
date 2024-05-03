package builtins

import (
	"fmt"
	"io"
)

func Help(w io.Writer, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: 'help <cmd>'")
	}

	switch args[0] {
	case "alias":
		fmt.Fprintf(w, "alias: create an alias for a command\n")
	case "cd":
		fmt.Fprintf(w, "cd: change working directory of the shell\n")
	case "echo":
		fmt.Fprintf(w, "echo: print text into the console\n")
	case "env":
		fmt.Fprintf(w, "env: prints all environment variables to console\n")
	case "type":
		fmt.Fprintf(w, "type: get file type of an input file\n")
	case "times":
		fmt.Fprintf(w, "times: get time since process was started, and time since command was last ran\n")
	default:
		fmt.Fprintf(w, "help: print help message for a command\n")
		fmt.Fprintf(w, "usage: 'help <cmd>'\n")
	}

	return nil
}
