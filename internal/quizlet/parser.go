package quizlet

import (
	"fmt"
	"regexp"
	"strings"
)

// Parse text which is a part of lesson review memo and return Word object.
func Parse(text string) (*Word, bool, error) {
	// Skip
	if strings.HasPrefix(text, "★") || strings.HasPrefix(text, "・") || text == "" {
		return nil, true, nil
	}

	// Useful Words and Phrases
	re := regexp.MustCompile(`([\w\s]+)\((\w+)\)\s*(\/.*\/)?\s*(.+)\/(.+)e\.g\.(.+)`)
	if re.MatchString(text) {
		match := re.FindStringSubmatch(text)
		return &Word{
			English:     strings.TrimSpace(match[1]),
			Kind:        strings.TrimSpace(match[2]),
			Japanese:    strings.TrimSpace(match[4]),
			Description: strings.TrimSpace(match[5]),
			Example:     strings.TrimSpace(match[6]),
		}, false, nil
	}

	// Useful Expressions
	re = regexp.MustCompile(`※(.+):(.*)`)
	if re.MatchString(text) {
		match := re.FindStringSubmatch(text)
		return &Word{
			English:     strings.TrimSpace(match[1]),
			Kind:        "expression",
			Japanese:    strings.TrimSpace(match[2]),
			Description: "",
			Example:     "",
		}, false, nil
	}

	return nil, false, fmt.Errorf("fail to parse '%s'", text)
}
