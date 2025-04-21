// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"golang.org/x/term"
)

// Utility functions
func Beep() { fmt.Print("\a") }

func checkSpeaker() bool {
	sampleRate := beep.SampleRate(44100)
	err := speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	if err != nil {
		fmt.Println("No speaker detected or unable to initialize audio output:", err)
		return false
	}
	return true
}

func PlayMusic(sound string, times int) {
	if !checkSpeaker() {
		return
	}

	f, err := os.Open(sound)
	if err != nil {
		return
	}
	defer f.Close()
	streamer, format, err := wav.Decode(f)
	if err != nil {
		return
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(beep.Loop(times, streamer), beep.Callback(func() {
		done <- true
	})))
	<-done
}

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}

func GoAtTopLeft() { fmt.Print("\033[H") }

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to clear screen: %v\n", err)
	}
}

func GetTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Warning: Unable to determine terminal size, using default 80x24.")
		width, height = 80, 24
	}
	return width, height
}
