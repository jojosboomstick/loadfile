// Package loadfile nothing ti do
package loadfile

import (
	"bufio"
	"os"
)

// LoadFile read strings from file and return it like slice
func LoadFile(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil {
		return nil, nil
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
