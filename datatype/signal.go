package datatype

import (
	"strconv"

	"github.com/mdouchement/zoia/binary"
)

// Signal defines the type of signal used.
type Signal uint8

func (t Signal) String() string {
	switch t {
	case binary.SignalMono:
		return "1in->1out"
	case binary.SignalStereo:
		return "stereo"
	default:
		return strconv.Itoa(int(t))
	}
}

// MarshalYAML implements yaml.Marshaler.
func (t Signal) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}
