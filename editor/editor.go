package editor

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path"
)

type Editor struct {
	path string
	file *os.File
}

func NewEditor() (Editor, error) {
	var (
		e   = Editor{}
		err error
	)
	e.path = temporaryFilePath()

	e.file, err = os.OpenFile(e.path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return e, err
	}

	return e, nil
}

func (e Editor) Write(p []byte) (int, error) {
	return e.file.Write(p)
}

func (e Editor) Read(p []byte) (int, error) {
	return e.file.Read(p)
}

func (e Editor) Close() error {
	return e.file.Close()
}

func (e *Editor) Launch(readonly bool) error {
	var (
		err error
	)

	if readonly {
		if err = e.file.Chmod(0400); err != nil {
			return err
		}
	}

	editorCommand := os.Getenv("VISUAL")
	if editorCommand == "" {
		editorCommand = os.Getenv("EDITOR")
	}
	if editorCommand == "" {
		return fmt.Errorf("open: no editor specified in $VISUAL, $EDITOR")
	}

	cmd := exec.Command(editorCommand, e.path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		return err
	}

	e.file, err = os.Open(e.path)
	if err != nil {
		return err
	}

	return nil
}

var hexBytes = []byte("0123456789abcdef")

func temporaryFilePath() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = hexBytes[rand.Int63()%16]
	}
	basename := fmt.Sprintf("%v-%v", "memo", string(b))
	return path.Join(os.TempDir(), basename)
}
