package main

import (
	"colorlines/cl"
)

func main() {
	// Load configuration file
	config := cl.GetConfig()
	// Call the function with the parsed arguments
	cl.DisplayRandomColoredBlocks(config.Colors, config.Lines, config.PieceSize, config.Shape)
	// Save the configuration file if needed
	if cl.SaveConfig {
		filePath := cl.GetConfigPath()
		cl.WriteConfig(filePath, config)
	}
}
