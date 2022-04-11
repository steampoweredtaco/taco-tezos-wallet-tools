package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"taco-exports/isdelve"

	"golang.org/x/term"
)

func getPassword() (string, error) {
	fmt.Print("Enter Password: ")
	var bytePassword []byte
	if isdelve.Enabled {
		reader := bufio.NewReader(os.Stdin)
		bytePassword, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		return bytePassword, nil
	}
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}

	password := string(bytePassword)
	return password, nil
}
