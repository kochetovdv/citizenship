package regulars

import (
	"fmt"
	"regexp"
	"strings"
)

// FindString finds a string in another string
func FindString(regularExpression string, body string) (string, error) {
	respRegExp := regexp.MustCompile(regularExpression)
	resp := respRegExp.FindString(body)
	if resp == "" {
		return "", fmt.Errorf("no date found in %s", body)
	}
	resp = strings.TrimSpace(resp)
	return resp, nil
}

// FindAllStrings2D returns 2 regulars in a 2D slice of strings
func FindAllStrings2D(regularExpression string, body string) ([][]string, error) {
	respRegExp := regexp.MustCompile(regularExpression)
	resp := respRegExp.FindAllStringSubmatch(body, -1)
	if len(resp) < 1 {
		return nil, fmt.Errorf("no date found in %s", body)
	}

	return resp, nil
}

// FindAllStrings finds all the matches of a regular expression in a string
func FindAllStrings(regularExpression string, body string) ([]string, error) {
	respRegExp := regexp.MustCompile(regularExpression)
	matches := respRegExp.FindAllStringSubmatch(body, -1)

	if len(matches) < 1 {
		return nil, fmt.Errorf("no matches found in %s", body)
	}

	var result []string
	for _, match := range matches {
		result = append(result, match[0])
	}

	return result, nil
}
