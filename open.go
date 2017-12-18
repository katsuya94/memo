package main

import (
	"errors"
	"fmt"
)

var (
	errorNotFound = errors.New("Memo not found")
)

func open(d date) error {
	fmt.Println(d)

	storage := storage(&localStorage{})

	raw, err := storage.retrive(d)

	if err == nil {
		fmt.Println(raw)
	} else if err == errorNotFound {
		fmt.Println("")
	} else {
		return err
	}

	return nil
}

type storage interface {
	retrive(date) (string, error)
	store(date, string) error
}

type localStorage struct{}

func (*localStorage) retrive(d date) (string, error) {
	return "", errorNotFound
}

func (*localStorage) store(d date, raw string) error {
	return nil
}
