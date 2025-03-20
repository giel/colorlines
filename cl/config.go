package cl

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var AppName = "colorlines"
var Version = "not set" // will be set by build.sh

// Config structure for the program
type Config struct {
	Colors    int    `json:"colors"`
	Lines     int    `json:"lines"`
	PieceSize int    `json:"piece_size"`
	Shape     string `json:"shape"`
}

// Load configuration from the file
func loadConfig(filePath string) (*Config, error) {
	config := &Config{
		Colors:    6,
		Lines:     2,
		PieceSize: 5,
		Shape:     "blocks",
	}

	// Check if the configuration file exists
	if _, err := os.Stat(filePath); err == nil {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

func GetConfig() *Config {
	// Define command-line flags
	colorsFlag := flag.Int("colors", 6, "Number of colors per line")
	linesFlag := flag.Int("lines", 1, "Number of lines to display")
	pieceSizeFlag := flag.Int("piece_size", 5, "Number of characters per piece")
	shapeFlag := flag.String("shape", "blocks", "Use 'blocks' or 'chars' characters to display")
	versionFlag := flag.Bool("version", false, "Display the version of the program")

	flag.Parse()

	// Handle the --version flag
	if *versionFlag {
		fmt.Printf("colorlines version %s\n", Version)
		os.Exit(0)
	}

	// Load configuration file
	configFilePath := os.ExpandEnv("$HOME/.config/colorlines/config.json")
	config, err := loadConfig(configFilePath)
	if err != nil {
		fmt.Printf("Error loading config file: %v\n", err)
		os.Exit(1)
	}

	// Override config values with command-line arguments if provided
	if *colorsFlag > 0 {
		config.Colors = *colorsFlag
	}
	if *linesFlag > 0 {
		config.Lines = *linesFlag
	}
	if *pieceSizeFlag > 0 {
		config.PieceSize = *pieceSizeFlag
	}
	// default is blocks
	config.Shape = "blocks"
	if *shapeFlag == "chars" {
		config.Shape = "chars"
	}
	return config
}
