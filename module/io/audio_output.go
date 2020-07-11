package io

import (
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/datatype"
	"github.com/mdouchement/zoia/module"
	"github.com/mdouchement/zoia/parameter"
)

// An AudioOutput connects audio from your ZOIA into the outside world.
type AudioOutput struct {
	module.Base
	Channel     datatype.AudioChannel
	GainEnabled bool
	Gain        datatype.Decibel
}

// ParseDedicatedFields parses dedicated fields of the module.
func (m *AudioOutput) ParseDedicatedFields() {
	m.Channel = datatype.AudioChannel(m.Raw[binary.OffsetAudioOutputChannel])

	if m.Raw[0] == 15 {
		m.GainEnabled = true
		m.Gain = datatype.Decibel(parameter.Uint16ToFloat32(m.Raw.Uint16(binary.OffsetAudioOutputGain), -100, 20))
	}
}

// Params returns an ordered list of parameters.
func (m *AudioOutput) Params() (params []module.Parameter) {
	if m.Channel == binary.ChannelLeft {
		params = append(params, module.Parameter{Name: "pedal output L", Pluggable: true})
	}
	if m.Channel == binary.ChannelRight {
		params = append(params, module.Parameter{Name: "pedal output R", Pluggable: true})
	}
	if m.Channel == binary.ChannelStereo {
		params = append(params, module.Parameter{Name: "pedal output L", Pluggable: true})
		params = append(params, module.Parameter{Name: "pedal output R", Pluggable: true})
	}

	if m.GainEnabled {
		params = append(params, module.Parameter{Name: "gain", Attribute: m.Gain})
	}

	return params
}
