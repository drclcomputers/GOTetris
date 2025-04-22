// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package game

import (
	"fmt"
	"gotetris/internal/util"
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
	PrintMode        int
	Sound            bool
	Stop             bool
	Pause            bool
	InputChan        chan keyboard.KeyEvent
}

func NewGame() *Game {
	return &Game{
		PrintMode: util.PRINTMODE,
		Sound:     util.SOUND,
		Pause:     false,
		InputChan: make(chan keyboard.KeyEvent),
	}
}

func (g *Game) pollInput() {
	for !g.Stop {
		char, key, err := keyboard.GetKey()
		if err != nil {
			//fmt.Println("Error reading keyboard input:", err)
			continue
		}
		g.InputChan <- keyboard.KeyEvent{Key: key, Rune: char}
	}
}

func (g *Game) handleKeyEvent(key keyboard.Key, char rune) {
	switch {
	case (key == keyboard.KeyEsc || char == 'q' || char == 'Q') && !g.Pause:
		g.Stop = true
	case (key == keyboard.KeyArrowLeft || char == 'a' || char == 'A') && !g.Pause:
		if g.canMove(g.CurrentShape, g.PosX-1, g.PosY) {
			g.PosX--
		}
	case (key == keyboard.KeyArrowRight || char == 'd' || char == 'D') && !g.Pause:
		if g.canMove(g.CurrentShape, g.PosX+1, g.PosY) {
			g.PosX++
		}
	case (key == keyboard.KeyArrowDown || char == 's' || char == 'S') && !g.Pause:
		if g.canMove(g.CurrentShape, g.PosX, g.PosY+1) {
			g.PosY++
		}
	case (char == 'm' || char == 'M' || char == 'r' || char == 'R') && !g.Pause:
		rotated := rotate(g.CurrentShape)
		if g.canMove(rotated, g.PosX, g.PosY) {
			g.CurrentShape = rotated
		}
	case key == keyboard.KeySpace:
		for g.canMove(g.CurrentShape, g.PosX, g.PosY+1) {
			g.PosY++
		}
	case char == 'p' || char == 'P':
		g.Pause = !g.Pause
	default:
	}
}

func (g *Game) gameLoop() {
	renderTicker := time.NewTicker(33 * time.Millisecond)
	defer renderTicker.Stop()

	dropTicker := time.NewTicker(1 * time.Second)
	defer dropTicker.Stop()

	for !g.Stop {
		select {
		case event := <-g.InputChan:
			g.handleKeyEvent(event.Key, event.Rune)

		case <-renderTicker.C:
			g.drawBoard()

		case <-dropTicker.C:
			if !g.Pause {
				if g.canMove(g.CurrentShape, g.PosX, g.PosY+1) {
					g.PosY++
				} else {
					g.lockToBoard()
					g.clearLines()
					g.randNewPiece()
				}
			}
		}
		if g.Stop {
			fmt.Printf("\033[%d;%dH", util.HEIGHT-5, 6)
			fmt.Print("Press Enter to continue...")
		}
	}
}

func (g *Game) Start() {
	Welcome()
	defer g.Goodbye()
	defer util.ShowCursor()

	if err := keyboard.Open(); err != nil {
		fmt.Println("Error initializing keyboard input:", err)
		os.Exit(1)
	}
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

	g.randNewPiece()

	go g.pollInput()

	g.gameLoop()
}

func Welcome() {
	util.TERM_WIDTH, util.TERM_HEIGHT = util.GetTerminalSize()

	fmt.Print(util.BG_BLACK)

	util.HideCursor()
	util.ClearScreen()

	if util.PRINTMODE == 4 {
		fmt.Print(util.BG_BLACK + util.GREEN)
	}

	fmt.Println("Welcome to GOTetris!")
	fmt.Println()
	fmt.Println("Controls: Arrow keys or WASD to move, M or R to rotate, Space to hard drop, P to pause, Q to quit")
	fmt.Println("Press Enter to start or type 'q' to exit...")

	var input string
	fmt.Scanln(&input)
	if input == "q" {
		fmt.Print(util.BLACK)
		util.ClearScreen()
		os.Exit(0)
	}
	util.ClearScreen()
}

func (g *Game) Goodbye() {
	util.ShowCursor()
	fmt.Print(util.BG_BLACK, util.BLACK)
	util.ClearScreen()

	if g.Sound {
		go util.PlayMusic(util.GAMEOVERMUSIC, 1)
	}

	fmt.Printf("Game Over!\nYour score is: %d\n\nThank you for playing!\n", g.Score)

	time.Sleep(2 * time.Second)

}
