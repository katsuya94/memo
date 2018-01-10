package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

type storage interface {
	retrieve(date) (string, error)
	store(date, string) error
	mostRecent() (date, error)
}

type localStorage struct{}

func (*localStorage) retrieve(d date) (string, error) {
	dir, err := memoDir()
	if err != nil {
		return "", err
	}

	path := path.Join(dir, d.String())
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errNotFound{}
	}

	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}

func (*localStorage) store(d date, contents string) error {
	dir, err := memoDir()
	if err != nil {
		return err
	}

	path := path.Join(dir, d.String())
	return ioutil.WriteFile(path, []byte(contents), 0644)
}

func (*localStorage) mostRecent() (date, error) {
	dir, err := memoDir()
	if err != nil {
		return date{}, err
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return date{}, err
	}

	mostRecentFilename := fileInfos[len(fileInfos)-1].Name()
	d, err := parseDate(mostRecentFilename)
	if err != nil {
		return date{}, err
	}

	return d, nil
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
