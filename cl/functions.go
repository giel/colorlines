package cl

import (
	"fmt"
	"math/rand"
	"strings"
)

// ANSI color codes
var colors = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
	"\033[37m", // White
}

const reset = "\033[0m"

// Generate a random uppercase letter (A-Z)
func randomChar() string {
	return string(rune(rand.Intn(26) + 'A'))
}

// Generate a random color from the list
func randomColor() string {
	return colors[rand.Intn(len(colors))]
}

// Generate a random block-like character
func randomBlock() string {
	blockChars := []string{
		"█", "▓", "▒", "░", // Full, medium, light blocks
		"▄", "▀", "▌", "▐", // Half blocks
		"▖", "▗", "▘", "▝", // Quarter blocks
		"▞", "▟", "▙", "▚", // Three-quarter blocks
		"▂", "▃", "▅", "▆", "▇", // Horizontal bar blocks
	}
	return blockChars[rand.Intn(len(blockChars))]
}

// Generate a random smiley character
func randomSmiley() string {
	blockChars := []string{
		"😀", "😃", "😄", "😁", // Smiling faces
		"😅", "😂", "🤣", "🥲", // Tear faces
	}
	return blockChars[rand.Intn(len(blockChars))]
}

// Display random colored block
func DisplayRandomColoredBlocks(numColors int, numLines int, pieceSize int, shape string) {
	var char1 string
	for line := 1; line <= numLines; line++ {
		for i := 0; i < numColors; i++ {
			color := randomColor()
			if shape == "blocks" {
				char1 = randomBlock() // Use a block-like character
			}
			if shape == "chars" {
				char1 = randomChar() // Use a regular character
			}
			if shape == "smileys" {
				char1 = randomSmiley() // Use a smiley character
			}
			fmt.Printf("%s%s%s", color, strings.Repeat(char1, pieceSize), reset)
		}
		fmt.Println()
	}
}
