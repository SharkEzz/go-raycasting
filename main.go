package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const WIDTH int = 1280
const HEIGHT int = 720
const OFFSET = 30

type Game struct {
	boundaries []Boundary
	particle   Particle
}

func (g *Game) Update() error {
	g.setParticleCursorPos(&g.particle)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("FPS: ", math.Round(ebiten.CurrentFPS())), 5, 5)

	drawBoundaries(screen, g.boundaries)
	g.particle.DrawParticle(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := initGame()

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Raycasting")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func initGame() *Game {
	game := Game{}

	// Boundaries
	game.boundaries = make([]Boundary, 4)
	game.boundaries = []Boundary{
		{OFFSET, OFFSET, float64(WIDTH) - OFFSET, OFFSET},
		{float64(WIDTH - OFFSET), OFFSET, float64(WIDTH - OFFSET), float64(HEIGHT - OFFSET)},
		{float64(WIDTH - OFFSET), float64(HEIGHT - OFFSET), OFFSET, float64(HEIGHT - OFFSET)},
		{OFFSET, float64(HEIGHT - OFFSET), OFFSET, OFFSET},
	}

	// Random boundaries
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 5; i++ {
		game.boundaries = append(game.boundaries, Boundary{
			float64(r1.Intn(WIDTH)),
			float64(r1.Intn(HEIGHT)),
			float64(r1.Intn(WIDTH)),
			float64(r1.Intn(HEIGHT)),
		})
	}

	// Particle
	game.particle = NewParticle(float64(WIDTH)/2, float64(HEIGHT)/2)

	return &game
}

func drawBoundaries(screen *ebiten.Image, boundaries []Boundary) {
	for _, boundary := range boundaries {
		ebitenutil.DrawLine(screen, boundary.StartX, boundary.StartY, boundary.StopX, boundary.StopY, color.RGBA{255, 255, 255, 255})
	}
}

func (g *Game) setParticleCursorPos(particle *Particle) {
	cursorPosX, cursorPosY := ebiten.CursorPosition()

	if cursorPosX < OFFSET {
		cursorPosX = OFFSET
	} else if cursorPosX > WIDTH-OFFSET {
		cursorPosX = WIDTH - OFFSET
	}

	if cursorPosY < OFFSET {
		cursorPosY = OFFSET
	} else if cursorPosY > HEIGHT-OFFSET {
		cursorPosY = HEIGHT - OFFSET
	}

	particle.MoveParticle(float64(cursorPosX), float64(cursorPosY), g.boundaries)
}
