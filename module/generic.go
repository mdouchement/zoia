package module

// A Generic is a fallback used until all modules have their own implementations.
type Generic struct {
	Base
}

// ParseDedicatedFields parses dedicated fields of the module.
func (m *Generic) ParseDedicatedFields() {
}

// Params returns an ordered list of paranmeters.
func (m *Generic) Params() []Parameter {
	return nil
}
