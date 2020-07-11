package binary

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// Raw is the original values of a patch entity.
type Raw []byte

// Bool converts a byte.
func (r Raw) Bool(at int) bool {
	return r[at] != 0
}

// Int converts a byte.
func (r Raw) Int(at int) int {
	return int(r[at])
}

// Uint8 converts a byte.
func (r Raw) Uint8(at int) uint8 {
	return uint8(r[at])
}

// Uint16 converts a 2 bytes sequence.
func (r Raw) Uint16(at int) uint16 {
	return binary.LittleEndian.Uint16(r[at : at+2])
}

// Uint32 converts a 4 bytes sequence.
func (r Raw) Uint32(at int) uint32 {
	return binary.LittleEndian.Uint32(r[at : at+4])
}

// Set sets the given byte to the raw value.
func (r Raw) Set(payload []byte, at int) {
	for i, b := range payload {
		r[at+i] = b
	}
}

func (r Raw) String() string {
	return hex.EncodeToString([]byte(r))
}

// DebugString returns a pretty printed []byte.
func (r Raw) DebugString() string {
	return fmt.Sprintf("%2d", r)
}

// MarshalYAML implements yaml.Marshaler.
func (r Raw) MarshalYAML() (interface{}, error) {
	return r.String(), nil
}
