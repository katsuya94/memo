package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func open(d date) error {
	storage := storage(&localStorage{})

	contents, err := storage.retrive(d)

	if _, ok := err.(errNotFound); ok {
		contents, err = blank(d)
	}

	if err != nil {
		return err
	}

	contents, err = edit(contents)
	if err != nil {
		return err
	}

	fmt.Println(parse(contents))

	return storage.store(d, contents)
}

var editorEnvVars = []string{"VISUAL", "EDITOR"}

func edit(contents string) (string, error) {
	f, err := ioutil.TempFile("", "memo")
	if err != nil {
		return contents, err
	}

	filename := f.Name()
	f.Close()

	err = ioutil.WriteFile(filename, []byte(contents), 0)
	if err != nil {
		return contents, err
	}

	editorCommand := ""

	for _, envVar := range editorEnvVars {
		if editorCommand != "" {
			break
		}
		editorCommand = os.Getenv(envVar)
	}

	if editorCommand == "" {
		return contents, fmt.Errorf(
			"open: no editor specified in $%v",
			strings.Join(editorEnvVars, ", $"),
		)
	}

	cmd := exec.Command(editorCommand, f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return contents, err
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return contents, err
	}

	return string(b), nil
}

func blank(d date) (string, error) {
	return "placeholder", nil
}
