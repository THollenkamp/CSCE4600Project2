package builtins

import (
	"fmt"
	"io"
	"time"
)

func Times(w io.Writer, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("usage: times takes no arguments")
	}

	// Get the user and system time
	userTime := time.Since(startTime) - systemTime
	systemTime = time.Since(startTime)

	// Print the times
	fmt.Fprintf(w, "user\t%s\n", userTime)
	fmt.Fprintf(w, "sys\t%s\n", systemTime)

	return nil
}

var startTime = time.Now()
var systemTime time.Duration
