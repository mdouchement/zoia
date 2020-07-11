package zoia

import (
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/datatype"
	"github.com/mdouchement/zoia/module"
)

// Version of Zoia firmware.
const Version = "1.11"

// A Patch is a set of synthesizer modules.
type Patch struct {
	Version     string
	Header      *Header
	Pages       []*Page
	Modules     []module.Module
	PatchCables []*PatchCable
}

// Parse parses the given binary patch.
func Parse(patch []byte) *Patch {
	p := &Patch{
		Header: ParseHeader(binary.ExtractHeader(patch)),
	}

	//
	// Modules

	offset := binary.SizeHeader
	for i := 0; i < p.Header.NumberOfModules; i++ {
		m := binary.ExtractModule(patch, offset)
		p.Modules = append(p.Modules, ParseModule(m))

		offset += len(m)
	}

	//
	// Patch cables

	numberOfPatchCables := int(binary.Raw(patch[offset : offset+2]).Uint16(0))
	offset += binary.SizeNumberOfPatchCables

	for i := 0; i < numberOfPatchCables; i++ {
		cable := binary.ExtractPatchCable(patch, offset)
		p.PatchCables = append(p.PatchCables, ParsePatchCable(cable))

		offset += len(cable)
	}

	//
	// Page details

	mumberOfPages := int(patch[offset])
	offset += binary.SizeNumberOfPages

	for i := 0; i < mumberOfPages; i++ {
		page := binary.ExtractPage(patch, offset)
		p.Pages = append(p.Pages, ParsePage(page))

		offset += len(page)
	}

	offset += binary.SizePagePadding

	//
	// Modules extra

	// Firmware < 1.10
	if offset == p.Header.PatchLength {
		for i, m := range p.Modules {
			e := module.Extra{
				Raw:   make(binary.Raw, binary.SizeModuleExtra),
				Color: m.GetBase().Color,
			}
			e.Raw[binary.OffsetModuleExtraColor] = byte(e.Color)

			p.Modules[i].SetExtra(e)
		}

		return p // <================================================== End Of Parsing 1.09
	}

	// Firmware >= 1.10
	for i := 0; i < p.Header.NumberOfModules; i++ {
		extra := binary.ExtractModuleExtra(patch, offset)
		p.Modules[i].SetExtra(module.Extra{
			Raw:   extra,
			Color: datatype.Color(extra[binary.OffsetModuleExtraColor]),
		})

		offset += len(extra)
	}

	return p
}

// Generate generates the binary patch.
// Non-destructive generation base on Raw modifications only.
func Generate(patch *Patch) []byte {
	return nil // TODO:
}
