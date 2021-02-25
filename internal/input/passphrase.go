package input

import (
	"errors"
	"fmt"

	"golang.org/x/term"
)

func SetPassphrase() (passphrase []byte, err error) {
	state, err := term.MakeRaw(0)
	if err != nil {
		return
	}
	defer term.Restore(0, state)

	terminal := term.NewTerminal(screen, "")

	var password, confirm string
	fmt.Print("Choose a passphrase: ")
	password, err = terminal.ReadPassword("")
	if err != nil {
		return
	}

	fmt.Print("Confirm your passphrase: ")
	confirm, err = terminal.ReadPassword("")
	if err != nil {
		return
	}

	if password != confirm {
		err = errors.New("passphrases don't match")
	}

	return []byte(password), err
}

func ReadPassphrase() (passphrase []byte, err error) {
	state, err := term.MakeRaw(0)
	if err != nil {
		return
	}
	defer term.Restore(0, state)

	terminal := term.NewTerminal(screen, "")

	var password string
	fmt.Print("Passphrase: ")
	password, err = terminal.ReadPassword("")

	return []byte(password), err
}
