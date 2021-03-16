// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.6

package grand

import (
	"fmt"
	"strconv"
	"strings"
)

type colorType int

const (
	ColorRGB colorType = iota + 1
	ColorRGBA
	ColorHEX
	ColorNAME
	ColorHSL
	ColorHSLA
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

func Color(t colorType) string {
	var s strings.Builder

	switch t {
		case ColorNAME:
			r := RandInt(0, len(value) - 1)
			s.WriteString(value[r])
			break

		case ColorRGB:
			r := RandInt(0, 255)
			g := RandInt(0, 255)
			b := RandInt(0, 255)

			s.WriteString("rgb(")
			s.WriteString(strconv.Itoa(r))
			s.WriteString(",")
			s.WriteString(strconv.Itoa(g))
			s.WriteString(",")
			s.WriteString(strconv.Itoa(b))
			s.WriteString(")")
			break

		case ColorRGBA:
			r := RandInt(0, 255)
			g := RandInt(0, 255)
			b := RandInt(0, 255)
			a := fmt.Sprintf("%.2f", Random())

			s.WriteString("rgba(")
			s.WriteString(strconv.Itoa(r))
			s.WriteString(",")
			s.WriteString(strconv.Itoa(g))
			s.WriteString(",")
			s.WriteString(strconv.Itoa(b))
			s.WriteString(",")
			s.WriteString(a)
			s.WriteString(")")
			break

		case ColorHSL:
			h := RandInt(0, 360)
			hs := RandInt(0, 100)
			l := RandInt(0, 100)

			s.WriteString("hsl(")
			s.WriteString(strconv.Itoa(h))
			s.WriteString(",")
			s.WriteString(strconv.Itoa(hs))
			s.WriteString("%,")
			s.WriteString(strconv.Itoa(l))
			s.WriteString("%)")
			break

		case ColorHSLA:
			h := RandInt(0, 360)
			hs := RandInt(0, 100)
			l := RandInt(0, 100)
			a := fmt.Sprintf("%.2f", Random())

			s.WriteString("hsla(")
			s.WriteString(strconv.Itoa(h))
			s.WriteString(",")
			s.WriteString(strconv.Itoa(hs))
			s.WriteString("%,")
			s.WriteString(strconv.Itoa(l))
			s.WriteString("%,")
			s.WriteString(a)
			s.WriteString(")")
			break

		case ColorHEX:
			r := RandInt(0, 255)
			g := RandInt(0, 255)
			b := RandInt(0, 255)

			s.WriteString("#")
			s.WriteString(strconv.FormatInt(int64(r), 16))
			s.WriteString(strconv.FormatInt(int64(g), 16))
			s.WriteString(strconv.FormatInt(int64(b), 16))
			s.WriteString("000")
			return s.String()[0:7]
	}

	return s.String()
}
