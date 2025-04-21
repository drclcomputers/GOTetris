// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package util

const (
	WIDTH       = 10
	HEIGHT      = 20
	MODIFYSCORE = 1000
	MINSPEED    = 50

	BLACK   = "\033[0m"
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	CYAN    = "\033[36m"
	WHITE   = "\033[37m"

	BG_BLACK   = "\033[40m"
	BG_RED     = "\033[41m"
	BG_GREEN   = "\033[42m"
	BG_YELLOW  = "\033[43m"
	BG_BLUE    = "\033[44m"
	BG_MAGENTA = "\033[45m"
	BG_CYAN    = "\033[46m"
	BG_WHITE   = "\033[47m"

	GAMEOVERMUSIC   = "assets/gameover2.wav"
	BACKGROUNDMUSIC = "assets/background.wav"
)

var Tetraminos = [][][]int{
	{{1, 1, 1, 1}},         // I
	{{1, 1}, {1, 1}},       // #
	{{0, 1, 0}, {1, 1, 1}}, // T
	{{1, 0, 0}, {1, 1, 1}}, // L
	{{0, 0, 1}, {1, 1, 1}}, // J
	{{0, 1, 1}, {1, 1, 0}}, // S
	{{1, 1, 0}, {0, 1, 1}}, // Z
}

var (
	PRINTMODE    = 3 // 3: no color, 1: background color, 2: foreground color, 4: Electronika 60
	SOUND        = false
	INITIALSPEED = 165

	AUX = 0

	TERM_HEIGHT = 24
	TERM_WIDTH  = 80
)

const TITLE = `
   ██████   ██████  ████████ ███████ ███████ █████   ██  ████
  ██       ██    ██    ██    ██         ██   ██   ██ ██ ██   
  ██  ████ ██    ██    ██    █████      ██   █████   ██   ██
  ██    ██ ██    ██    ██    ██         ██   ██   ██ ██     ██ 
   ██████   ██████     ██    ███████    ██   ██   ██ ██  ████        
`

const VER = "0.6"
const SLOGAN = "     From RSR with love! <3"

var FG_COLORS = []string{BLACK, RED, GREEN, YELLOW, BLUE, MAGENTA, CYAN, WHITE}
var BG_COLORS = []string{BG_BLACK, BG_RED, BG_GREEN, BG_YELLOW, BG_BLUE, BG_MAGENTA, BG_CYAN, BG_WHITE}
