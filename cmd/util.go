package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func confirm(recommended bool, message string) (bool, error) {
	if recommended {
		fmt.Printf("%v (Y/n) ", message)
	} else {
		fmt.Printf("%v (y/N) ", message)
	}

	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return recommended, err
	}

	b = bytes.TrimSpace(b)
	b = bytes.ToLower(b)

	switch string(b) {
	case "":
		return recommended, nil
	case "n":
		return false, nil
	case "y":
		return true, nil
	default:
		return confirm(recommended, "Please type y or n.")
	}
}
