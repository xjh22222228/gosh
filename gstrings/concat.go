// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.1

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
