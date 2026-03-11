package utils

import "math"

func ToRadian(value float64) float64 {
	return value * (math.Pi / 180)
}

func MapValue(x, inMin, inMax, outMin, outMax float64) float64 {
	if inMax == inMin {
		return outMin
	}

	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}

func ClampValue(value, minValue, maxValue float64) float64 {
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}
