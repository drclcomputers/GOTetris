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
	"sort"
	"strconv"
	"strings"
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
	case key == keyboard.KeyEsc || char == 'q' || char == 'Q':
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
	case (key == keyboard.KeyArrowUp || char == 'w' || char == 'W') && !g.Pause:
		rotated := rotate(g.CurrentShape)
		if g.canMove(rotated, g.PosX, g.PosY) {
			g.CurrentShape = rotated
		}
	case key == keyboard.KeySpace && !g.Pause:
		for g.canMove(g.CurrentShape, g.PosX, g.PosY+1) {
			g.PosY++
		}
	case char == 'p' || char == 'P':
		g.Pause = !g.Pause
		util.LASTPAUSESTATE = !util.LASTPAUSESTATE
	default:
	}
}

func (g *Game) gameLoop() {
	renderTicker := time.NewTicker(33 * time.Millisecond)
	defer renderTicker.Stop()

	dropTicker := time.NewTicker(time.Duration(util.DROPSPEED) * time.Second)
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
			gameOverAnimation()
		}
	}
}

func gameOverAnimation() {
	for i := 0; i < 3; i++ {
		fmt.Printf("\033[%d;%dH", util.TERM_HEIGHT/2, util.TERM_WIDTH/2-3)
		if util.PRINTMODE < 3 {
			fmt.Print(util.BG_BLACK + util.RED + "GAME OVER")
		} else if util.PRINTMODE == 3 {
			fmt.Print(util.BG_BLACK + util.WHITE + "GAME OVER")
		} else {
			fmt.Print(util.BG_BLACK + util.GREEN + "GAME OVER")
		}
		time.Sleep(350 * time.Millisecond)
		fmt.Printf("\033[%d;%dH", util.TERM_HEIGHT/2, util.TERM_WIDTH/2-3)
		fmt.Print(strings.Repeat(" ", 9))
		time.Sleep(350 * time.Millisecond)
	}

	fmt.Printf("\033[%d;%dH", util.TERM_HEIGHT/2, util.TERM_WIDTH/2-3)
	if util.PRINTMODE < 3 {
		fmt.Print(util.BG_BLACK + util.RED + "GAME OVER")
	} else if util.PRINTMODE == 3 {
		fmt.Print(util.BG_BLACK + util.WHITE + "GAME OVER")
	} else {
		fmt.Print("GAME OVER")
	}

	fmt.Printf("\033[%d;%dH", 11, 6)
	if util.PRINTMODE <= 3 {
		fmt.Print(util.BG_BLACK + util.WHITE + "Press Enter to continue...")
	} else {
		fmt.Print(util.BG_BLACK + util.GREEN + "Press Enter to continue...")
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

func printInfoStart() {
	if util.PRINTMODE == 4 {
		fmt.Print(util.GREEN)
	}

	fmt.Println("Welcome to GOTetris!")
	fmt.Println()
	fmt.Println("Controls: Arrow keys or WASD to move, W or UP key to rotate, Space to hard drop, P to pause, Q to quit")
	fmt.Println()
	fmt.Println("Highest score: ")
	printHighScores()
	fmt.Println()
	if util.ENDLESS {
		fmt.Println("Endless (Relaxed) Mode: Play at a steady pace without increasing speed.")
		util.PAUSED = "      Paused - Relax and enjoy!"
	} else {
		fmt.Println("Marathon Mode: Play until the board fills up, with increasing speed as the player clears more lines.")
	}
	fmt.Println()

	fmt.Println("Press Enter to start or type 'q' to exit...")

	var input string
	fmt.Scanln(&input)
	if input == "q" {
		fmt.Print(util.BLACK)
		util.ClearScreen()
		os.Exit(0)
	}
}

func Welcome() {
	util.TERM_WIDTH, util.TERM_HEIGHT = util.GetTerminalSize()

	util.HideCursor()
	fmt.Print(util.BG_BLACK + util.WHITE)

	util.ClearScreen()

	printInfoStart()

	fmt.Print(util.BG_BLACK)
	util.ClearScreen()
}

func (g *Game) printInfoEnd() {
	fmt.Println("Game Over!")
	fmt.Println()
	fmt.Println("Highest score: ")
	printHighScores()
	fmt.Printf("Your score is: %d", g.Score)
	fmt.Println()
	fmt.Println("Thank you for playing!")
}

func (g *Game) Goodbye() {
	util.ShowCursor()
	fmt.Print(util.BG_BLACK, util.BLACK)
	util.ClearScreen()

	if g.Sound {
		go util.PlayMusic(util.GAMEOVERMUSIC, 1)
	}

	g.printInfoEnd()

	g.writeHighScores()

	if g.Sound {
		time.Sleep(2 * time.Second)
	}

}

func readHighScores() []int {
	file, _ := os.ReadFile("score.txt")
	lines := strings.Split(string(file), "\n")
	scores := []int{}
	for _, line := range lines {
		if score, err := strconv.Atoi(line); err == nil {
			scores = append(scores, score)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	return scores
}

func printHighScores() {
	fmt.Println("Top Scores:")
	scores := readHighScores()
	for i, score := range scores {
		fmt.Printf("%d. %d\n", i+1, score)
		if i == 4 {
			break
		}
	}
}

func (g *Game) writeHighScores() {
	if g.Score == 0 {
		return
	}

	file, err := os.OpenFile("score.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error saving high score:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d\n", g.Score))
	if err != nil {
		fmt.Println("Error writing high score:", err)
	}
}
