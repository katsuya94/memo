package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	errNotFound = errors.New("Memo not found")
)

func open(d date) error {
	fmt.Println(d)

	storage := storage(&localStorage{})

	raw, err := storage.retrive(d)

	fmt.Println(raw)

	if err == nil {
		return edit()
	} else if err == errNotFound {
		return edit()
	} else {
		return err
	}
}

type storage interface {
	retrive(date) (string, error)
	store(date, string) error
}

type localStorage struct{}

func (*localStorage) retrive(d date) (string, error) {
	return "", errNotFound
}

func (*localStorage) store(d date, raw string) error {
	return nil
}

func edit() error {
	f, err := ioutil.TempFile("", "memo")
	if err != nil {
		return err
	}

	cmd := exec.Command("vim", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
