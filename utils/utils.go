package utils

import (
	"math/rand/v2"
	"os"
	"strings"
)

func GetRandomProxy() string {
	proxyFile, err := os.ReadFile("./data/proxies.txt")
	if err != nil {
		return ""
	}

	proxyLines := strings.Split(string(proxyFile), "\n")
	if len(proxyLines) == 0 {
		return ""
	}

	randomIndex := rand.IntN(len(proxyLines))
	if randomIndex < 0 || randomIndex >= len(proxyLines) {
		return ""
	}
	return strings.TrimSpace(proxyLines[randomIndex])
}

func GetRandomToken() string {
	tokenFile, err := os.ReadFile("./data/accs.txt")
	if err != nil {
		return ""
	}

	tokenLines := strings.Split(string(tokenFile), "\n")
	if len(tokenLines) == 0 {
		return ""
	}

	randomIndex := rand.IntN(len(tokenLines))
	if randomIndex < 0 || randomIndex >= len(tokenLines) {
		return ""
	}
	return strings.TrimSpace(tokenLines[randomIndex])
}
