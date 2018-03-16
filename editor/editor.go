package editor

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
)

type Editor struct{}

func (e Editor) Read(s string) error {
	path := temporaryFilePath()

	err := ioutil.WriteFile(path, []byte(s), 0444)
	if err != nil {
		return err
	}

	return e.run(path)
}

func (e Editor) Edit(s string) (string, error) {
	path := temporaryFilePath()

	err := ioutil.WriteFile(path, []byte(s), 0644)
	if err != nil {
		return s, err
	}

	err = e.run(path)
	if err != nil {
		return s, err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return s, err
	}

	return string(b), nil
}

func (e Editor) run(path string) error {
	editorCommand := os.Getenv("VISUAL")
	if editorCommand == "" {
		editorCommand = os.Getenv("EDITOR")
	}
	if editorCommand == "" {
		return fmt.Errorf("open: no editor specified in $VISUAL, $EDITOR")
	}

	cmd := exec.Command(editorCommand, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

var hexRunes = []rune("0123456789abcdef")

func temporaryFilePath() string {
	runes := make([]rune, 8)
	for i := range runes {
		runes[i] = hexRunes[rand.Int63()%16]
	}
	basename := fmt.Sprintf("%v-%v", "memo", string(runes))
	return path.Join(os.TempDir(), basename)
}
