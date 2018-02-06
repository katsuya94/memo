package main

import (
	"fmt"
	"regexp"
	"strings"
)

type sectionInfo struct {
	tags []string
	name string
}

type section struct {
	info sectionInfo
	body string
}

func parse(contents string) ([]section, error) {
	lines := strings.Split(contents, "\n")
	sections := []section{}

	var currentInfo, nextInfo sectionInfo
	var i, j int

	currentInfo = sectionInfo{}

	for true {
		for ; j < len(lines); j++ {
			var ok bool
			var err error
			nextInfo, ok, err = sectionHeader(lines[j])
			if err != nil {
				return []section{}, err
			}
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

	if sections[0].body == "" {
		return sections[1:], nil
	}

	return sections, nil
}

var (
	sectionHeaderRegexp     = regexp.MustCompile(`^ {0,3}(\*{3}|-{3}|_{3})`)
	sectionHeaderInfoRegexp = regexp.MustCompile(`^ {0,3}(\*{3}|-{3}|_{3}) *([0-9a-zA-Z_-]+|(\+[0-9a-zA-Z_-]+ *)*) *$`)
)

func sectionHeader(line string) (sectionInfo, bool, error) {
	var info sectionInfo

	if sectionHeaderRegexp.MatchString(line) {
		if ok := sectionHeaderInfoRegexp.MatchString(line); !ok {
			return info, true, errMalformedSectionHeader{line}
		}

		match := sectionHeaderInfoRegexp.FindStringSubmatch(line)

		if match[2] == "" {
			return info, true, nil
		}

		if match[2][0] == '+' {
			tokens := strings.Split(match[2], " ")
			for _, token := range tokens {
				if token != "" {
					info.tags = append(info.tags, token[1:])
				}
			}
		} else {
			info.name = match[2]
		}

		return info, true, nil
	}

	return info, false, nil
}

func dump(section section) string {
	if section.info.name != "" {
		return fmt.Sprintf("--- %v\n%v\n", section.info.name, section.body)
	} else if len(section.info.tags) > 0 {
		tagsWithoutLeadingPlus := strings.Join(section.info.tags, " +")
		return fmt.Sprintf("--- +%v\n%v\n", tagsWithoutLeadingPlus, section.body)
	} else {
		return fmt.Sprintf("---\n%v\n", section.body)
	}
}
