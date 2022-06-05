package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/SharkEzz/go-raycasting/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const WIDTH int = 1400
const HEIGHT int = 720

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
	drawBoundaries(screen, g.boundaries)
	g.particle.DrawParticle(screen)

	scene := *g.particle.Scene

	w := (WIDTH / 2) / len(scene)
	for i := 0; i < len(scene); i++ {

		sq := scene[i] * scene[i]
		wSq := (WIDTH / 2) * (WIDTH / 2)
		c := uint8(utils.MapValue(sq, 0, float64(wSq), 255, 0))
		h := utils.MapValue(1/scene[i], 0, 0.02, 0, float64(HEIGHT))

		ebitenutil.DrawRect(screen, float64(i*w+(WIDTH/2)), (float64(HEIGHT)-h)/2, float64(w), h, color.RGBA{c, c, c, 0xFF})
	}
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
		{0, 0, float64(WIDTH / 2), 0},
		{float64(WIDTH / 2), 0, float64(WIDTH / 2), float64(HEIGHT)},
		{float64(WIDTH / 2), float64(HEIGHT), 0, float64(HEIGHT)},
		{0, float64(HEIGHT), 0, 0},
	}

	// Random boundaries
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 5; i++ {
		game.boundaries = append(game.boundaries, Boundary{
			float64(r1.Intn(WIDTH / 2)),
			float64(r1.Intn(HEIGHT)),
			float64(r1.Intn(WIDTH / 2)),
			float64(r1.Intn(HEIGHT)),
		})
	}

	// Particle
	game.particle = NewParticle(float64(WIDTH/2)/2, float64(HEIGHT)/2)

	return &game
}

func drawBoundaries(screen *ebiten.Image, boundaries []Boundary) {
	for _, boundary := range boundaries {
		ebitenutil.DrawLine(screen, boundary.StartX, boundary.StartY, boundary.StopX, boundary.StopY, color.RGBA{255, 255, 255, 255})
	}
}

func (g *Game) setParticleCursorPos(particle *Particle) {
	mouseX, mouseY := ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		particle.Rotate(-0.03)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		particle.Rotate(0.03)
	}

	particle.MoveParticle(float64(mouseX), float64(mouseY), &g.boundaries)
}
