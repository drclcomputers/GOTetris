// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package cmd

import (
	"gotetris/internal/game"
	"gotetris/internal/util"

	"strings"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Tetris game",
	Run: func(cmd *cobra.Command, args []string) {
		switch strings.ToLower(printMode) {
		case "1", "background":
			util.PRINTMODE = 1
		case "2", "foreground":
			util.PRINTMODE = 2
		case "3", "nocolor":
			util.PRINTMODE = 3
		case "4", "60", "electronika":
			util.PRINTMODE = 4
		default:
			util.PRINTMODE = 3
		}
		util.SOUND = sound

		switch strings.ToLower(diff) {
		case "1", "easy":
			util.INITIALSPEED = 250
		case "2", "intermediate":
			util.INITIALSPEED = 175
		case "3", "hard":
			util.INITIALSPEED = 135
		default:
			util.INITIALSPEED = 175
		}

		g := game.NewGame()
		g.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
