// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package game

import (
	"gotetris/internal/util"
	"math/rand"
)

func (g *Game) canMove(shape [][]int, posX, posY int) bool {
	for y := 0; y < len(shape); y++ {
		for x := 0; x < len(shape[0]); x++ {
			if shape[y][x] != 0 {
				nX := posX + x
				nY := posY + y
				if nY < 0 || nX < 0 || nX >= util.WIDTH || nY >= util.HEIGHT {
					return false
				}
				if nY >= 0 && g.Board[nY][nX] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func rotate(shape [][]int) [][]int {
	height := len(shape)
	width := len(shape[0])
	rotated := make([][]int, width)
	for i := range rotated {
		rotated[i] = make([]int, height)
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rotated[x][height-1-y] = shape[y][x]
		}
	}
	return rotated
}

func (g *Game) lockToBoard() {
	for y := 0; y < len(g.CurrentShape); y++ {
		for x := 0; x < len(g.CurrentShape[0]); x++ {
			if g.CurrentShape[y][x] != 0 {
				g.Board[g.PosY+y][g.PosX+x] = g.CurrentShapeType + 1
			}
		}
	}
}

func (g *Game) clearLines() int {
	cleared := 0
	for y := util.HEIGHT - 1; y >= 0; y-- {
		full := true

		for x := 0; x < util.WIDTH; x++ {
			if g.Board[y][x] == 0 {
				full = false
				break
			}
		}

		if full {
			for row := y; row > 0; row-- {
				g.Board[row] = g.Board[row-1]
			}
			g.Board[0] = [util.WIDTH]int{}
			cleared++
			y++
		}
	}
	g.Score += cleared * 100
	return cleared
}

func (g *Game) randNewPiece() {
	if g.NextShape == nil {
		g.NextShapeType = rand.Intn(len(util.Tetraminos))
		g.NextShape = util.Tetraminos[g.NextShapeType]
	}

	g.CurrentShapeType = g.NextShapeType
	g.CurrentShape = g.NextShape

	g.NextShapeType = rand.Intn(len(util.Tetraminos))
	g.NextShape = util.Tetraminos[g.NextShapeType]

	g.PosX = (util.WIDTH - len(g.CurrentShape[0])) / 2
	g.PosY = 0

	if !g.canMove(g.CurrentShape, g.PosX, g.PosY) {
		g.Stop = true
		return
	}
}
