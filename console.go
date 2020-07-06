package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/eiannone/keyboard"
)


var trim string = "^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^\n"

func NewMapConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 2)
	actions[0] = "Arrow | WASD Keys for Movement"
	actions[1] = "x | q | esc to exit"

	console.Actions = actions
	return console
}


func DisplayConsole() {

	console := Console{}

	// Open keyboard
	if err := keyboard.Open(); err != nil {
		logError(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	// Main Menu Loop
	menuLoop:
	for {
		
		// Generate Map fields
		terminal := ""

		// Print from top left to bottom right
		for y := gameWorld.YMax - 1; y >= 0 ; y-- {
			line := []rune{}
			for x := 0; x < gameWorld.XMax; x++ {
				tile := world.Tile{
					X: x,
					Y: y,
				}
				line = append(line, gameWorld.Tiles[tile])
			}
			terminal += string(line) + "\n"
		}

		// Diplay
		color.Green("%s", trim)
		color.Cyan("%s", terminal)
		color.Green("%s", trim)
		console.DisplayActions()

		char, key, err := keyboard.GetKey()
		if err != nil {
			logError(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
        if key == keyboard.KeyEsc {
			break
		}

		switch {
			// Move Left
			case key == keyboard.KeyArrowLeft, char == 'a':
				fmt.Printf("rune %q, key %X\r\n", char, key)

			// Move Right
			case key == keyboard.KeyArrowRight, char == 'd':
				fmt.Printf("rune %q, key %X\r\n", char, key)

			// Move Up
			case key == keyboard.KeyArrowUp, char == 'w':
				fmt.Printf("rune %q, key %X\r\n", char, key)

			// Move Down
			case key == keyboard.KeyArrowDown, char == 's':
				fmt.Printf("rune %q, key %X\r\n", char, key)

			// Exit
			case key == keyboard.KeyEsc, key == keyboard.KeyCtrlC,  key == keyboard.KeyCtrlD, char == 'q', char == 'x':
				break menuLoop
		}
	}
}

// DisplayActions
func (console *Console) DisplayActions() {

	// List Potential Actions
	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		color.Cyan("%d. %s\n", (number + 1), option)
	}
}

// ChooseAction
func (console *Console) ChooseAction() (int, bool) {

	// Open keyboard
	if err := keyboard.Open(); err != nil {
		logError(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	// List Potential Actions
	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		color.Cyan("%d. %s\n", (number + 1), option)
	}

	// Read action as key input
	char, key, err := keyboard.GetKey()
	if err != nil {
		logError(err)
	}
	fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)


	option, err := strconv.Atoi(string(char))
	fmt.Printf("Casted to action %d\r\n", option)

	switch {
		// Exit
		case key == keyboard.KeyEsc, key == keyboard.KeyCtrlC,  key == keyboard.KeyCtrlD, char == 'q', char == 'x':
			return -1, true
		case err == nil && option >= 0:
			return option, false
		default:
			return -1, false
	}
}

func logError(err error) {
	color.Red("\n\n-------------------\nEncountered Error: %s\n-------------------\n\n", err.Error())
}
