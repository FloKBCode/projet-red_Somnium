package character

import (
	"fmt"
	"regexp"
	"strings"
)

// validateName vérifie que le nom contient uniquement des lettres.
func validateName(input string) string {
	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	for !re.MatchString(input) {
		fmt.Print("⚠️  Ce nom n'est pas autorisé... Choisis un nom fait uniquement de lettres : ")
		fmt.Scanln(&input)
	}
	return input
}

// normalizeString met la première lettre en majuscule, les autres en minuscule.
func normalizeString(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// selectRace demande au joueur de choisir une race onirique.
func selectRace() string {
	fmt.Println("\n🌙 Choisis la forme que ton esprit adoptera dans le Labyrinthe :")
	fmt.Println("1. Humain – équilibre fragile entre force et volonté.")
	fmt.Println("2. Elfe – une conscience fine et une magie subtile.")
	fmt.Println("3. Nain – une endurance forgée dans la pierre des songes.")
	fmt.Println("4. Spectre – éthéré, mais à la vitalité instable.")
	fmt.Println("5. Abysse – né de l'ombre, puissant mais consumé de l'intérieur.")
	
	var choice int
	for {
		fmt.Print("👉 Ton choix (1-5) : ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			continue
		}
		switch choice {
		case 1:
			return "Humain"
		case 2:
			return "Elfe"
		case 3:
			return "Nain"
		case 4:
			return "Spectre"
		case 5:
			return "Abysse"
		default:
			fmt.Println("❌ Ce reflet ne peut exister ici... recommence.")
		}
	}
}

// selectClass demande au joueur de choisir une classe de combat.
func selectClass() string {
	fmt.Println("\n⚔️  Quelle voie suit ton esprit dans ce cauchemar ?")
	fmt.Println("1. Guerrier – une force brute, une arme lourde.")
	fmt.Println("2. Mage – la maîtrise des arcanes du rêve.")
	fmt.Println("3. Voleur – rapide, précis, insaisissable.")
	fmt.Println("4. Occultiste – manipule les ombres au prix de sa santé.")
	
	var choice int
	for {
		fmt.Print("👉 Ton choix (1-4) : ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			continue
		}
		switch choice {
		case 1:
			return "Guerrier"
		case 2:
			return "Mage"
		case 3:
			return "Voleur"
		case 4:
			return "Occultiste"
		default:
			fmt.Println("❌ Cette voie n'existe pas dans le Labyrinthe...")
		}
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

	// Sécurité pour éviter les stats négatives ou nulles
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
	var rawName string
	fmt.Print("💤 Ton esprit dérive... quel est ton nom dans ce rêve ? ")
	fmt.Scanln(&rawName)

	validatedName := validateName(rawName)
	name := normalizeString(validatedName)

	race := selectRace()
	class := selectClass()

	maxHP, maxMana := getBaseStats(race, class)
	
	// ✅ Correction : utiliser maxHP et maxMana
	hero := InitCharacter(name, race, class, maxHP, maxMana)
	
	fmt.Printf("\n✨ %s... ton reflet prend forme : %s %s.\n", hero.Name, hero.Race, hero.Class)
	fmt.Printf("💖 Vitalité : %d | 🔮 Essence : %d\n", hero.PvMax, hero.ManaMax)
	fmt.Println("Ton voyage commence dans les profondeurs de Somnium...")
	fmt.Println("Souviens-toi : chaque mort dans ce lieu laisse des traces...")

	return hero
}