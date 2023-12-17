package utils

type C struct {
	I int
	J int
}

func (coordinate C) Add(other C) C {
	return C{
		I: coordinate.I + other.I,
		J: coordinate.J + other.J,
	}
}

func (coordinate C) DistanceTo(to C) int {
	return Abs(coordinate.I-to.I) + Abs(coordinate.J-to.J)
}

type Direction = C

var Down = Direction{I: 1, J: 0}
var Up = Direction{I: -1, J: 0}
var Right = Direction{I: 0, J: 1}
var Left = Direction{I: 0, J: -1}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
