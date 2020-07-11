package effect

import (
	"time"

	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/datatype"
	"github.com/mdouchement/zoia/module"
	"github.com/mdouchement/zoia/parameter"
)

// A Compressor is a vastly useful audio tool that controls
// your signal level according to changes in input level.
type Compressor struct {
	module.Base
	Signal           datatype.Signal
	NumberOfControls int
	Treshold         datatype.Decibel
	AttackEnabled    bool
	Attack           time.Duration
	ReleaseEnabled   bool
	Release          time.Duration
	RatioEnabled     bool
	Ratio            datatype.Ratio
	Sidechain        datatype.Sidechain
}

// ParseDedicatedFields parses dedicated fields of the module.
func (m *Compressor) ParseDedicatedFields() {
	m.Signal = datatype.Signal(m.Raw[binary.OffsetCompressorSignal])
	m.Sidechain = datatype.Sidechain(m.Raw[binary.OffsetCompressorSidechain])

	m.NumberOfControls = m.Raw.Int(binary.OffsetCompressorNbOfControls)
	paramOffset := binary.OffsetCompressorControlsStart

	m.Treshold = datatype.Decibel(parameter.Uint16ToFloat32(m.Raw.Uint16(paramOffset), -80, 0))

	m.AttackEnabled = m.Raw.Bool(binary.OffsetCompressorAttackEnabled)
	if m.AttackEnabled {
		paramOffset += binary.SizeControl
		v := parameter.Uint16ToFloat32(m.Raw.Uint16(paramOffset), 0, 10)
		m.Attack = time.Duration(v * 1_000_000) // time.Millisecond
	}

	m.ReleaseEnabled = m.Raw.Bool(binary.OffsetCompressorReleaseEnabled)
	if m.ReleaseEnabled {
		paramOffset += binary.SizeControl
		v := parameter.Uint16ToFloat32(m.Raw.Uint16(paramOffset), 0.01, 2)
		m.Release = time.Duration(v * 1_000_000_000) // time.Second
	}

	m.RatioEnabled = m.Raw.Bool(binary.OffsetCompressorRatioEnabled)
	if m.RatioEnabled {
		paramOffset += binary.SizeControl
		m.Ratio = datatype.Ratio{
			Numerator:   float32(1 + float64(m.Raw.Uint16(paramOffset))/0xFFFF*19),
			Denominator: 1,
			IsInfinity: func(n, _ float32) bool {
				return n >= 20
			},
		}
	}
}

// Params returns an ordered list of parameters.
func (m *Compressor) Params() (params []module.Parameter) {
	params = append(params, module.Parameter{Name: "audio inL", Pluggable: true})
	if m.Signal == binary.SignalStereo {
		params = append(params, module.Parameter{Name: "audio inR", Pluggable: true})
	}

	//

	params = append(params, module.Parameter{Name: "treshold", Pluggable: true, Attribute: m.Treshold})
	if m.AttackEnabled {
		params = append(params, module.Parameter{Name: "attack", Pluggable: true, Attribute: m.Attack})
	}
	if m.ReleaseEnabled {
		params = append(params, module.Parameter{Name: "attack", Pluggable: true, Attribute: m.Release})
	}
	if m.RatioEnabled {
		params = append(params, module.Parameter{Name: "attack", Pluggable: true, Attribute: m.Ratio})
	}

	//

	params = append(params, module.Parameter{Name: "audio outL", Pluggable: true})
	if m.Signal == binary.SignalStereo {
		params = append(params, module.Parameter{Name: "audio outR", Pluggable: true})
	}

	return params
}
