package gstrings

import (
	"errors"
	"regexp"
)

func Search(s, regex string) (int, error) {
	v, err := regexp.Compile(regex)

	if err != nil {
		return -1, err
	}

	result := v.FindStringIndex(s)

	if len(result) >= 1 {
		return result[0], nil
	}

	return -1, errors.New("not found")
}
