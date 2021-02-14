package input

import (
	"fmt"

	"golang.org/x/term"
)

func ReadCID() (CID string, err error) {
	state, err := term.MakeRaw(0)
	if err != nil {
		return
	}
	defer term.Restore(0, state)

	terminal := term.NewTerminal(screen, "")

	fmt.Print("File identifier: ")
	CID, err = terminal.ReadLine()

	return
}
