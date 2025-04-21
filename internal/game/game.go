// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package game

import (
	"fmt"
	"gotetris/internal/util"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
)

type Game struct {
	Board            [util.HEIGHT][util.WIDTH]int
	CurrentShape     [][]int
	CurrentShapeType int
	NextShape        [][]int
	NextShapeType    int
	PosX, PosY       int
	Score            int
	Speed            int
	PrintMode        int
	Sound            bool
	Stop             bool
	Pause            bool
	InputChan        chan keyboard.KeyEvent
}

func NewGame() *Game {
	return &Game{
		Speed:     util.INITIALSPEED,
		PrintMode: util.PRINTMODE,
		Sound:     util.SOUND,
		Pause:     false,
		InputChan: make(chan keyboard.KeyEvent),
	}
}

func (g *Game) asyncInputHandler() {
	err := keyboard.Open()
	if err != nil {
		log.Fatalf("Failed to initialize keyboard input: %v", err)
	}
	defer keyboard.Close()

	for !g.Stop {
		char, key, err := keyboard.GetKey()
		if err != nil {
			//fmt.Println("Error reading keyboard input:", err)
			continue
		}
		g.InputChan <- keyboard.KeyEvent{Rune: char, Key: key}
	}
}

func (g *Game) keyHandler() {
	select {
	case event := <-g.InputChan:
		switch {
		case (event.Key == keyboard.KeyEsc || event.Rune == 'q' || event.Rune == 'Q') && !g.Pause:
			g.Stop = true
		case (event.Key == keyboard.KeyArrowLeft || event.Rune == 'a' || event.Rune == 'A') && !g.Pause:
			if g.canMove(g.CurrentShape, g.PosX-1, g.PosY) {
				g.PosX--
			}
		case (event.Key == keyboard.KeyArrowRight || event.Rune == 'd' || event.Rune == 'D') && !g.Pause:
			if g.canMove(g.CurrentShape, g.PosX+1, g.PosY) {
				g.PosX++
			}
		case (event.Key == keyboard.KeyArrowDown || event.Rune == 's' || event.Rune == 'S') && !g.Pause:
			if g.canMove(g.CurrentShape, g.PosX, g.PosY+1) {
				g.PosY++
			}
		case (event.Key == keyboard.KeySpace || event.Rune == 'r' || event.Rune == 'R') && !g.Pause:
			rotated := rotate(g.CurrentShape)
			if g.canMove(rotated, g.PosX, g.PosY) {
				g.CurrentShape = rotated
			}
		case event.Rune == 'p' || event.Rune == 'P':
			if g.Pause {
				fmt.Printf("\033[%d;%dH", util.HEIGHT-5, 6)
				fmt.Print("      ")
			}
			g.Pause = !g.Pause
		default:
		}
	default:

	}
}

func (g *Game) gameLoop() {
	for !g.Stop {
		for g.Pause {
			g.keyHandler()
			fmt.Printf("\033[%d;%dH", util.HEIGHT-5, 6)
			fmt.Print("Paused")
		}

		g.keyHandler()

		g.drawBoard()

		if g.canMove(g.CurrentShape, g.PosX, g.PosY+1) {
			g.PosY++
		} else {
			g.lockToBoard()
			g.clearLines()
			g.randNewPiece()
		}

		if g.Stop {
			fmt.Printf("\033[%d;%dH", util.HEIGHT-5, 6)
			fmt.Print("Press Enter to continue...")
		}

		time.Sleep(time.Duration(g.Speed) * time.Millisecond)

		g.adjustScore()
	}
}

func (g *Game) Start() {
	Welcome()
	defer g.Goodbye()
	defer util.ShowCursor()
	defer keyboard.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		util.ShowCursor()
		keyboard.Close()
		os.Exit(0)
	}()

	if g.Sound && !util.CheckSpeaker() {
		fmt.Println("Warning: No speaker detected. Sound will be disabled.")
		g.Sound = false
	}

	if g.Sound {
		go util.PlayMusic(util.BACKGROUNDMUSIC, -1)
	}

	rand.Seed(time.Now().UnixNano())

	go g.asyncInputHandler()

	g.randNewPiece()
	g.gameLoop()
}

func Welcome() {
	util.TERM_WIDTH, util.TERM_HEIGHT = util.GetTerminalSize()

	util.HideCursor()
	util.ClearScreen()

	if util.PRINTMODE == 4 {
		fmt.Print(util.GREEN)
	}

	fmt.Println("Welcome to GOTetris!")
	fmt.Println("Controls: Arrow keys to move, Space to rotate, Q to quit")
	fmt.Println("Press Enter to start...")

	var input string
	fmt.Scanln(&input)
	util.ClearScreen()
}

func (g *Game) Goodbye() {
	util.ShowCursor()
	util.ClearScreen()

	if g.Sound {
		go util.PlayMusic(util.GAMEOVERMUSIC, 1)
	}

	fmt.Printf("Game Over!\nYour score is: %d\n\nThank you for playing!\n", g.Score)

	time.Sleep(2 * time.Second)

	fmt.Print(util.BLACK, util.BG_BLACK)
}
