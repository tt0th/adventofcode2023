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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
