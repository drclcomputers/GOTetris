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
	offset := (util.TERM_WIDTH - 2*util.WIDTH) / 2
	var builder strings.Builder

	util.GoAtTopLeft()
	builder.WriteString(util.TITLE + "\n\n")

	for y := 0; y < util.HEIGHT; y++ {
		if util.PRINTMODE == 4 {
			builder.WriteString(util.BG_BLACK + util.GREEN)
		}

		if y != 2 {
			builder.WriteString(strings.Repeat(" ", offset) + "│ ")
		} else {
			if util.PRINTMODE == 1 || util.PRINTMODE == 2 {
				builder.WriteString(util.RED)
			}
			builder.WriteString(util.SLOGAN)
			builder.WriteString(strings.Repeat(" ", offset-len(util.SLOGAN)))
			if util.PRINTMODE == 1 || util.PRINTMODE == 2 {
				builder.WriteString(util.WHITE)
			}
			builder.WriteString("│ ")
		}

		for x := 0; x < util.WIDTH; x++ {
			cell := g.getCellValue(x, y)
			builder.WriteString(g.renderCell(cell))
		}

		if util.PRINTMODE == 4 {
			builder.WriteString(util.BG_BLACK + util.GREEN)
		}
		builder.WriteString(" │")

		if y == 5 {
			builder.WriteString(strings.Repeat(" ", 4))
			builder.WriteString(fmt.Sprintf("Score: %d", g.Score))
		}

		if y == 7 {
			builder.WriteString(strings.Repeat(" ", 4))
			builder.WriteString("Next:")
		}

		builder.WriteString("\n")
	}

	builder.WriteString(strings.Repeat(" ", offset) + "└" + strings.Repeat("─", 2*(util.WIDTH+1)) + "┘")
	fmt.Print(builder.String())

	g.renderNextTetramino()
}

func (g *Game) getCellValue(x, y int) int {
	if g.PosY <= y && y < g.PosY+len(g.CurrentShape) &&
		g.PosX <= x && x < g.PosX+len(g.CurrentShape[0]) &&
		g.CurrentShape[y-g.PosY][x-g.PosX] != 0 {
		return g.CurrentShapeType + 1
	}
	return g.Board[y][x]
}

func (g *Game) renderCell(cell int) string {
	switch g.PrintMode {
	case 1:
		return util.BG_COLORS[cell] + "  " + util.BLACK
	case 2:
		return util.BG_BLACK + util.FG_COLORS[cell] + "[]" + util.BLACK
	case 4:
		if cell == 0 {
			return util.BG_BLACK + util.GREEN + ". "
		}
		return util.BG_BLACK + util.GREEN + "[]"
	default:
		if cell == 0 {
			return ". "
		}
		return "[]"
	}
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
