package game

import (
	"fmt"
	"somnium/character"
	"somnium/combat"
	"somnium/shop"
	"somnium/ui"
	"time"
)

func MainMenu() {
	var player character.Character
	created := false

	for {
		if created {
			ui.ClearScreen(&player)
		}
		if !created {
			fmt.Println("\n✨ Bienvenue dans le Labyrinthe des Cauchemars...")
			fmt.Println("1. Créer un nouveau personnage")
			fmt.Println("2. Charger une partie")
			var choice int
			fmt.Print("Votre choix : ")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				player = character.CharacterCreation()
				fmt.Println("🎉 Personnage créé avec succès !")
			case 2:
				fmt.Println("\n📂 Chargement de la partie...")
			time.Sleep(1 * time.Second)
			loadedPlayer, err := LoadGame()
			if err != nil {
				fmt.Println("❌ Impossible de charger la partie :", err)
				ui.PressEnterToContinue(&player)
			} else {
				player = *loadedPlayer
				created = true
				fmt.Println("✅ Partie chargée avec succès !")
				ui.PressEnterToContinue(&player)
				ui.ClearScreen(&player)
			}
		}

		displayMenuOptions()
		choice = handleUserInput()

		switch choice {
		case 1:
			fmt.Println("\n🌀 Plongée dans tes souvenirs...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			player.DisplayInfo()
			ui.PressEnterToContinue(&player)
		case 2:
			fmt.Println("\n📦 Tu fouilles ton inventaire :")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			character.AccessInventory(&player)
			ui.PressEnterToContinue(&player)
		case 3:
			fmt.Println("\n🏪 Le marchand apparaît dans un éclair de lumière...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			shop.MerchantMenu(&player)
		case 4:
			fmt.Println("\n⚒️ Dans la forge résonne le métal...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			shop.ForgeMenu(&player)
		case 5:
			fmt.Println("\n⚔️ Tu t'entraînes dans une arène onirique...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			combat.TrainingFight(&player)
		case 6:
			fmt.Println("📜 Exploration d'une couche du Labyrinthe :")
			if err := ExploreLayer(&player); err != nil {
				fmt.Println("❌ Erreur :", err)
			}
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
		case 7:
			fmt.Println("\n📜 Quêtes disponibles :")
			// Afficher les quêtes disponibles (fonctionnalité à implémenter)
		case 8:
			fmt.Println("\n💾 Sauvegarde de la partie...")
			time.Sleep(1 * time.Second)
			SaveGame(&player)
			ui.ClearScreen(&player)
		case 9:
			fmt.Println("\n📂 Chargement de la partie...")
			time.Sleep(1 * time.Second)
			loadedPlayer, err := LoadGame()
			if err != nil {
				fmt.Println("❌ Impossible de charger la partie :", err)
				ui.PressEnterToContinue(&player)
			} else {
				player = *loadedPlayer
				created = true
				fmt.Println("✅ Partie chargée avec succès !")
				ui.PressEnterToContinue(&player)
				ui.ClearScreen(&player)
			}
		case 10:
			DisplayHiddenArtists()
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
		case 11:
			fmt.Println("\n🌙 Ton esprit retourne doucement dans le coma...")
			return
		default:
			fmt.Println("❌ Choix invalide, réessaie.")
		}
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
	fmt.Println("7. Exploration d'une couche")
	fmt.Println("8. Sauvegarder la partie")
	fmt.Println("9. Charger une partie")
	fmt.Println("10. Qui sont-ils")
	fmt.Println("11. Quitter")
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
