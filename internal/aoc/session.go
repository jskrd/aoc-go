package aoc

import (
	"errors"
	"fmt"
	"os"
)

func GetSession() (string, error) {
	data, err := os.ReadFile("session.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("session file not found")
		}
		return "", fmt.Errorf("error reading session file: %w", err)
	}

	return string(data), nil
}
