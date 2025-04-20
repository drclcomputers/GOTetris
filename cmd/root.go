// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	printMode string
	sound     bool
	diff      string
)

var rootCmd = &cobra.Command{
	Use:   "gotetris",
	Short: "Go Tetris is a terminal Tetris game",
	Run: func(cmd *cobra.Command, args []string) {
		// Show help if no subcommand is provided
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&printMode, "printmode", "p", "nocolor", "Print mode: (1)background, (2)foreground, (3)nocolor")
	rootCmd.PersistentFlags().BoolVarP(&sound, "sound", "s", false, "Enable sound")
	rootCmd.PersistentFlags().StringVarP(&diff, "speed", "d", "intermediate", "Difficulty: (1)easy, (2)intermediate, (3)hard")
}
