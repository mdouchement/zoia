package datatype

import (
	"strconv"

	"github.com/mdouchement/zoia/binary"
)

// Color defines a module color.
type Color uint8

func (t Color) String() string {
	switch t {
	case binary.ColorExtraBlue:
		return "blue"
	case binary.ColorExtraGreen:
		return "green"
	case binary.ColorExtraRed:
		return "red"
	case binary.ColorExtraYello:
		return "yellow"
	case binary.ColorExtraAqua:
		return "aqua"
	case binary.ColorExtraMagenta:
		return "magenta"
	case binary.ColorExtraWhite:
		return "white"
	case binary.ColorExtraOrange:
		return "orange"
	case binary.ColorExtraLime:
		return "lime"
	case binary.ColorExtraSky:
		return "sky"
	case binary.ColorExtraPurple:
		return "purple"
	case binary.ColorExtraSurf:
		return "surf"
	case binary.ColorExtraPink:
		return "pink"
	case binary.ColorExtraPeach:
		return "peach"
	case binary.ColorExtraMango:
		return "mango"
	default:
		return strconv.Itoa(int(t))
	}
}

// HexColor returns the color in hexadecimal format.
func (t Color) HexColor() string {
	switch t {
	case binary.ColorExtraBlue:
		return "#0000ff"
	case binary.ColorExtraGreen:
		return "#00ff00"
	case binary.ColorExtraRed:
		return "#ff0000"
	case binary.ColorExtraYello:
		return "#ffff00"
	case binary.ColorExtraAqua:
		return "#00ffff"
	case binary.ColorExtraMagenta:
		return "#ff00ff"
	case binary.ColorExtraWhite:
		return "#ffffff"
	case binary.ColorExtraOrange:
		return "#ffa500"
	case binary.ColorExtraLime:
		return "#00ff00"
	case binary.ColorExtraSky:
		return "#87ceeb"
	case binary.ColorExtraPurple:
		return "#800080"
	case binary.ColorExtraSurf:
		return "#00c5cd"
	case binary.ColorExtraPink:
		return "#ffc0cb"
	case binary.ColorExtraPeach:
		return "#ffe5b4"
	case binary.ColorExtraMango:
		return "#ffc324"
	default:
		return "#000000"
	}
}

// MarshalYAML implements yaml.Marshaler.
func (t Color) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}
