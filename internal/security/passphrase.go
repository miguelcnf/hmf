package security

import (
	"crypto/sha256"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
	"golang.org/x/crypto/pbkdf2"
)

func GeneratePassphrase() (passphrase []byte, err error) {
	list, err := diceware.Generate(10)
	if err != nil {
		return nil, err
	}

	words := strings.Join(list, " ")
	return []byte(words), nil
}

var (
	salt = []byte("0148848b-3950-46e8-a031-188f4f732eb1")
)

func keyStretching(passphrase []byte) (key []byte) {
	return pbkdf2.Key(passphrase, salt, 1000, 32, sha256.New)
}
