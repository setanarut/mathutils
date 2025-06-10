package main

import (
	"fmt"

	"github.com/setanarut/mathutils"
)

func main() {
	fmt.Println(mathutils.LinSpace(0, 3, 3))
	fmt.Println(mathutils.Fract(-1.3234))
	fmt.Println(mathutils.Fract(1.3234))
	fmt.Println(mathutils.Lerp(0, 10, 0.6))
	a := mathutils.OppositeAngle(mathutils.Radians(-90))
	b := mathutils.Degrees(a)
	fmt.Println(b)
}
