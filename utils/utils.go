package utils

import "math"

func ToRadian(value float64) float64 {
	return value * (math.Pi / 180)
}

func MapValue(x, in_min, in_max, out_min, out_max float64) float64 {
	return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min
}
