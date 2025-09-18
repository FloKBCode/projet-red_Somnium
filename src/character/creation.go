package character

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"somnium/ui"
	"strings"
)

// validateName vérifie que le nom contient uniquement des lettres (y compris accentuées).
func validateName(input string) string {
	// Utilisation de \p{L} pour accepter toutes les lettres Unicode (Éléonore, José, etc.)
	re := regexp.MustCompile(`^\p{L}+$`)
	reader := bufio.NewReader(os.Stdin)

	for !re.MatchString(input) {
		ui.PrintError("⚠️  Ce nom n'est pas autorisé... Choisis un nom fait uniquement de lettres : ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
	}
	return input
}

// normalizeString met la première lettre en majuscule, les autres en minuscule.
func normalizeString(s string) string {
	if len(s) == 0 {
		return s
	}
	// On sépare au cas où le joueur tape plusieurs mots
	words := strings.Fields(s)
	for i := range words {
		words[i] = strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
	}
	return strings.Join(words, " ")
}

// selectRace demande au joueur de choisir une race onirique.
func selectRace() string {
	fmt.Println("\n🌙 Choisis la forme que ton esprit adoptera dans le Labyrinthe :")
	options := []string{
		"Humain – équilibre fragile entre force et volonté.",
		"Elfe – une conscience fine et une magie subtile.",
		"Nain – une endurance forgée dans la pierre des songes.",
		"Spectre – éthéré, mais à la vitalité instable.",
		"Abysse – né de l'ombre, puissant mais consumé de l'intérieur.",
	}

	for {
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}
		fmt.Print("👉 Ton choix (1-5) : ")

		var choice int
		if _, err := fmt.Scanln(&choice); err == nil && choice >= 1 && choice <= len(options) {
			return strings.Fields(options[choice-1])[0] // renvoie juste le mot clé (ex : "Humain")
		}
		ui.PrintError("❌ Ce reflet ne peut exister ici... recommence.")
	}
}

// selectClass demande au joueur de choisir une classe de combat.
func selectClass() string {
	fmt.Println("\n⚔️  Quelle voie suit ton esprit dans ce cauchemar ?")
	options := []string{
		"Guerrier – une force brute, une arme lourde.",
		"Mage – la maîtrise des arcanes du rêve.",
		"Voleur – rapide, précis, insaisissable.",
		"Occultiste – manipule les ombres au prix de sa santé.",
	}

	for {
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}
		fmt.Print("👉 Ton choix (1-4) : ")

		var choice int
		if _, err := fmt.Scanln(&choice); err == nil && choice >= 1 && choice <= len(options) {
			return strings.Fields(options[choice-1])[0] // renvoie juste le mot clé
		}
		ui.PrintError("❌ Cette voie n'existe pas dans le Labyrinthe...")
	}
}

// getBaseStats retourne les PV/Mana de base en fonction de la race ET de la classe.
func getBaseStats(race, class string) (maxHP, maxMana int) {
	// Base par race
	switch race {
	case "Humain":
		maxHP, maxMana = 100, 50
	case "Elfe":
		maxHP, maxMana = 80, 80
	case "Nain":
		maxHP, maxMana = 120, 30
	case "Spectre":
		maxHP, maxMana = 60, 100
	case "Abysse":
		maxHP, maxMana = 150, 20
	default:
		maxHP, maxMana = 100, 50
	}

	// Ajustement par classe
	switch class {
	case "Guerrier":
		maxHP += 20
		maxMana -= 10
	case "Mage":
		maxHP -= 10
		maxMana += 30
	case "Voleur":
		maxHP -= 5
		maxMana += 10
	case "Occultiste":
		maxHP -= 20
		maxMana += 40
	}

	// Sécurité
	if maxHP < 1 {
		maxHP = 1
	}
	if maxMana < 0 {
		maxMana = 0
	}

	return maxHP, maxMana
}

// CharacterCreation crée un nouveau personnage complet.
func CharacterCreation() Character {
	reader := bufio.NewReader(os.Stdin)

	ui.PrintInfo("💤 Ton esprit dérive... quel est ton nom dans ce rêve ? ")
	rawName, _ := reader.ReadString('\n')
	rawName = strings.TrimSpace(rawName)

	validatedName := validateName(rawName)
	name := normalizeString(validatedName)

	race := selectRace()
	class := selectClass()

	maxHP, maxMana := getBaseStats(race, class)

	hero := InitCharacter(name, race, class, maxHP, maxMana)

	// Message immersif
	fmt.Printf("\n✨ %s... ton reflet prend forme : %s %s.\n", hero.Name, hero.Race, hero.Class)
	fmt.Printf("💖 Vitalité : %d | 🔮 Essence : %d\n", hero.PvMax, hero.ManaMax)
	fmt.Println("Ton voyage commence dans les profondeurs de Somnium...")
	ui.PrintError("Souviens-toi : chaque mort dans ce lieu laisse des traces...")

	return hero
}