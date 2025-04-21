// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package game

import (
	"fmt"
	"gotetris/internal/util"
	"strings"
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

	offset := (util.TERM_WIDTH - 2*util.WIDTH) / 2
	util.GoAtTopLeft()

	fmt.Print(util.TITLE)
	fmt.Print("\n\n")

	for y := 0; y < util.HEIGHT; y++ {
		if util.PRINTMODE == 4 {
			fmt.Print(util.BG_BLACK + util.GREEN)
		}

		if y != 2 {
			fmt.Print(strings.Repeat(" ", offset))
			fmt.Print("│ ")
		} else {
			if util.PRINTMODE == 1 || util.PRINTMODE == 2 {
				fmt.Print(util.RED)
			}
			fmt.Print(util.SLOGAN)
			fmt.Print(strings.Repeat(" ", offset-len(util.SLOGAN)))
			if util.PRINTMODE == 1 || util.PRINTMODE == 2 {
				fmt.Print(util.WHITE)
			}
			fmt.Print("│ ")
		}

		for x := 0; x < util.WIDTH; x++ {
			cell := tempBoard[y][x]
			switch g.PrintMode {
			case 1:
				switch cell {
				case 0:
					fmt.Print(util.BG_BLACK + "  ")
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
				fmt.Print(util.BG_BLACK)
				switch cell {
				case 0:
					fmt.Print(util.WHITE + ". ")
				case 1:
					fmt.Print(util.RED + "[]")
				case 2:
					fmt.Print(util.GREEN + "[]")
				case 3:
					fmt.Print(util.YELLOW + "[]")
				case 4:
					fmt.Print(util.BLUE + "[]")
				case 5:
					fmt.Print(util.MAGENTA + "[]")
				case 6:
					fmt.Print(util.CYAN + "[]")
				case 7:
					fmt.Print(util.WHITE + "[]")
				}
			case 4:
				if cell == 0 {
					fmt.Print(util.BG_BLACK, util.GREEN, ". ")
				} else {
					fmt.Print(util.BG_BLACK, util.GREEN, "[]")
				}
			default:
				if cell == 0 {
					fmt.Print(". ")
				} else {
					fmt.Print("[]")
				}
			}
			fmt.Print(util.BG_BLACK + util.WHITE)
		}
		if util.PRINTMODE == 4 {
			fmt.Print(util.BG_BLACK + util.GREEN)
		}
		fmt.Print(" │")

		if y == 5 {
			fmt.Print(strings.Repeat(" ", 4))
			fmt.Printf("Score: %d", g.Score)
		}

		if y == 7 {
			fmt.Print(strings.Repeat(" ", 4))
			fmt.Print("Next:")
		}

		fmt.Println()
	}

	fmt.Print(strings.Repeat(" ", offset))
	fmt.Print("└")
	fmt.Print(strings.Repeat("─", 2*(util.WIDTH+1)))
	fmt.Print("┘")

	g.renderNextTetramino()
}

func (g *Game) renderNextTetramino() {
	nextTetraminoRow := 18
	nextTetraminoCol := util.WIDTH*2 + 10 + (util.TERM_WIDTH-2*util.WIDTH)/2

	for y := 0; y < len(g.NextShape); y++ {
		fmt.Printf("\033[%d;%dH", nextTetraminoRow+y, nextTetraminoCol)
		for x := 0; x < len(g.NextShape[0]); x++ {
			if g.NextShape[y][x] != 0 {
				fmt.Print("[]")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print(strings.Repeat(" ", 5))
	}
	if len(g.NextShape) == 1 {
		fmt.Printf("\033[%d;%dH", nextTetraminoRow+1, nextTetraminoCol)
		fmt.Print(strings.Repeat(" ", 10))
	}
}
