package hmf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	ipfs "github.com/ipfs/go-ipfs-api"

	"github.com/miguelcnf/hmf/pkg/security"
)

func StoreSecure(file string) (CID string, passphrase []byte, err error) {
	// generate passphrase
	passphrase, err = security.GeneratePassphrase()
	if err != nil {
		return "", nil, fmt.Errorf("error generating passphrase: %v", err)
	}

	CID, err = Store(file, passphrase)

	return
}

func Store(file string, passphrase []byte) (CID string, err error) {
	// read file data
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	// encrypt data
	encrypted, err := security.Encrypt(data, passphrase)
	if err != nil {
		return "", fmt.Errorf("error encrypting data: %v", err)
	}

	// connect to running local node
	sh := ipfs.NewShell("localhost:5001")

	// add encrypted file
	CID, err = sh.Add(bytes.NewBuffer(encrypted))
	if err != nil {
		return "", fmt.Errorf("error adding file to ipfs: %v", err)
	}

	return CID, nil
}

func Read(encryptedFile string, passphrase []byte) error {
	// read encrypted file data
	encrypted, err := ioutil.ReadFile(encryptedFile)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	data, err := security.Decrypt(encrypted, passphrase)
	if err != nil {
		return fmt.Errorf("error decrypting data: %v", err)
	}

	// write file in-place
	err = ioutil.WriteFile(encryptedFile, data, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

func Get(CID, dst string, passphrase []byte) error {
	// connect to running local node
	sh := ipfs.NewShell("localhost:5001")

	// get encrypted file and write it to fs destination
	err := sh.Get(CID, dst)
	if err != nil {
		return fmt.Errorf("error getting file from ipfs: %v", err)
	}

	return Read(dst, passphrase)
}

func Hold(CID string) error {
	// connect to running local node
	sh := ipfs.NewShell("localhost:5001")

	// get file and write it to a temp fs destination
	err := sh.Get(CID, CID)
	if err != nil {
		return fmt.Errorf("error getting file from ipfs: %v", err)
	}

	// pin file
	err = sh.Pin(CID)
	if err != nil {
		return fmt.Errorf("error pinning file on ipfs: %v", err)
	}

	return nil
}
