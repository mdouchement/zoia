package datatype

import "fmt"

// Ratio defines a ratio defined by a parameter.
type Ratio struct {
	Numerator   float32
	Denominator float32
	IsInfinity  func(n, d float32) bool
}

func (t Ratio) String() string {
	if t.IsInfinity != nil && t.IsInfinity(t.Numerator, t.Denominator) {
		return "infinity"
	}
	return fmt.Sprintf("%.1f:%.1f", t.Numerator, t.Denominator)
}

// MarshalYAML implements yaml.Marshaler.
func (t Ratio) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}
