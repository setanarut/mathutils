package mathutils

import (
	"math"
)

// Tau is two times pi, representing a full circle in radians. https://oeis.org/A019692
const Tau = 6.2831853071795864769252867665590057683943387987502

const (
	degToRad float64 = math.Pi / 180.0
	radToDeg float64 = 180.0 / math.Pi
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
	if x >= 0 {
		return x - math.Floor(x)
	}
	return x - math.Ceil(x)
}

// Clamp returns value clamped to [min, max]
func Clamp(value, min, max float64) float64 {

	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Lerp performs linear interpolation between start and end based on t.
// t should be in the range [0, 1] where 0 returns start and 1 returns end.
// If t is outside this range, it will extrapolate.
func Lerp(start, end, t float64) float64 {
	return start + t*(end-start)
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
	tValues := LinSpace(0, Tau, n)
	for i, t := range tValues {
		tValues[i] = math.Sin(t) * amplitude
	}
	return tValues
}

// OppositeAngle returns opposite angle. Unit is radians.
// It adds π to the angle and wraps it around if it exceeds 2π.
func OppositeAngle(angle float64) float64 {
	result := angle + math.Pi
	if result >= Tau {
		return result - Tau
	}
	return result
}
