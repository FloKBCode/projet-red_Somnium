package game

import (
	"fmt"
	"somnium/character"
	"somnium/combat"
	"somnium/shop"
	"somnium/ui"
	"time"
	"somnium/quest"
)

func MainMenu() {
	var player character.Character
	created := false

	for {
		if created {
			ui.ClearScreen(&player)
			CheckAllAchievements(&player) // ✅ Vérifier les succès à chaque retour au menu
		}
		if !created {
			ui.PrintInfo("\n✨ Bienvenue dans le Labyrinthe des Cauchemars...")
			ui.PrintInfo("1. Créer un nouveau personnage")
			ui.PrintInfo("2. Charger une partie")
			var choice int
			ui.PrintInfo("Votre choix : ")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				player = character.CharacterCreation()
				created = true
				ui.PrintSuccess("🎉 Personnage créé avec succès !")
				ui.PressEnterToContinue(&player)
			case 2:
				ui.PrintInfo("\n📂 Chargement de la partie...")
				time.Sleep(1 * time.Second)
				loadedPlayer, err := LoadGame()
				if err != nil {
					ui.PrintError("❌ Impossible de charger la partie")
					ui.PressEnterToContinue(&player)
				} else {
					player = *loadedPlayer
					created = true
					ui.PrintSuccess("✅ Partie chargée avec succès !")
					ui.PressEnterToContinue(&player)
				}
			}
			continue
		}

		displayMenuOptions()
		choice := handleUserInput()

		switch choice {
		case 1:
			ui.PrintInfo("\n🌀 Plongée dans tes souvenirs...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			player.DisplayInfo()
			ui.PressEnterToContinue(&player)
		case 2:
			ui.PrintInfo("\n📦 Tu fouilles ton inventaire :")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			handleInventoryMenu(&player)
		case 3:
			ui.PrintInfo("\n🏪 Le marchand apparaît dans un éclair de lumière...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			shop.MerchantMenu(&player)
		case 4:
			ui.PrintInfo("\n⚒️ Dans la forge résonne le métal...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			shop.ForgeMenu(&player)
		case 5:
			ui.PrintInfo("\n⚔️ Tu t'entraînes dans une arène onirique...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			combat.TrainingFight(&player)
			ui.PressEnterToContinue(&player)
		case 6:
			ui.PrintInfo("\n📜 Quêtes disponibles :")
			quest.ShowQuestMenu(&player)
			ui.PressEnterToContinue(&player)
		case 7:
			ui.PrintInfo("\n🌀 Exploration d'une couche du Labyrinthe...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			if err := ExploreLayer(&player); err != nil {
				ui.PrintError(fmt.Sprintf("❌ Erreur : %v", err))
			}
			ui.PressEnterToContinue(&player)
		case 8:
			ui.PrintInfo("\n💾 Sauvegarde de la partie...")
			time.Sleep(1 * time.Second)
			if err := SaveGame(&player); err != nil {
				ui.PrintError("❌ Erreur de sauvegarde")
			} else {
				ui.PrintSuccess("✅ Partie sauvegardée !")
			}
			ui.PressEnterToContinue(&player)
		case 9:
			ui.PrintInfo("\n📂 Chargement de la partie...")
			time.Sleep(1 * time.Second)
			loadedPlayer, err := LoadGame()
			if err != nil {
				ui.PrintError("❌ Impossible de charger la partie")
			} else {
				player = *loadedPlayer
				ui.PrintSuccess("✅ Partie chargée avec succès !")
			}
			ui.PressEnterToContinue(&player)
		case 10:
			ui.PrintInfo("\n🌙 Ton esprit pénètre dans le sanctuaire des songes accomplis...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			ShowAchievementsMenu(&player) 
			ui.PressEnterToContinue(&player)
		case 11:
			DisplayHiddenArtists()
			ui.PressEnterToContinue(&player)
		case 0:
			ui.PrintInfo("\n🌙 Ton esprit retourne doucement dans le coma...")
			ui.PressEnterToContinue(&player)
			ui.ClearScreen(&player)
			return
		default:
			ui.PrintError("❌ Choix invalide, réessaie.")
		}
	}
}

func displayMenuOptions() {
	ui.PrintInfo("\n=== Menu Principal ===")
	ui.PrintInfo("1. Afficher informations personnage")
	ui.PrintInfo("2. Accéder à l'inventaire")
	ui.PrintInfo("3. Marchand")
	ui.PrintInfo("4. Forgeron")
	ui.PrintInfo("5. Entrainement")
	ui.PrintInfo("6. Quêtes disponibles")
	ui.PrintInfo("7. Exploration d'une couche")
	ui.PrintInfo("8. Sauvegarder la partie")
	ui.PrintInfo("9. Charger une partie")
	ui.PrintInfo("10. Salle des Trophées") // ✅ NOUVEAU
	ui.PrintInfo("11. Qui sont-ils")
	ui.PrintInfo("0. Quitter")
}

func handleUserInput() int {
	var choice int
	ui.PrintInfo("Votre choix: ")
	fmt.Scanln(&choice)
	return choice
}

// ✅ NOUVEAU : Menu d'inventaire amélioré avec gestion des armes
func handleInventoryMenu(player *character.Character) {
	for {
		ui.PrintInfo(fmt.Sprintf("\n=== Inventaire (%d/%d) ===", len(player.Inventory), player.InventorySize))
		
		if len(player.Inventory) == 0 {
			ui.PrintInfo("Votre sac est vide.")
			return
		}

		// Afficher inventaire avec types d'objets
		for i, item := range player.Inventory {
			itemType := ""
			if weapon, isWeapon := character.Weapons[item]; isWeapon {
				itemType = fmt.Sprintf(" ⚔️ (+%d dégâts)", weapon.Damage)
			} else if isEquipment(item) {
				itemType = " 🛡️ (Équipement)"
			} else if isPotion(item) {
				itemType = " 🧪 (Consommable)"
			}
			
			ui.PrintInfo(fmt.Sprintf("%d. %s%s", i+1, item, itemType))
		}

		ui.PrintInfo("A. Équiper une arme")
		ui.PrintInfo("E. Gérer l'équipement")
		ui.PrintInfo("0. Retour")

		var choice string
		ui.PrintInfo("Utiliser quel objet ? ")
		fmt.Scanln(&choice)

		switch choice {
		case "0":
			return
		case "A", "a":
			handleWeaponEquip(player)
		case "E", "e":
			player.DisplayEquipment()
			ui.PressEnterToContinue(player)
		default:
			// Essayer de convertir en nombre pour utiliser un objet
			var itemChoice int
			if _, err := fmt.Sscanf(choice, "%d", &itemChoice); err == nil {
				if itemChoice >= 1 && itemChoice <= len(player.Inventory) {
					item := player.Inventory[itemChoice-1]
					useInventoryItem(player, item)
				} else {
					ui.PrintError("❌ Choix invalide")
				}
			} else {
				ui.PrintError("❌ Choix invalide")
			}
		}
	}
}

// ✅ NOUVEAU : Gestion équipement d'arme depuis l'inventaire
func handleWeaponEquip(player *character.Character) {
	weapons := []string{}
	weaponIndices := []int{}
	
	for i, item := range player.Inventory {
		if _, isWeapon := character.Weapons[item]; isWeapon {
			weapons = append(weapons, item)
			weaponIndices = append(weaponIndices, i)
		}
	}
	
	if len(weapons) == 0 {
		ui.PrintError("❌ Aucune arme dans votre inventaire !")
		return
	}
	
	ui.PrintInfo("\n--- Armes disponibles ---")
	for i, weapon := range weapons {
		weaponData := character.Weapons[weapon]
		ui.PrintInfo(fmt.Sprintf("%d. %s (+%d dégâts)", i+1, weapon, weaponData.Damage))
	}
	
	var choice int
	ui.PrintInfo("👉 Quelle arme équiper ? ")
	fmt.Scanln(&choice)
	
	if choice >= 1 && choice <= len(weapons) {
		selectedWeapon := weapons[choice-1]
		player.EquipWeapon(selectedWeapon)
	} else {
		ui.PrintError("❌ Choix invalide")
	}
}

// ✅ NOUVEAU : Utilisation améliorée des objets
func useInventoryItem(player *character.Character, item string) {
	// Vérifier si c'est une arme
	if _, isWeapon := character.Weapons[item]; isWeapon {
		ui.PrintInfo(fmt.Sprintf("Voulez-vous équiper %s ? (o/n)", item))
		var equipChoice string
		fmt.Scanln(&equipChoice)
		if equipChoice == "o" || equipChoice == "oui" {
			player.EquipWeapon(item)
		}
		return
	}
	
	// Vérifier si c'est un équipement
	if isEquipment(item) {
		ui.PrintInfo(fmt.Sprintf("Voulez-vous équiper %s ? (o/n)", item))
		var equipChoice string
		fmt.Scanln(&equipChoice)
		if equipChoice == "o" || equipChoice == "oui" {
			player.EquipItem(item)
		}
		return
	}
	
	// Utiliser l'objet normalement
	switch item {
	case "Potion de vie":
		if player.PvCurr >= player.PvMax {
			ui.PrintError("💖 Vos PV sont déjà au maximum !")
			return
		}
		ui.PrintSuccess("💖 Vos PV augmentent !")
	case "Potion de mana":
		if player.ManaCurr >= player.ManaMax {
			ui.PrintError("🔮 Votre mana est déjà au maximum !")
			return
		}
		ui.PrintSuccess("🔮 Votre mana augmente !")
	}
	
	player.UseItem(item)
}

// Fonctions utilitaires
func isEquipment(item string) bool {
	equipmentItems := []string{
		"Chapeau de l'Errant", "Chapeau de l'aventurier",
		"Tunique des Songes", "Tunique de l'aventurier", 
		"Bottes de l'Oublié", "Bottes de l'aventurier",
	}
	
	for _, equipment := range equipmentItems {
		if item == equipment {
			return true
		}
	}
	return false
}

func isPotion(item string) bool {
	potions := []string{
		"Potion de vie", "Potion de mana", "Potion de poison",
	}
	
	for _, potion := range potions {
		if item == potion {
			return true
		}
	}
	return false
}

// --- Affichage des artistes cachés ---
func DisplayHiddenArtists() {
	ui.PrintInfo("\n🎨 Les artistes cachés sont : ABBA et Spielberg (détails dans les missions bonus).")
}