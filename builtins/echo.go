package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	fmt.Fprintln(w, strings.Join(args, " "))
	return nil
}
