package ui

import (
	"bufio"
	"fmt"
	"os"
	"somnium/character"
	"strconv"
	"strings"
)

func GetUserInput(prompt string, player *character.Character) string {
	fmt.Printf("%s — %s\n> ", player.Name, prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetUserChoice(min, max int, player *character.Character) int {
	for {
		input := GetUserInput(fmt.Sprintf("Choisissez un nombre entre %d et %d", min, max), player)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < min || choice > max {
			PrintError(fmt.Sprintf("%s hésite… le choix doit être entre %d et %d", player.Name, min, max))
			continue
		}
		return choice
	}
}

func ValidateInput(input string, validOptions []string, player *character.Character) bool {
	for _, option := range validOptions {
		if strings.EqualFold(input, option) {
			return true
		}
	}
	PrintError(fmt.Sprintf("%s ne peut pas choisir cela… (%s)", player.Name, input))
	return false
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PressEnterToContinue(player *character.Character) {
	fmt.Printf("%s flotte dans l'obscurité… appuyez sur Entrée pour continuer.\n", player.Name)
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func DisplayMenu(title string, options []string, player *character.Character) int {
	fmt.Println("═══════════════ " + title + " ═══════════════")
	for i, opt := range options {
		fmt.Printf("%d) %s\n", i+1, opt)
	}
	return GetUserChoice(1, len(options), player)
}
