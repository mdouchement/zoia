package parameter

import "math"

// Uint16ToFloat32 converts value to its float32 representation according to min & max.
func Uint16ToFloat32(value uint16, min, max float32) float32 {
	return ScaleUint16ToFloat32(value, 0xFFFF, min, max)
}

// Float32ToUint16 converts value to its uint16 representation according to min & max.
func Float32ToUint16(value float32, min, max float32) uint16 {
	return ScaleFloat32ToUint16(value, 0xFFFF, min, max)
}

// ScaleUint16ToFloat32 converts value to its float32 representation according to min & max.
// Scale is the upper boundary of the value.
func ScaleUint16ToFloat32(value, scale uint16, min, max float32) float32 {
	minAbs := float32(math.Abs(float64(min)))
	maxAbs := float32(math.Abs(float64(max)))

	p := float32(value) / float32(scale)
	return p*(minAbs+maxAbs) + min
}

// ScaleFloat32ToUint16 converts value to its uint16 representation according to min & max.
// Scale is the upper boundary of the value.
func ScaleFloat32ToUint16(value float32, scale uint16, min, max float32) uint16 {
	minAbs := math.Abs(float64(min))
	maxAbs := math.Abs(float64(max))

	value -= min
	value /= float32(minAbs + maxAbs)
	return uint16(value) * scale
}

// Uint16ToFloat32Log converts value to its logarithm float32 representation according to log scale.
func Uint16ToFloat32Log(value uint16, log float32) float32 {
	v := math.Log10(float64(value) / 0xFFFF)
	return float32(v * float64(log))
}

// Float32LogToUint16 converts value to its linear uint16 representation according to log scale.
func Float32LogToUint16(value, log float32) uint16 {
	v := math.Exp(float64(value)/float64(log)*math.Ln10) * 0xFFFF
	return uint16(math.Round(v))
}
