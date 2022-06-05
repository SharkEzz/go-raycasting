package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/SharkEzz/go-raycasting/utils"
)

type Particle struct {
	PosX, PosY, Rotation float64
	Rays                 []Ray
	Scene                *[]float64
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

func (p *Particle) MoveParticle(posX, posY float64, boundaries *[]Boundary) {
	p.PosX = posX
	p.PosY = posY

	var scene []float64

	for index := range p.Rays {
		p.Rays[index].SetOrigin(utils.Point2D{
			X: posX,
			Y: posY,
		})
		p.Rays[index].SetDirection(utils.ToRadian(p.Rotation))

		record := math.Inf(0)
		var closest *utils.Point2D

		for _, boundary := range *boundaries {
			intersect := p.Rays[index].Cast(&boundary)
			if intersect == nil {
				continue
			}

			distance := math.Sqrt(math.Pow(intersect.X-p.PosX, 2) + math.Pow(intersect.Y-p.PosY, 2))
			if distance >= record {
				continue
			}

			record = distance
			closest = intersect
		}

		if closest != nil {
			p.Rays[index].SetStop(*closest)
		}
		scene = append(scene, record)
	}

	p.Scene = &scene
}

func NewParticle(posX, posY float64) Particle {
	rays := []Ray{}

	for i := -30; i < 40; i += 1 {
		rays = append(rays, NewRay(utils.Point2D{X: posX, Y: posY}, utils.ToRadian(float64(i))))
	}

	return Particle{
		Rays:     rays,
		Rotation: 0,
		PosX:     posX,
		PosY:     posY,
	}
}
