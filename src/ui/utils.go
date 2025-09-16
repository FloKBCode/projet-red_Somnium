package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Interface minimale : pas besoin du struct complet
type Named interface {
	GetName() string
}

func GetUserInput(prompt string, player Named) string {
	fmt.Printf("%s — %s\n> ", player.GetName(), prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetUserChoice(min, max int, player Named) int {
	for {
		input := GetUserInput(fmt.Sprintf("Choisissez un nombre entre %d et %d", min, max), player)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < min || choice > max {
			PrintError(fmt.Sprintf("%s hésite… le choix doit être entre %d et %d", player.GetName(), min, max))
			continue
		}
		return choice
	}
}
