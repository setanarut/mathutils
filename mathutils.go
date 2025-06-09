package mathutils

import "math"

const (
	degToRad = math.Pi / 180.0
	radToDeg = 180.0 / math.Pi
	twoPi    = math.Pi * 2.0
)

// Radians converts an angle measured in degrees to its value in radians.
func Radians(degrees float64) float64 {
	return degrees * degToRad
}

// Degrees converts an angle measured in radians to its value in degrees.
func Degrees(radians float64) float64 {
	return radians * radToDeg
}

// MapRange maps a value v from one range [a, b] to another range [c, d].
func MapRange(v, a, b, c, d float64) float64 {
	return (v-a)/(b-a)*(d-c) + c
}

// Fract returns the fractional part of x.
func Fract(x float64) float64 {
	return x - math.Floor(x)
}

// Clamp returns value clamped to [low, high]
func Clamp(value, low, high float64) float64 {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}

// LinSpace returns a slice of float64 values spaced evenly between min and max.
//
// If n is less than or equal to 1, it returns a slice with only the min value.
func LinSpace(min, max float64, n int) []float64 {
	if n <= 1 {
		return []float64{min}
	}
	d := max - min
	l := float64(n) - 1
	res := make([]float64, n)
	for i := range res {
		res[i] = (min + (float64(i)*d)/l)
	}
	return res
}

// SinSpace returns a slice of sine values spaced evenly between 0 and 2π.
//
// It generates n points, one period of the sine wave, scaled by the given amplitude.
//
//	amplitude // Amplitude of the sine wave
//	n // Number of points
func SinSpace(amplitude float64, n int) []float64 {
	tValues := LinSpace(0, twoPi, n)
	for i, t := range tValues {
		tValues[i] = math.Sin(t) * amplitude
	}
	return tValues
}
