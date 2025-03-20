package main

import (
	"colorlines/cl"
)

func main() {
	// Load configuration file
	config := cl.GetConfig()
	// Call the function with the parsed arguments
	cl.DisplayRandomColoredBlocks(config.Colors, config.Lines, config.PieceSize, config.Shape)
}
