package input

import (
	"io"
	"os"
)

var (
	screen = struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
)
