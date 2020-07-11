package zoia

import (
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/datatype"
	"github.com/mdouchement/zoia/parameter"
)

// A PatchCable reprents a patch cable between 2 modules.
type PatchCable struct {
	Raw               binary.Raw
	Gain              datatype.Decibel
	SourceModule      int
	SourceSignal      datatype.PatchCableSignal
	DestinationModule int
	DestinationSignal datatype.PatchCableSignal
}

// ParsePatchCable returns a new PatchCable.
func ParsePatchCable(cable binary.Raw) *PatchCable {
	return &PatchCable{
		Raw:               cable,
		Gain:              datatype.Decibel(parameter.ScaleUint16ToFloat32(cable.Uint16(binary.OffsetPatchCableGain), 11200, -100, 12)),
		SourceModule:      cable.Int(binary.OffsetPatchCableModuleSource),
		SourceSignal:      datatype.PatchCableSignal(cable[binary.OffsetPatchCableSignalSource]),
		DestinationModule: cable.Int(binary.OffsetPatchCableModuleDestination),
		DestinationSignal: datatype.PatchCableSignal(cable[binary.OffsetPatchCableSignalDestination]),
	}
}
