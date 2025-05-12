package disk

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func FindFiles(
	root string,
	fileExtensions *[]string,
	ignoredPathPatterns *[]string,
) ([]string, error) {
	var matches []string

	fmt.Println("Searching for source and header files...")
	err := filepath.WalkDir(root, func(path string, dir fs.DirEntry, err error) error {
		for _, ignoredPattern := range *ignoredPathPatterns {
			if strings.Contains(path, ignoredPattern) {
				return nil
			}
		}

		if err != nil {
			fmt.Printf("error accessing path: %q\n", path)
			return err
		}

		// ignore directories.
		if dir.IsDir() {
			return nil
		}

		filename := dir.Name()
		for _, ext := range *fileExtensions {
			pattern := fmt.Sprintf("*.%s", ext)
			matched, matchErr := filepath.Match(pattern, filename)
			if matchErr != nil {
				fmt.Printf("failed to match pattern %q with file %q", pattern, filename)
				return err
			}

			if matched {
				matches = append(matches, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return matches, nil
}
