package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/miguelcnf/hmf/internal/hmf"
	"github.com/miguelcnf/hmf/internal/input"
)

func main() {
	// Not a proper cli but meh... :)
	var (
		share = flag.Bool("share", false, "share file (incompatible with other actions)")
		read  = flag.Bool("read", false, "read file (incompatible with other actions)")
		get   = flag.Bool("get", false, "get file (incompatible with other actions)")
		hold  = flag.Bool("hold", false, "hold file (incompatible with other actions)")

		file     = flag.String("file", "", "path to file")
		password = flag.Bool("password", false, "enable password mode (default is false)")
	)
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	go func() {
		<-sigs
		os.Exit(128)
	}()

	switch {
	case *share:
		if *file == "" {
			flag.Usage()
			os.Exit(2)
		}

		var (
			link       string
			passphrase []byte
			holdErr    error
		)
		if *password {
			passphrase, err := input.SetPassphrase()
			if err != nil {
				fmt.Printf("Unable to set passphrase: %v\n", err)
				os.Exit(1)
			}

			link, holdErr = hmf.Store(*file, passphrase)
		} else {
			link, passphrase, holdErr = hmf.StoreSecure(*file)
			if holdErr == nil {
				fmt.Printf("Make sure to write down your 10 word generated passphrase: %s\n", passphrase)
			}
		}
		if holdErr != nil {
			fmt.Printf("Failed to hold file: %v\n", holdErr)
			os.Exit(1)
		}

		fmt.Printf("Here's your file identifier: %s\n", link)
	case *read:
		if *file == "" {
			flag.Usage()
			os.Exit(2)
		}

		passphrase, err := input.ReadPassphrase()
		if err != nil {
			fmt.Printf("Unable to read passphrase: %v\n", err)
			os.Exit(1)
		}

		err = hmf.Read(*file, passphrase)
		if err != nil {
			fmt.Printf("Failed to read file: %v\n", err)
			os.Exit(1)
		}
	case *get:
		if *file == "" {
			flag.Usage()
			os.Exit(2)
		}

		CID, err := input.ReadCID()
		if err != nil {
			fmt.Printf("Unable to read file identifier: %v\n", err)
			os.Exit(1)
		}

		passphrase, err := input.ReadPassphrase()
		if err != nil {
			fmt.Printf("Unable to read passphrase: %v\n", err)
			os.Exit(1)
		}

		err = hmf.Get(CID, *file, passphrase)
		if err != nil {
			fmt.Printf("Failed to get file: %v\n", err)
			os.Exit(1)
		}
	case *hold:
		CID, err := input.ReadCID()
		if err != nil {
			fmt.Printf("Unable to read file identifier: %v\n", err)
			os.Exit(1)
		}

		err = hmf.Hold(CID)
		if err != nil {
			fmt.Printf("Failed to hold file: %v\n", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(2)
	}

	fmt.Println("Success!")
}
