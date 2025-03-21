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
	if config.SaveConfig {
		filePath := cl.GetConfigPath()
		// Reset the SaveConfig flag to avoid saving the same configuration again
		config.SaveConfig = false
		cl.WriteConfig(filePath, config)
	}
}
