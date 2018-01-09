package main

import (
	"regexp"
	"strings"
)

type sectionInfo struct {
	tags       []string
	name       string
	persistent bool
}

type section struct {
	info sectionInfo
	body string
}

func parse(contents string) []section {
	lines := strings.Split(contents, "\n")
	sections := []section{}

	var currentInfo, nextInfo sectionInfo
	var i, j int

	currentInfo = sectionInfo{}

	for true {
		for ; j < len(lines); j++ {
			var ok bool
			nextInfo, ok = sectionHeader(lines[j])
			if ok {
				break
			}
		}

		section := section{
			info: currentInfo,
			body: strings.Join(lines[i:j], "\n"),
		}
		sections = append(sections, section)

		if j >= len(lines) {
			break
		}

		currentInfo = nextInfo
		j++
		i = j
	}

	return sections
}

var sectionHeaderRegexp = regexp.MustCompile(`^ {0,3}(\*{3}|-{3}|_{3})`)

func sectionHeader(line string) (sectionInfo, bool) {
	if sectionHeaderRegexp.MatchString(line) {
		return sectionInfo{}, true
	} else {
		return sectionInfo{}, false
	}
}
