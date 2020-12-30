package gstrings

import (
	"strings"
)

func Concat(s ...string) string  {
	var builder strings.Builder

	for _, v := range s {
		builder.WriteString(v)
	}

	return builder.String()
}
