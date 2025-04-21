# 🎮 GoTetris: A Terminal-Based Tetris Game 🧩

Welcome to **GoTetris**, a fun and nostalgic Tetris game built entirely in Go! 🚀 This project brings the classic block-stacking puzzle game to your terminal, complete with customizable settings, colorful visuals, and optional sound effects. 🎵

---

## ✨ Features

- 🕹️ **Classic Tetris Gameplay**: Relive the timeless fun of Tetris in your terminal.
- 🎨 **Customizable Visuals**: Choose between no color, background color, or foreground color modes.
- 🎵 **Sound Effects**: Enable or disable sound effects and background music.
- ⚡ **Adjustable Difficulty**: Play at your preferred speed—Easy, Intermediate, or Hard.
- ⌨️ **Keyboard Controls**:
  - **Arrow Keys or WASD**: Move and rotate blocks.
  - **Spacebar**: Rotate the current piece.
  - **Q**: Quit the game.

---

## 🚀 Getting Started

### Prerequisites
- Go 1.24.2 or higher installed on your system.
- ALSA development package

### Installation
1. Clone the repository:
```bash
git clone https://github.com/yourusername/gotetris.git
cd gotetris
```

2. Build the project
```bash
go build -o gotetris
```

3. Run the game
```bash
./gotetris start
```

---

## 🎮 How to Play
1. Start the game:
```bash
./gotetris start
```
2. Use the following controls to play:

* Arrow Keys or WASD: Move the blocks left, right, or down.
* Spacebar: Rotate the current block.
* Q or ESC: Quit the game.

3. Stack the blocks to clear lines and score points! 🏆

---

## ⚙️ Command-Line Options
Customize your game experience with the following flags:

Flag |	Description	Default
-----|----------------------------------------------------------------------------
-p, --printmode	| Set the print mode: background, foreground, or nocolor | nocolor
-s, --sound	| Enable or disable sound effects |	false
-d, --speed	| Set the difficulty: easy, intermediate, or hard | intermediate

Example:
```bash
./gotetris start -p background -s true -d hard
```

---

## 📂 Project Structure
* cmd/: Contains the Cobra CLI commands.
* internal/game/: Core game logic and rendering.
* internal/util/: Utility functions and constants.
* assets/: Sound files for background music and game over effects.

---

## 🛠️ Development
Dependencies
This project uses the following Go libraries:

* [Cobra](https://github.com/spf13/cobra) for CLI commands.
* [Beep](https://github.com/faiface/beep) for sound effects.
* [Keyboard](https://github.com/eiannone/keyboard) for keyboard input handling.
Run Locally
1. Install dependencies:
```bash
go mod tidy
```
2. Run the game:
```bash
go run .
```

---

## 🎉 Contributing
We welcome contributions! 🛠️ Feel free to open issues or submit pull requests to improve the game.

---

## 📜 License
This project is licensed under the MIT License. See the LICENSE file for details.

Music by [Melody Ayres-Griffiths](https://pixabay.com/users/melodyayresgriffiths-27269767/?utm_source=link-attribution&utm_medium=referral&utm_campaign=music&utm_content=148250) from [Pixabay](https://pixabay.com//?utm_source=link-attribution&utm_medium=referral&utm_campaign=music&utm_content=148250)

---

## ❤️ Acknowledgments
* Inspired by the classic Tetris game. 🎵
* Built with love using Go. 🐹
