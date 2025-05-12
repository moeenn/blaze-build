package config

import (
	"blazebuild/internal/toolbelt"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
)

const CONFIG_NAME = "bb-project.json"

type Config struct {
	Root            *string   `json:"root"`
	Extensions      *[]string `json:"extensions"`
	IgnoredPatterns *[]string `json:"ignoredPatterns"`
}

var defaultConfig Config = Config{
	Root:            toolbelt.Ref("."),
	Extensions:      toolbelt.Ref([]string{"cpp", "hpp"}),
	IgnoredPatterns: toolbelt.Ref([]string{"build", "CMakeFiles"}),
}

func (c *Config) SetDefaults() {
	if c.Extensions == nil || len(*c.Extensions) == 0 {
		c.Extensions = defaultConfig.Extensions
	}

	if c.IgnoredPatterns == nil || len(*c.IgnoredPatterns) == 0 {
		c.IgnoredPatterns = defaultConfig.IgnoredPatterns
	}
}

func NewConfigFromFile() (*Config, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to detect current director: %w", err)
	}

	filepath := path.Join(currentDir, CONFIG_NAME)

	// check if the config file exists.
	_, err = os.Stat(filepath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Printf("config file does not exist: using default options.\n")
			return &defaultConfig, nil
		}
	}

	contentBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read contents of the config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(contentBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file as JSON: %w", err)
	}

	config.SetDefaults()
	return &config, nil
}
