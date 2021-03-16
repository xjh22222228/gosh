// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.6

package gfunctools

import (
	"fmt"
	"github.com/xjh22222228/gosh/gstr"
	"strings"
)

func Currency(num float64, symbol ...string) string {
	currencySymbol := "$"

	if len(symbol) > 0 {
		currencySymbol = symbol[0]
	}

	toFixed := fmt.Sprintf("%.2f", num)
	floatN := gstr.Slice(toFixed, -3)
	numPrefix := gstr.Slice(toFixed, 0, -3)
	length := len(numPrefix)
	var str strings.Builder

	for i := 1; i <= length; i++ {
		str.WriteString(string(numPrefix[length - i]))
		if i > 2 && i % 3 == 0 && i != length {
			str.WriteString(",")
		}
	}

	str.WriteString(currencySymbol)

	return gstr.Reverse(str.String()) + floatN
}
