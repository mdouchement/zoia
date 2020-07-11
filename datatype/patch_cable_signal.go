package datatype

import (
	"fmt"
)

// PatchCableSignal defines the type of signal used.
// The signal (left, right, cv, etc.) is dynamic depending on cables definition in the patch.
type PatchCableSignal uint8

func (t PatchCableSignal) String() string {
	return fmt.Sprintf("io#%d", t)
}

// MarshalYAML implements yaml.Marshaler.
func (t PatchCableSignal) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}
