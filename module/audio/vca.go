package audio

import (
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/datatype"
	"github.com/mdouchement/zoia/module"
	"github.com/mdouchement/zoia/parameter"
)

// A VCA interprets incoming CV at the level control and boost or cut the volume.
type VCA struct {
	module.Base
	Signal       datatype.Signal
	LevelControl datatype.Decibel
}

// ParseDedicatedFields parses dedicated fields of the module.
func (m *VCA) ParseDedicatedFields() {
	m.Signal = datatype.Signal(m.Raw[binary.OffsetVCASignal])
	m.LevelControl = datatype.Decibel(
		parameter.Uint16ToFloat32Log(m.Raw.Uint16(binary.OffsetVCALevelControl), 20),
	)
}

// Params returns an ordered list of parameters.
func (m *VCA) Params() (params []module.Parameter) {
	params = append(params, module.Parameter{Name: "audio in1", Pluggable: true})
	if m.Signal == binary.SignalStereo {
		params = append(params, module.Parameter{Name: "audio in2", Pluggable: true})
	}

	//

	params = append(params, module.Parameter{Name: "level control", Pluggable: true, Attribute: m.LevelControl})

	//

	params = append(params, module.Parameter{Name: "audio out1", Pluggable: true})
	if m.Signal == binary.SignalStereo {
		params = append(params, module.Parameter{Name: "audio out2", Pluggable: true})
	}

	return params
}
