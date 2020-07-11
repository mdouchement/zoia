package io

import (
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/datatype"
	"github.com/mdouchement/zoia/module"
)

// An AudioInput connects audio from the outside world into the grid.
type AudioInput struct {
	module.Base
	Channel datatype.AudioChannel
}

// ParseDedicatedFields parses dedicated fields of the module.
func (m *AudioInput) ParseDedicatedFields() {
	m.Channel = datatype.AudioChannel(m.Raw[binary.OffsetAudioInputChannel])
}

// Params returns an ordered list of parameters.
func (m *AudioInput) Params() (params []module.Parameter) {
	if m.Channel == binary.ChannelLeft {
		params = append(params, module.Parameter{Name: "pedal input L", Pluggable: true})
	}
	if m.Channel == binary.ChannelRight {
		params = append(params, module.Parameter{Name: "pedal input R", Pluggable: true})
	}
	if m.Channel == binary.ChannelStereo {
		params = append(params, module.Parameter{Name: "pedal input L", Pluggable: true})
		params = append(params, module.Parameter{Name: "pedal input R", Pluggable: true})
	}

	return params
}
