package zoia

import (
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/caption"
)

// Header of a patch.
type Header struct {
	Raw             binary.Raw
	Name            string
	PatchLength     int
	NumberOfModules int
}

// ParseHeader returns a new Header.
func ParseHeader(header binary.Raw) *Header {
	h := &Header{
		Raw:             header,
		Name:            caption.TrimToString(header[binary.OffsetHeaderName : binary.OffsetHeaderName+binary.SizeCaption]),
		PatchLength:     int(header.Uint16(0) * 4), // TODO: not sure that the patch length is on 2 bytes, but it's enough to handle the maximum size.
		NumberOfModules: header.Int(binary.OffsetHeaderNumberOfModules),
	}

	return h
}

// SetName set the name of the patch.
// Invalid character are trimmed and the name is limited to 15 characters.
func (h *Header) SetName(name string) {
	name = caption.TrimString(name)
	if len(name) > binary.SizeCaption {
		name = name[:binary.SizeCaption]
	}
	h.Name = name
	h.Raw.Set(caption.Raw([]byte(name), binary.SizeCaption), binary.OffsetHeaderName)
}
