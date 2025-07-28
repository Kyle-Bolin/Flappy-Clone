package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	BirdSize     = 30
	Gravity      = 0.5
	JumpPower    = -10
	PipeWidth    = 80
	PipeGap      = 200
	PipeSpeed    = 3
)

type Game struct {
	bird       Bird
	pipes      []Pipe
	score      int
	gameOver   bool
	background color.Color
}

type Bird struct {
	x, y    float64
	velocity float64
}

type Pipe struct {
	x, y    float64
	height  float64
	passed  bool
}

func NewGame() *Game {
	return &Game{
		bird: Bird{
			x:        100,
			y:        ScreenHeight / 2,
			velocity: 0,
		},
		pipes:      []Pipe{},
		score:      0,
		gameOver:   false,
		background: color.RGBA{135, 206, 235, 255}, // Sky blue
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}

	// Bird physics
	g.bird.velocity += Gravity
	g.bird.y += g.bird.velocity

	// Jump on space or click
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.bird.velocity = JumpPower
	}

	// Generate pipes
	if len(g.pipes) == 0 || g.pipes[len(g.pipes)-1].x < ScreenWidth-300 {
		height := rand.Float64()*(ScreenHeight-PipeGap-100) + 50
		g.pipes = append(g.pipes, Pipe{
			x:       ScreenWidth,
			y:       0,
			height:  height,
			passed:  false,
		})
	}

	// Update pipes
	for i := range g.pipes {
		g.pipes[i].x -= PipeSpeed

		// Check collision with bird
		if g.checkCollision(g.pipes[i]) {
			g.gameOver = true
		}

		// Score points
		if !g.pipes[i].passed && g.pipes[i].x < g.bird.x {
			g.score++
			g.pipes[i].passed = true
		}
	}

	// Remove off-screen pipes
	if len(g.pipes) > 0 && g.pipes[0].x < -PipeWidth {
		g.pipes = g.pipes[1:]
	}

	// Check if bird hits boundaries
	if g.bird.y < 0 || g.bird.y > ScreenHeight {
		g.gameOver = true
	}

	return nil
}

func (g *Game) checkCollision(pipe Pipe) bool {
	birdRight := g.bird.x + BirdSize
	birdLeft := g.bird.x
	birdTop := g.bird.y
	birdBottom := g.bird.y + BirdSize

	pipeRight := pipe.x + PipeWidth
	pipeLeft := pipe.x

	// Check horizontal collision
	if birdRight > pipeLeft && birdLeft < pipeRight {
		// Check vertical collision with top pipe
		if birdTop < pipe.height {
			return true
		}
		// Check vertical collision with bottom pipe
		if birdBottom > pipe.height+PipeGap {
			return true
		}
	}

	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(g.background)

	// Draw bird
	ebitenutil.DrawRect(screen, g.bird.x, g.bird.y, BirdSize, BirdSize, color.RGBA{255, 255, 0, 255})

	// Draw pipes
	for _, pipe := range g.pipes {
		// Top pipe
		ebitenutil.DrawRect(screen, pipe.x, pipe.y, PipeWidth, pipe.height, color.RGBA{0, 128, 0, 255})
		// Bottom pipe
		bottomY := pipe.height + PipeGap
		bottomHeight := ScreenHeight - bottomY
		ebitenutil.DrawRect(screen, pipe.x, bottomY, PipeWidth, bottomHeight, color.RGBA{0, 128, 0, 255})
	}

	// Draw score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, basicfont.Face7x13, 10, 30, color.White)

	if g.gameOver {
		gameOverText := "Game Over! Press R to restart"
		text.Draw(screen, gameOverText, basicfont.Face7x13, ScreenWidth/2-100, ScreenHeight/2, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Flappy Bird Clone")

	game := NewGame()

	// Handle restart
	ebiten.SetKeyCallback(func(key ebiten.Key) {
		if key == ebiten.KeyR && game.gameOver {
			*game = *NewGame()
		}
	})

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
} 