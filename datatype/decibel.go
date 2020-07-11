package datatype

import "fmt"

// Decibel defines a value expressed in decibel (dB).
type Decibel float32

func (t Decibel) String() string {
	return fmt.Sprintf("%.2f dB", t)
}

// MarshalYAML implements yaml.Marshaler.
func (t Decibel) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}
