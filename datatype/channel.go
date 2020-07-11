package datatype

import (
	"strconv"

	"github.com/mdouchement/zoia/binary"
)

// AudioChannel defines the audio channel used.
type AudioChannel uint8

func (t AudioChannel) String() string {
	switch t {
	case binary.ChannelStereo:
		return "stereo"
	case binary.ChannelLeft:
		return "left"
	case binary.ChannelRight:
		return "right"
	default:
		return strconv.Itoa(int(t))
	}
}

// MarshalYAML implements yaml.Marshaler.
func (t AudioChannel) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}
