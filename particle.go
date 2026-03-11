package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/SharkEzz/go-raycasting/utils"
)

const FOV = 70
const HALF_FOV = FOV / 2

type Particle struct {
	PosX, PosY, Heading float64
	Rays                []Ray
	Scene               []float64
}

func (p *Particle) DrawParticle(screen *ebiten.Image) {
	for _, ray := range p.Rays {
		ebitenutil.DrawLine(
			screen,
			ray.StartPos.X,
			ray.StartPos.Y,
			ray.StopX,
			ray.StopY,
			color.RGBA{255, 255, 255, 255})
	}
}

func (p *Particle) MoveParticle(posX, posY float64, boundaries []Boundary) {
	p.PosX = posX
	p.PosY = posY

	if len(p.Scene) != len(p.Rays) {
		p.Scene = make([]float64, len(p.Rays))
	}

	for index := range p.Rays {
		p.Rays[index].SetOrigin(utils.Point2D{
			X: posX,
			Y: posY,
		})

		record := MAX_VIEW_DISTANCE
		var closest *utils.Point2D

		for _, boundary := range boundaries {
			intersect := p.Rays[index].Cast(boundary)
			if intersect == nil {
				continue
			}

			distance := math.Hypot(intersect.X-p.PosX, intersect.Y-p.PosY)
			if distance >= record {
				continue
			}

			record = distance
			closest = intersect
		}

		if closest != nil {
			p.Rays[index].SetStop(*closest)
		}

		// fish-eye correction using delta between ray angle and camera heading
		correctedDistance := record * math.Cos(p.Rays[index].Angle-p.Heading)
		if correctedDistance < 0 {
			correctedDistance = MAX_VIEW_DISTANCE
		}
		p.Scene[index] = correctedDistance
	}
}

func (p *Particle) Rotate(angle float64) {
	p.Heading += angle

	for index := range p.Rays {
		rayOffset := float64(index - HALF_FOV)
		p.Rays[index].SetAngle(utils.ToRadian(rayOffset) + p.Heading)
	}
}

func NewParticle(posX, posY float64) Particle {
	rays := []Ray{}

	for i := -HALF_FOV; i < HALF_FOV; i += 1 {
		rays = append(rays, NewRay(utils.Point2D{X: posX, Y: posY}, utils.ToRadian(float64(i))))
	}

	return Particle{
		Rays:  rays,
		Scene: make([]float64, len(rays)),
		PosX:  posX,
		PosY:  posY,
	}
}
