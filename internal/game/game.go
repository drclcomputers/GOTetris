package game

import (
	"fmt"
	"gotetris/internal/util"
	"log"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

type Game struct {
	Board            [util.HEIGHT][util.WIDTH]int
	CurrentShape     [][]int
	CurrentShapeType int
	PosX, PosY       int
	Score            int
	Speed            int
	PrintMode        int
	Sound            bool
	Stop             bool
	InputChan        chan keyboard.KeyEvent
}

func NewGame() *Game {
	return &Game{
		Speed:     util.INITIALSPEED,
		PrintMode: util.PRINTMODE,
		Sound:     util.SOUND,
		InputChan: make(chan keyboard.KeyEvent),
	}
}

func (g *Game) asyncInputHandler() {
	err := keyboard.Open()
	if err != nil {
		log.Fatalf("Failed to initialize keyboard input: %v", err)
	}
	defer keyboard.Close()
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		g.InputChan <- keyboard.KeyEvent{Rune: char, Key: key}
	}
}

func (g *Game) keyHandler() {
	select {
	case event := <-g.InputChan:
		switch {
		case event.Key == keyboard.KeyEsc || event.Rune == 'q' || event.Rune == 'Q':
			g.Stop = true
		case event.Key == keyboard.KeyArrowLeft || event.Rune == 'a' || event.Rune == 'A':
			if g.canMove(g.CurrentShape, g.PosX-1, g.PosY) {
				g.PosX--
			}
		case event.Key == keyboard.KeyArrowRight || event.Rune == 'd' || event.Rune == 'D':
			if g.canMove(g.CurrentShape, g.PosX+1, g.PosY) {
				g.PosX++
			}
		case event.Key == keyboard.KeyArrowDown || event.Rune == 's' || event.Rune == 'S':
			if g.canMove(g.CurrentShape, g.PosX, g.PosY+1) {
				g.PosY++
			}
		case event.Key == keyboard.KeySpace:
			rotated := rotate(g.CurrentShape)
			if g.canMove(rotated, g.PosX, g.PosY) {
				g.CurrentShape = rotated
			}
		}
	default:

	}
}

func (g *Game) gameLoop() {
	for !g.Stop {
		g.keyHandler()

		g.drawBoard()

		if g.canMove(g.CurrentShape, g.PosX, g.PosY+1) {
			g.PosY++
		} else {
			g.lockToBoard()
			g.clearLines()
			g.randNewPiece()
		}

		time.Sleep(time.Duration(g.Speed) * time.Millisecond)

		g.adjustScore()
	}
}

func (g *Game) Start() {
	Welcome()
	defer g.Goodbye()

	if g.Sound {
		go util.PlayMusic(util.BACKGROUNDMUSIC, -1)
	}

	rand.Seed(time.Now().UnixNano())

	go g.asyncInputHandler()

	g.randNewPiece()
	g.gameLoop()
}

func Welcome() {
	util.HideCursor()
	util.ClearScreen()
	fmt.Println("Welcome to GOTetris!")
	fmt.Println("Controls: Arrow keys to move, Space to rotate, Q to quit")
	fmt.Println("Press Enter to start...")

	var input string
	fmt.Scanln(&input)
	util.ClearScreen()
}

func (g *Game) Goodbye() {
	util.ClearScreen()

	if g.Sound {
		go util.PlayMusic(util.GAMEOVERMUSIC, 1)
	}

	fmt.Printf("Game Over!\nYour score is: %d\n\nThank you for playing!", g.Score)
	util.ShowCursor()

	time.Sleep(1 * time.Second)
}
