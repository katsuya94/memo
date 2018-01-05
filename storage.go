package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

type storage interface {
	retrive(date) (string, error)
	store(date, string) error
}

type localStorage struct{}

func (*localStorage) retrive(d date) (string, error) {
	return "", errNotFound{}
}

func (*localStorage) store(d date, contents string) error {
	dir, err := memoDir()
	if err != nil {
		return err
	}

	path := path.Join(dir, d.String())
	return ioutil.WriteFile(path, []byte(contents), 0644)
}

func memoDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	dir := path.Join(usr.HomeDir, ".memo")

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	return dir, nil
}
