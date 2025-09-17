package game

import (
	"fmt"
	"somnium/character"
	"somnium/combat"
	"somnium/shop"
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
			fmt.Printf("Personnage créé avec succès !\n")
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
			shop.MerchantMenu(&player)
		case 4:
			fmt.Println("\n⚒️ La forge résonne du métal...")
			shop.ForgeMenu(&player)
		case 5:
			fmt.Println("\n⚔️ Tu t'entraînes dans une arène onirique...")
			combat.TrainingFight(&player)
		case 6:
			fmt.Println("\n📜 Quêtes disponibles :")
			// Afficher les quêtes disponibles (fonctionnalité à implémenter)
		case 7:
			fmt.Println("\n💾 Sauvegarde de la partie...")
			// Sauvegarder la partie (fonctionnalité à implémenter)
		case 8:
			fmt.Println("\n📂 Chargement de la partie...")
			// Charger une partie sauvegardée (fonctionnalité à implémenter)
		case 9:
			DisplayHiddenArtists()
		case 10:
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
	fmt.Println("6. Quêtes disponibles")
	fmt.Println("7. Sauvegarder la partie")
	fmt.Println("8. Charger une partie")
	fmt.Println("9. Qui sont-ils")
	fmt.Println("10. Quitter")
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
