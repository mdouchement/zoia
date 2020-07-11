package zoia

import (
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/caption"
	"github.com/mdouchement/zoia/datatype"
	modulepkg "github.com/mdouchement/zoia/module"
	"github.com/mdouchement/zoia/module/audio"
	"github.com/mdouchement/zoia/module/effect"
	"github.com/mdouchement/zoia/module/io"
)

// ParseModule returns a new Module.
func ParseModule(module binary.Raw) (m modulepkg.Module) {
	l := len(module)

	base := modulepkg.Base{
		Raw:      module,
		Type:     datatype.Type(module[binary.OffsetModuleID]),
		Name:     caption.TrimToString(module[l-1-binary.SizeCaption : l-1]), // FIXME: do not works for all modules
		Color:    datatype.Color(module[binary.OffsetModuleColor]),
		Page:     module.Int(binary.OffsetModulePage),
		Position: module.Int(binary.OffsetModulePosition),
	}

	switch base.Type {
	case binary.ModuleIDAudioInput:
		m = &io.AudioInput{
			Base: base,
		}
	case binary.ModuleIDAudioOutput:
		m = &io.AudioOutput{
			Base: base,
		}
	case binary.ModuleIDVCA:
		m = &audio.VCA{
			Base: base,
		}
	case binary.ModuleIDCompressor:
		m = &effect.Compressor{
			Base: base,
		}
	default:
		m = &modulepkg.Generic{
			Base: base,
		}
	}

	m.ParseDedicatedFields()
	return m
}
