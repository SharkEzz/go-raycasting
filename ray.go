package main

import (
	"math"

	"github.com/SharkEzz/go-raycasting/utils"
)

type Ray struct {
	StartPos  utils.Point2D
	Direction utils.Point2D
	Angle     float64
	StopX     float64
	StopY     float64
}

// Return the intersection point between a ray and a boundary, nil if there is none
func (r *Ray) Cast(boundary *Boundary) *utils.Point2D {
	x1 := boundary.StartX
	y1 := boundary.StartY
	x2 := boundary.StopX
	y2 := boundary.StopY

	x3 := r.StartPos.X
	y3 := r.StartPos.Y
	x4 := r.StopX
	y4 := r.StopY

	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)

	if den == 0 {
		return nil
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / den
	u := ((x1-x3)*(y1-y2) - (y1-y3)*(x1-x2)) / den

	if t >= 0 && t <= 1 && u >= 0 {
		pX := x1 + t*(x2-x1)
		pY := y1 + t*(y2-y1)

		return &utils.Point2D{
			X: pX, Y: pY,
		}
	}

	return nil
}

func (r *Ray) SetOrigin(origin utils.Point2D) {
	r.StartPos = origin
	r.StopX = origin.X + r.Direction.X*30
	r.StopY = origin.Y + r.Direction.Y*30
}

func (r *Ray) SetStop(stop utils.Point2D) {
	r.StopX = stop.X
	r.StopY = stop.Y
}

func NewRay(startPos utils.Point2D, angle float64) Ray {
	ray := Ray{
		StartPos: startPos,
		Angle:    angle,
		Direction: utils.Point2D{
			X: math.Cos(angle),
			Y: math.Sin(angle),
		},
	}
	ray.StopX = startPos.X + ray.Direction.X*30
	ray.StopY = startPos.Y + ray.Direction.Y*30

	return ray
}
