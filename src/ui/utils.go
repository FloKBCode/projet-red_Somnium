package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Named interface {
	GetName() string
}

// GetUserInput récupère l'entrée utilisateur avec un prompt personnalisé
// Prend en paramètre le message à afficher et le joueur
// Retourne la chaîne de caractères saisie par l'utilisateur
func GetUserInput(prompt string, player Named) string {
	fmt.Printf("%s — %s\n> ", player.GetName(), prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// GetUserChoice récupère un choix numérique de l'utilisateur
// Prend en paramètre les valeurs minimum et maximum autorisées et le joueur
// Retourne le nombre choisi par l'utilisateur
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

// ValidateInput vérifie si l'entrée utilisateur est valide
// Prend en paramètre l'entrée à valider, les options valides et le joueur
// Retourne true si l'entrée est valide, false sinon
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

// ClearScreen nettoie l'écran du terminal
// Prend en paramètre le joueur pour personnaliser le message
func ClearScreen(player Named) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Printf("%s ouvre une nouvelle page de son rêve...\n", player.GetName())
}

// PressEnterToContinue attend que l'utilisateur appuie sur Entrée
// Prend en paramètre le joueur pour personnaliser le message
func PressEnterToContinue(player Named) {
	fmt.Printf("%s — Appuyez sur Entrée pour continuer...\n", player.GetName())
	bufio.NewReader(os.Stdin).ReadString('\n')
}
