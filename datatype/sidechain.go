package datatype

import (
	"strconv"

	"github.com/mdouchement/zoia/binary"
)

// Sidechain defines the type of sidechain used.
type Sidechain uint8

func (t Sidechain) String() string {
	switch t {
	case binary.SidechainInternal:
		return "internal"
	case binary.SidechainExternal:
		return "external"
	default:
		return strconv.Itoa(int(t))
	}
}

// MarshalYAML implements yaml.Marshaler.
func (t Sidechain) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}
