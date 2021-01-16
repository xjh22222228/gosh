// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.6

package grandom

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	C_RGB = iota
	C_RGBA = iota
	C_HEX = iota
	C_COLOR = iota
)

var value = [139]string{
	"LightPink", "Pink", "Crimson",
	"LavenderBlush", "PaleVioletRed",
	"HotPink", "DeepPink", "MediumVioletRed",
	"Orchid", "Thistle", "Plum",
	"Violet", "Magenta", "Fuchsia", "DarkMagenta",
	"Purple", "MediumOrchid", "DarkViolet",
	"DarkOrchid", "BlueViolet", "MediumPurple",
	"MediumSlateBlue", "SlateBlue", "DarkSlateBlue",
	"Lavender", "GhostWhite", "Blue",
	"MediumBlue", "MidnightBlue", "DarkBlue",
	"Navy", "RoyalBlue", "CornflowerBlue",
	"LightSteelBlue", "LightSlateGray", "SlateGray",
	"DodgerBlue", "AliceBlue", "SteelBlue",
	"LightSkyBlue", "SkyBlue", "DeepSkyBlue",
	"LightBlue", "PowderBlue", "CadetBlue",
	"Azure", "LightCyan", "PaleTurquoise",
	"Cyan", "Aqua", "DarkTurquoise",
	"DarkSlateGray", "DarkCyan", "Teal",
	"MediumTurquoise", "LightSeaGreen", "Turquoise",
	"Aquamarine", "MediumAquamarine", "MediumSpringGreen",
	"MintCream", "SpringGreen", "MediumSeaGreen",
	"SeaGreen", "Honeydew", "LightGreen",
	"PaleGreen", "DarkSeaGreen", "LimeGreen",
	"Lime", "ForestGreen", "Green",
	"DarkGreen", "Chartreuse", "LawnGreen",
	"GreenYellow", "DarkOliveGreen", "YellowGreen",
	"OliveDrab", "Beige", "LightGoldenrodYellow",
	"Ivory", "LightYellow", "Yellow",
	"Olive", "DarkKhaki", "LemonChiffon",
	"PaleGoldenrod", "Khaki", "Gold",
	"Cornsilk", "Goldenrod", "DarkGoldenrod",
	"FloralWhite", "OldLace", "Wheat",
	"Moccasin", "Orange", "PapayaWhip",
	"BlanchedAlmond", "NavajoWhite", "AntiqueWhite",
	"Tan", "BurlyWood", "Bisque",
	"DarkOrange", "Linen", "Peru",
	"PeachPuff", "SandyBrown", "Chocolate",
	"SaddleBrown", "Seashell", "Sienna",
	"LightSalmon", "Coral", "OrangeRed",
	"DarkSalmon", "Tomato", "MistyRose",
	"Salmon", "Snow", "LightCoral",
	"RosyBrown", "IndianRed", "Red",
	"Brown", "FireBrick", "DarkRed",
	"Maroon", "White", "WhiteSmoke",
	"Gainsboro", "LightGrey", "Silver",
	"DarkGray", "Gray", "DimGray",
	"Black",
}

func Color(mode int) string {
	var sCol strings.Builder

	switch mode {
		case C_COLOR:
			r := RandInt(0, len(value))
			sCol.WriteString(value[r])
			break
		case C_RGB:
			r := RandInt(0, 255)
			g := RandInt(0, 255)
			b := RandInt(0, 255)

			sCol.WriteString("rgb(")
			sCol.WriteString(strconv.Itoa(r))
			sCol.WriteString(",")
			sCol.WriteString(strconv.Itoa(g))
			sCol.WriteString(",")
			sCol.WriteString(strconv.Itoa(b))
			sCol.WriteString(")")
			break
		case C_RGBA:
			r := RandInt(0, 255)
			g := RandInt(0, 255)
			b := RandInt(0, 255)
			a := fmt.Sprintf("%.2f", Random())

			sCol.WriteString("rgba(")
			sCol.WriteString(strconv.Itoa(r))
			sCol.WriteString(",")
			sCol.WriteString(strconv.Itoa(g))
			sCol.WriteString(",")
			sCol.WriteString(strconv.Itoa(b))
			sCol.WriteString(",")
			sCol.WriteString(a)
			sCol.WriteString(")")
			break
		case C_HEX:
			r := RandInt(0, 255)
			g := RandInt(0, 255)
			b := RandInt(0, 255)

			sCol.WriteString("#")
			sCol.WriteString(strconv.FormatInt(int64(r), 16))
			sCol.WriteString(strconv.FormatInt(int64(g), 16))
			sCol.WriteString(strconv.FormatInt(int64(b), 16))
			sCol.WriteString("000000")
			return sCol.String()[0:7]
	}

	return sCol.String()
}
