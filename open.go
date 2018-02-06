package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

func open(d, today date) error {
	if time.Time(d).After(time.Time(today)) {
		return fmt.Errorf("open: can't open future memo")
	}

	storage := storage(&localStorage{})

	contents, err := storage.retrieve(d)

	if time.Time(d).Before(time.Time(today)) {
		if err != nil {
			return err
		}

		return read(contents)
	}

	if _, ok := err.(errNotFound); ok {
		var (
			mostRecentDate     date
			mostRecentContents string
			sections           []section
		)

		mostRecentDate, err = storage.mostRecent()
		if err != nil {
			return err
		}

		if mostRecentDate == zeroDate { // if the date is empty
			mostRecentContents = ""
		} else {
			mostRecentContents, err = storage.retrieve(mostRecentDate)
			if err != nil {
				return err
			}
		}

		sections, err = parse(mostRecentContents)
		if err != nil {
			return err
		}

		namedSections := []section{}
		for _, section := range sections {
			if section.info.name != "" {
				namedSections = append(namedSections, section)
			}
		}

		dumpedSections := make([]string, len(namedSections))
		for i, s := range namedSections {
			dumpedSections[i] = dump(s)
		}

		if len(dumpedSections) == 0 {
			contents = "---\n"
		} else {
			contents = strings.Join(dumpedSections, "")
		}
	}
	if err != nil {
		return err
	}

	contents, err = edit(contents)
	if err != nil {
		return err
	}

	_, err = parse(contents)

	for err != nil {
		_, ok := err.(errMalformedSectionHeader)
		if !ok {
			break
		}

		fmt.Println(err.Error())

		var discard bool
		discard, err = confirm(false, "Discard changes?")
		if err != nil {
			return err
		}
		if discard {
			return nil
		}

		contents, err = edit(contents)
		if err != nil {
			return err
		}

		_, err = parse(contents)
	}
	if err != nil {
		return err
	}

	return storage.store(d, contents)
}

func read(contents string) error {
	filename := tempFilename("memo")

	err := ioutil.WriteFile(filename, []byte(contents), 0444)
	if err != nil {
		return err
	}

	return editor(filename)
}

func edit(contents string) (string, error) {
	filename := tempFilename("memo")

	err := ioutil.WriteFile(filename, []byte(contents), 0644)
	if err != nil {
		return contents, err
	}

	err = editor(filename)
	if err != nil {
		return contents, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return contents, err
	}

	return string(bytes), nil
}

var hex = []rune("0123456789abcdef")

func tempFilename(base string) string {
	runes := make([]rune, 8)
	for i := range runes {
		runes[i] = hex[rand.Int63()%16]
	}
	basename := fmt.Sprintf("%v-%v", base, string(runes))
	return path.Join(os.TempDir(), basename)
}

func editor(filename string) error {
	editorCommand := os.Getenv("VISUAL")
	if editorCommand == "" {
		editorCommand = os.Getenv("EDITOR")
	}
	if editorCommand == "" {
		return fmt.Errorf("open: no editor specified in $VISUAL, $EDITOR")
	}

	cmd := exec.Command(editorCommand, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
