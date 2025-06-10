package mathutils

import (
	"math"
	"testing"
)

func TestLerp(t *testing.T) {
	tests := []struct {
		start, end, t, want float64
	}{
		{0, 10, 0, 0},
		{0, 10, 1, 10},
		{0, 10, 0.5, 5},
		{5, 15, 0.25, 7.5},
		{-10, 10, 0.75, 5},
		{100, 200, 0.33, 133},
		{1, 1, 0.5, 1},
		{0, 10, -0.5, -5},
		{0, 10, 1.5, 15},
	}
	for _, tt := range tests {
		got := Lerp(tt.start, tt.end, tt.t)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("Lerp(%v, %v, %v) = %v; want %v", tt.start, tt.end, tt.t, got, tt.want)
		}
	}
}

func TestLinSpace(t *testing.T) {
	t.Run("basic LinSpace test", func(t *testing.T) {
		min := 0.0
		max := 10.0
		n := 5
		want := []float64{0, 2.5, 5, 7.5, 10}
		got := LinSpace(min, max, n)
		for i := range got {
			if math.Abs(got[i]-want[i]) > 1e-9 {
				t.Errorf("LinSpace(%v, %v, %v)[%d] = %v; want %v", min, max, n, i, got[i], want[i])
			}
		}
	})
}

func TestOppositeAngle(t *testing.T) {
	tests := []struct {
		angle, want float64
	}{
		{0, math.Pi},
		{math.Pi / 2, 3 * math.Pi / 2},
		{math.Pi, 0},
		{3 * math.Pi / 2, math.Pi / 2},
		{math.Pi / 4, 5 * math.Pi / 4},
	}
	for _, tt := range tests {
		got := OppositeAngle(tt.angle)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("OppositeAngle(%v) = %v; want %v", tt.angle, got, tt.want)
		}
	}
}
