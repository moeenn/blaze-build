package main

import (
	"blazebuild/internal/config"
	"fmt"

	"blazebuild/internal/disk"
	"os"
)

func run() error {
	config, err := config.NewConfigFromFile()
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	matches, err := disk.FindFiles(*config.Root, config.Extensions, config.IgnoredPatterns)
	if err != nil {
		return fmt.Errorf("failed find files: %w", err)
	}

	fmt.Printf("\n\n----------------------------------------------------\n")
	for _, match := range matches {
		fmt.Printf("match: %s\n", match)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
