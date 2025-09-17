package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func ValidateInput(input string, validOptions []string, player Named) bool {
	input = strings.ToLower(strings.TrimSpace(input))
	for _, option := range validOptions {
		if strings.ToLower(option) == input {
			return true
		}
	}
	PrintError(fmt.Sprintf("%s ne comprend pas ce choix.", player.GetName()))
	return false
}

func ClearScreen(player Named) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("%s ouvre une nouvelle page de son rêve...\n", player.GetName())
}

func PressEnterToContinue(player Named) {
	fmt.Printf("%s — Appuyez sur Entrée pour continuer...\n", player.GetName())
	bufio.NewReader(os.Stdin).ReadString('\n')
}
