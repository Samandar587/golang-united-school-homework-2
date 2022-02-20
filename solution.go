package main

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type num int

func CalcSquare(sideLen float64, sidesNum num) float64 {

	switch sidesNum {
	case 3:
		return (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	case 4:
		return sideLen * sideLen
	case 0:
		return math.Pi * math.Pow(sideLen, 2)
	default:
		return 0

	}
}

func main() {
	const (
		SidesTriangle = 3
		SidesSquare   = 4
		SidesCircle   = 0
	)
	fmt.Println(CalcSquare(2.0, SidesTriangle))
	fmt.Println(CalcSquare(10.0, SidesSquare))
	fmt.Println(CalcSquare(2.0, SidesCircle))

}
