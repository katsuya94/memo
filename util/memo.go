package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	sectionHeaderRegexp            = regexp.MustCompile(`^ {0,3}(\*{3}|-{3}|_{3})`)
	sectionHeaderInformationRegexp = regexp.MustCompile(`^ {0,3}(\*{3}|-{3}|_{3}) *([0-9a-zA-Z_-]+|(\+[0-9a-zA-Z_-]+ *)*) *\n$`)
)

type Section struct {
	Tags []string
	Name string
	Body []byte
}

type Memo []Section

func (m Memo) HasNamedSection(name string) bool {
	for _, s := range m {
		if s.Name == name {
			return true
		}
	}
	return false
}

type MemoFormatError struct {
	line    int
	message string
}

func (e MemoFormatError) Error() string {
	return fmt.Sprintf("line %v: %v", e.line, e.message)
}

func ReadMemo(rd io.Reader) (Memo, error) {
	var (
		r = bufio.NewReader(rd)
		i = 0
		m = Memo{}
	)

	for {
		i += 1
		line, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		section, ok, err := parseLine(line)
		if err != nil {
			return nil, MemoFormatError{i, err.Error()}
		}

		if ok {
			m = append(m, section)
		} else if len(m) == 0 {
			return nil, MemoFormatError{i, "found line before the first section header"}
		} else {
			m[len(m)-1].Body = append(m[len(m)-1].Body, line...)
		}
	}

	return m, nil
}

func parseLine(line []byte) (Section, bool, error) {
	if !sectionHeaderRegexp.Match(line) {
		return Section{}, false, nil
	}

	match := sectionHeaderInformationRegexp.FindSubmatch(line)
	if match == nil {
		return Section{}, true, fmt.Errorf("found malformed section header")
	}

	tagsOrName := match[2]

	if len(tagsOrName) == 0 {
		return Section{}, true, nil
	} else if tagsOrName[0] == '+' {
		s := Section{}
		for _, token := range bytes.Split(match[2], []byte{' '}) {
			if len(token) > 0 {
				s.Tags = append(s.Tags, string(token[1:]))
			}
		}
		return s, true, nil
	} else {
		return Section{Name: string(match[2])}, true, nil
	}
}

func WriteMemo(m Memo, wr io.Writer) error {
	var (
		w = bufio.NewWriter(wr)
	)

	for _, s := range m {
		fmt.Fprint(w, "---")

		if s.Name != "" {
			fmt.Fprintf(w, " %v", s.Name)
		}

		if len(s.Tags) > 0 {
			fmt.Fprintf(w, " +%v", strings.Join(s.Tags, " +"))
		}

		fmt.Fprint(w, "\n")

		if _, err := w.Write(s.Body); err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}
