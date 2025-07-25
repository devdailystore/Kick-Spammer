package utils

import (
	"math/rand/v2"
	"os"
	"strings"
)

func GetRandomFromFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	lines := strings.Split(string(file), "\n")
	if len(lines) == 0 {
		return ""
	}

	randomIndex := rand.IntN(len(lines))
	if randomIndex < 0 || randomIndex >= len(lines) {
		return ""
	}
	return strings.TrimSpace(lines[randomIndex])
}
