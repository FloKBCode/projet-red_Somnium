package game

import (
	"fmt"
	"somnium/character"
)

func MainMenu() {
	var player character.Character
	created := false

	for {
		// Création du personnage si pas encore fait
		if !created {
			player = character.CharacterCreation()
			created = true
			fmt.Println("\n✨ Ton esprit s'éveille dans le Labyrinthe des Cauchemars...")
			fmt.Println("Personnage créé avec succès !\n")
		}

		displayMenuOptions()
		choice := handleUserInput()

		switch choice {
		case 1:
			fmt.Println("\n🌀 Plongée dans tes souvenirs...")
			player.DisplayInfo()
		case 2:
			fmt.Println("\n📦 Tu fouilles ton inventaire :")
			character.AccessInventory(&player)
		case 3:
			fmt.Println("\n🏪 Le marchand apparaît dans un éclair de lumière...")
			fmt.Println("Marchand (à implémenter)")
		case 4:
			fmt.Println("\n⚒️ La forge résonne du métal...")
			fmt.Println("Forgeron (à implémenter)")
		case 5:
			fmt.Println("\n⚔️ Tu t'entraînes dans une arène onirique...")
			fmt.Println("Entrainement (à implémenter)")
		case 6:
			DisplayHiddenArtists()
		case 7:
			fmt.Println("\n🌙 Ton esprit retourne doucement dans le coma...")
			return
		default:
			fmt.Println("❌ Choix invalide, réessaie.")
		}
	}
}

// --- Affichage du menu ---
func displayMenuOptions() {
	fmt.Println("\n=== Menu Principal ===")
	fmt.Println("1. Afficher informations personnage")
	fmt.Println("2. Accéder à l'inventaire")
	fmt.Println("3. Marchand")
	fmt.Println("4. Forgeron")
	fmt.Println("5. Entrainement")
	fmt.Println("6. Qui sont-ils")
	fmt.Println("7. Quitter")
}

// --- Lecture du choix utilisateur ---
func handleUserInput() int {
	var choice int
	fmt.Print("Votre choix: ")
	fmt.Scanln(&choice)
	return choice
}

// --- Affichage des artistes cachés ---
func DisplayHiddenArtists() {
	fmt.Println("\n🎨 Les artistes cachés sont : ABBA et Spielberg (détails dans les missions bonus).")
}
