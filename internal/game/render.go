package game

import (
	"fmt"
	"gotetris/internal/util"
)

func (g *Game) drawBoard() {
	tempBoard := g.Board
	for y := range tempBoard {
		tempBoard[y] = g.Board[y]
	}

	for y := 0; y < len(g.CurrentShape); y++ {
		for x := 0; x < len(g.CurrentShape[0]); x++ {
			if g.CurrentShape[y][x] != 0 {
				if g.PosY+y >= 0 && g.PosY+y < util.HEIGHT && g.PosX+x >= 0 && g.PosX+x < util.WIDTH {
					tempBoard[g.PosY+y][g.PosX+x] = g.CurrentShapeType + 1
				}
			}
		}
	}

	util.GoAtTopLeft()

	for y := 0; y < util.HEIGHT; y++ {
		fmt.Print("│ ")
		for x := 0; x < util.WIDTH; x++ {
			cell := tempBoard[y][x]
			switch g.PrintMode {
			case 1:
				switch cell {
				case 0:
					fmt.Print(util.BG_BLACK + util.BLACK + "  ")
				case 1:
					fmt.Print(util.BG_RED + "  ")
				case 2:
					fmt.Print(util.BG_GREEN + "  ")
				case 3:
					fmt.Print(util.BG_YELLOW + "  ")
				case 4:
					fmt.Print(util.BG_BLUE + "  ")
				case 5:
					fmt.Print(util.BG_MAGENTA + "  ")
				case 6:
					fmt.Print(util.BG_CYAN + "  ")
				case 7:
					fmt.Print(util.BG_WHITE + "  ")
				}
			case 2:
				switch cell {
				case 0:
					fmt.Print(util.BLACK)
				case 1:
					fmt.Print(util.RED)
				case 2:
					fmt.Print(util.GREEN)
				case 3:
					fmt.Print(util.YELLOW)
				case 4:
					fmt.Print(util.BLUE)
				case 5:
					fmt.Print(util.MAGENTA)
				case 6:
					fmt.Print(util.CYAN)
				case 7:
					fmt.Print(util.WHITE)
				}
				fmt.Print("[]")
			default:
				if cell == 0 {
					fmt.Print(". ")
				} else {
					fmt.Print("[]")
				}
			}
			fmt.Print(util.BG_BLACK + util.BLACK)
		}
		fmt.Println(" │")
	}
	fmt.Print("└")
	for i := 0; i < 2*(util.WIDTH+1); i++ {
		fmt.Print("─")
	}
	fmt.Print("┘")
	fmt.Printf("\nScore: %d\n", g.Score)
}
