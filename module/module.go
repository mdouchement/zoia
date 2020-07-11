package module

import (
	"fmt"

	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/caption"
	"github.com/mdouchement/zoia/datatype"
)

// A Module represents one of Zoia's synth module.
type Module interface {
	// Base returns the Base details of a module.
	GetBase() *Base
	// ParseDedicatedFields parses dedicated fields of the module.
	ParseDedicatedFields()
	// SetExtra set the extra of the module.
	SetExtra(extra Extra)
	// Params returns an ordered list of parameters.
	Params() []Parameter
}

// A Base holds all the common informations about a module.
type Base struct {
	Raw      binary.Raw
	Type     datatype.Type
	Name     string
	Color    datatype.Color
	Page     int
	Position int
	Extra    Extra
}

// A Extra holds extra information about a module.
type Extra struct {
	Raw   binary.Raw
	Color datatype.Color
}

// GetBase returns the Base details of a module.
func (m *Base) GetBase() *Base {
	return m
}

// SetExtra set the extra of the module.
func (m *Base) SetExtra(extra Extra) {
	m.Extra = extra
}

// SetName set the name of the module.
// Invalid character are trimmed and the name is limited to 15 characters.
func (m *Base) SetName(name string) {
	name = caption.TrimString(name)
	if len(name) > binary.SizeCaption {
		name = name[:binary.SizeCaption]
	}
	m.Name = name
	m.Raw.Set(caption.Raw([]byte(name), binary.SizeCaption), len(m.Raw)-1-binary.SizeCaption)
}

////////////////////////
//                    //
// Parameter          //
//                    //
////////////////////////

// A Parameter holds a module's parameter details.
type Parameter struct {
	Name      string
	Attribute fmt.Stringer
	Pluggable bool
}

func (p *Parameter) String() string {
	if p.Attribute == nil {
		return fmt.Sprintf("%s: %s", p.Name, p.Attribute)
	}
	return p.Name
}
