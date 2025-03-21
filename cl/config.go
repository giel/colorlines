package cl

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var AppName = "colorlines"
var Version = "not set" // will be set by build.sh
var ConfigPath = "$HOME/.config/colorlines/config.json"

// Config structure for the program
type Config struct {
	Colors     int    `json:"colors"`
	Lines      int    `json:"lines"`
	PieceSize  int    `json:"piece_size"`
	Shape      string `json:"shape"`
	SaveConfig bool   `json:"save_config"`
}

// write config to file
func WriteConfig(filePath string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	dir := filepath.Dir(filePath)       // Extract directory path
	err = os.MkdirAll(dir, os.ModePerm) // Create directories if they don't exist
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Load configuration from the file
func loadConfig(filePath string) (*Config, error) {
	config := &Config{
		Colors:     6,
		Lines:      2,
		PieceSize:  5,
		Shape:      "blocks",
		SaveConfig: false,
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
	colorsFlag := flag.Int("colors", 5, "Number of colors per line")
	linesFlag := flag.Int("lines", 1, "Number of lines to display")
	pieceSizeFlag := flag.Int("piece_size", 5, "Number of characters per piece")
	shapeFlag := flag.String("shape", "blocks", "Use 'blocks', 'chars' or 'smileys' characters to display")
	versionFlag := flag.Bool("version", false, "Display the version of the program")
	saveConfigFlag := flag.Bool("save_config", false, "Save the configuration to the file")

	flag.Parse()

	// Handle the --version flag
	if *versionFlag {
		fmt.Printf("%s version %s\n", AppName, Version)
		os.Exit(0)
	}

	// Load configuration file
	configFilePath := GetConfigPath()
	config, err := loadConfig(configFilePath)
	if err != nil {
		fmt.Printf("Error loading config file: %v\n", err)
		os.Exit(1)
	}

	// Check values of shape from config
	if config.Shape != "" && !CheckShapeValue(config.Shape) {
		fmt.Println("Invalid shape value in config.json. Use 'blocks', 'chars' or 'smileys'")
		os.Exit(1)
	}

	// Use explicitly provided command-line flags to override the configuration file
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "colors" {
			config.Colors = *colorsFlag
		}
		if f.Name == "lines" {
			config.Lines = *linesFlag
		}
		if f.Name == "piece_size" {
			config.PieceSize = *pieceSizeFlag
		}
		if f.Name == "shape" {
			config.Shape = *shapeFlag
		}
		if f.Name == "save_config" {
			config.SaveConfig = *saveConfigFlag
		}
	})

	// Override empty config values with defaults of command-line arguments
	if config.Colors == 0 {
		config.Colors = *colorsFlag
	}
	if config.Lines == 0 {
		config.Lines = *linesFlag
	}
	if config.PieceSize == 0 {
		config.PieceSize = *pieceSizeFlag
	}
	if config.Shape == "" {
		if !CheckShapeValue(*shapeFlag) {
			fmt.Println("Invalid shape value on commandline. Use 'blocks', 'chars' or 'smileys'")
			os.Exit(1)
		}
		config.Shape = *shapeFlag
	}

	return config
}

func GetConfigPath() string {
	configFilePath := os.ExpandEnv(ConfigPath)
	return configFilePath
}

func CheckShapeValue(shape string) bool {
	shapeValues := []string{"blocks", "chars", "smileys"}
	for _, value := range shapeValues {
		if shape == value {
			return true
		}
	}
	return false
}
