package orientation

import "fmt"

const RotationIncrement Orientation = 90

type Orientation int

const (
	N Orientation =  90 * iota
	E Orientation =  90 * iota
	S Orientation =  90 * iota
	W Orientation =  90 * iota
)

func North() (Orientation) {
	return N
}

func East() (Orientation) {
	return E
}

func  South() (Orientation) {
	return S
}

func  West() (Orientation) {
	return W
}

func String2Orientation(orientation string) (o Orientation) {
	switch string(orientation) {
	case "N":
		o = N
	case "S":
		o = S
	case "W":
		o = W
	case "E":
		o = E
	default:
		fmt.Printf("Invalid Orientation %s : Use North as default value", orientation)
		o = N
	}
	return o
}

func Orientation2String(o Orientation) (s string) {
	if o % 360 == 0 {
		return "N"
	} else if o % 270 == 0 {
		 return "W"
	} else if o % 180 == 0 {
		 return "S"
	} else if o % 90 == 0 {
		 return "E"
	} else {
		return ""
	}
}
