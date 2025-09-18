package character

import (
	"fmt"
	"somnium/ui"
	"strings"
)

// RemoveFromInventory supprime un objet
func (c *Character) RemoveFromInventory(item string) bool {
	for i, invItem := range c.Inventory {
		if invItem == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

// TakePoison consomme une potion de poison
func (c *Character) TakePoison() bool {
	if !c.RemoveFromInventory("Potion de poison") {
		ui.PrintError("❌ Pas de potion de poison !")
		return false
	}
	ui.PrintInfo("☠️ Vous buvez la potion de poison...")
	c.PoisonEffect()
	return true
}

// CanUpgradeInventorySlot vérifie si le personnage peut améliorer son inventaire
func (c *Character) CanUpgradeInventorySlot() (bool, int, string) {
	upgradeCost := 30
	maxUpgrades := 3
	
	if c.XPUpgrades >= maxUpgrades {
		return false, 0, "🚫 Votre sac ne peut pas être agrandi davantage."
	}
	if c.Money < upgradeCost {
		return false, upgradeCost, "💰 Pas assez de fragments pour améliorer votre sac."
	}
	
	return true, upgradeCost, ""
}

// UpgradeInventorySlot améliore la taille de l'inventaire du personnage
func (c *Character) UpgradeInventorySlot() bool {
	upgradeCost := 30
	maxUpgrades := 3

	if c.XPUpgrades >= maxUpgrades {
		ui.PrintError("🚫 Votre sac ne peut pas être agrandi davantage.")
		return false
	}
	if c.Money < upgradeCost {
		ui.PrintError("💰 Pas assez de fragments pour améliorer votre sac.")
		return false
	}

	c.Money -= upgradeCost
	c.InventorySize += 10
	c.XPUpgrades++
	ui.PrintSuccess(fmt.Sprintf("🎒 Votre sac s'élargit (+10 emplacements). Capacité : %d", c.InventorySize))
	return true
}

// ✅ AMÉLIORÉ : AccessInventory avec gestion des armes et équipements
func AccessInventory(player *Character) {
	for {
		ui.PrintInfo(fmt.Sprintf("\n=== Inventaire (%d/%d) ===", len(player.Inventory), player.InventorySize))
		if len(player.Inventory) == 0 {
			ui.PrintInfo("Votre sac est vide.")
			return
		}

		// ✅ Affichage amélioré avec types d'objets
		for i, item := range player.Inventory {
			itemDescription := getItemDescription(item)
			ui.PrintInfo(fmt.Sprintf("%d. %s%s", i+1, item, itemDescription))
		}
		
		ui.PrintInfo("A. Équiper une arme")
		ui.PrintInfo("E. Gérer l'équipement")
		ui.PrintInfo("W. Voir les armes disponibles")
		ui.PrintInfo("0. Retour")

		var choice string
		ui.PrintInfo("👉 Votre choix : ")
		fmt.Scanln(&choice)

		switch strings.ToUpper(choice) {
		case "0":
			return
		case "A":
			handleWeaponSelection(player)
		case "E":
			handleEquipmentManagement(player)
		case "W":
			displayAvailableWeapons(player)
		default:
			// Essayer d'utiliser un objet par numéro
			var itemIndex int
			if _, err := fmt.Sscanf(choice, "%d", &itemIndex); err == nil {
				if itemIndex >= 1 && itemIndex <= len(player.Inventory) {
					item := player.Inventory[itemIndex-1]
					handleItemUsage(player, item, itemIndex-1)
				} else {
					ui.PrintError("❌ Numéro invalide")
				}
			} else {
				ui.PrintError("❌ Choix invalide")
			}
		}
	}
}

// ✅ NOUVEAU : Description détaillée des objets
func getItemDescription(item string) string {
	// Vérifier si c'est une arme
	if weapon, isWeapon := Weapons[item]; isWeapon {
		return fmt.Sprintf(" ⚔️ (Arme: +%d dégâts, %s)", weapon.Damage, weapon.Type)
	}
	
	// Vérifier si c'est un équipement
	if isEquipmentItem(item) {
		return " 🛡️ (Équipement)"
	}
	
	// Vérifier si c'est une potion
	switch item {
	case "Potion de vie":
		return " 🧪 (+30 PV)"
	case "Potion de mana":
		return " 🔮 (+20 Mana)"
	case "Potion de poison":
		return " ☠️ (Dangereux !)"
	}
	
	// Objets spéciaux
	switch item {
	case "Pierre de Résurrection":
		return " 💎 (Résurrection automatique)"
	case "Amulette de Protection":
		return " 🛡️ (Protection contre pièges)"
	case "Cristal de Mana":
		return " 💠 (+50 Mana)"
	case "Essence spirituelle":
		return " ✨ (Matériau rare)"
	case "Fragment d'âme":
		return " 🌟 (Matériau mystique)"
	}
	
	return " 📦 (Objet)"
}

// ✅ NOUVEAU : Gestion de sélection d'arme
func handleWeaponSelection(player *Character) {
	weapons := getAvailableWeapons(player)
	
	if len(weapons) == 0 {
		ui.PrintError("❌ Aucune arme dans votre inventaire !")
		ui.PressEnterToContinue(player)
		return
	}
	
	ui.PrintInfo("\n--- Armes disponibles ---")
	for i, weaponName := range weapons {
		weapon := Weapons[weaponName]
		currentMark := ""
		if player.Weapon.Name == weaponName {
			currentMark = " ✅ (Équipée)"
		}
		ui.PrintInfo(fmt.Sprintf("%d. %s (+%d dégâts, %s)%s", i+1, weaponName, weapon.Damage, weapon.Type, currentMark))
	}
	
	var choice int
	ui.PrintInfo("👉 Quelle arme équiper (0 pour retour) ? ")
	fmt.Scanln(&choice)
	
	if choice == 0 {
		return
	}
	
	if choice >= 1 && choice <= len(weapons) {
		selectedWeapon := weapons[choice-1]
		if player.Weapon.Name == selectedWeapon {
			ui.PrintInfo("Cette arme est déjà équipée !")
		} else {
			player.EquipWeapon(selectedWeapon)
		}
	} else {
		ui.PrintError("❌ Choix invalide")
	}
	
	ui.PressEnterToContinue(player)
}

// ✅ NOUVEAU : Gestion de l'équipement
func handleEquipmentManagement(player *Character) {
	ui.PrintInfo("\n--- Équipement Actuel ---")
	player.DisplayEquipment()
	
	ui.PrintInfo("\n--- Actions ---")
	ui.PrintInfo("1. Équiper un objet")
	ui.PrintInfo("2. Déséquiper un objet")
	ui.PrintInfo("0. Retour")
	
	var choice int
	ui.PrintInfo("👉 Votre choix : ")
	fmt.Scanln(&choice)
	
	switch choice {
	case 1:
		handleEquipItem(player)
	case 2:
		handleUnequipItem(player)
	case 0:
		return
	default:
		ui.PrintError("❌ Choix invalide")
	}
	
	ui.PressEnterToContinue(player)
}

// ✅ NOUVEAU : Équiper un objet depuis l'inventaire
func handleEquipItem(player *Character) {
	equipments := getAvailableEquipments(player)
	
	if len(equipments) == 0 {
		ui.PrintError("❌ Aucun équipement dans votre inventaire !")
		return
	}
	
	ui.PrintInfo("\n--- Équipements disponibles ---")
	for i, item := range equipments {
		ui.PrintInfo(fmt.Sprintf("%d. %s", i+1, item))
	}
	
	var choice int
	ui.PrintInfo("👉 Quel équipement ? ")
	fmt.Scanln(&choice)
	
	if choice >= 1 && choice <= len(equipments) {
		selectedEquipment := equipments[choice-1]
		player.EquipItem(selectedEquipment)
	} else {
		ui.PrintError("❌ Choix invalide")
	}
}

// ✅ NOUVEAU : Déséquiper un objet
func handleUnequipItem(player *Character) {
	ui.PrintInfo("\n--- Que déséquiper ? ---")
	ui.PrintInfo("1. Tête")
	ui.PrintInfo("2. Torse")
	ui.PrintInfo("3. Pieds")
	ui.PrintInfo("4. Arme")
	
	var choice int
	ui.PrintInfo("👉 Votre choix : ")
	fmt.Scanln(&choice)
	
	switch choice {
	case 1:
		player.UnequipItem("Head")
	case 2:
		player.UnequipItem("Chest")
	case 3:
		player.UnequipItem("Feet")
	case 4:
		if player.Weapon.Name != "" {
			if player.AddToInventory(player.Weapon.Name) {
				ui.PrintSuccess(fmt.Sprintf("⚔️ %s déséquipée", player.Weapon.Name))
				player.Weapon = Weapon{} // Arme vide
			} else {
				ui.PrintError("🎒 Inventaire plein !")
			}
		} else {
			ui.PrintError("❌ Aucune arme équipée !")
		}
	default:
		ui.PrintError("❌ Choix invalide")
	}
}

// ✅ NOUVEAU : Obtenir les armes disponibles dans l'inventaire
func getAvailableWeapons(player *Character) []string {
	var weapons []string
	for _, item := range player.Inventory {
		if _, isWeapon := Weapons[item]; isWeapon {
			weapons = append(weapons, item)
		}
	}
	return weapons
}

// ✅ NOUVEAU : Afficher les armes disponibles
func displayAvailableWeapons(player *Character) {
	weapons := getAvailableWeapons(player)
	
	if len(weapons) == 0 {
		ui.PrintError("❌ Aucune arme dans votre inventaire !")
		ui.PressEnterToContinue(player)
		return
	}
	
	ui.PrintInfo("\n--- Armes dans l'inventaire ---")
	for _, weaponName := range weapons {
		weapon := Weapons[weaponName]
		equipped := ""
		if player.Weapon.Name == weaponName {
			equipped = " ✅ (Équipée)"
		}
		ui.PrintInfo(fmt.Sprintf("⚔️ %s (+%d dégâts, %s)%s", weaponName, weapon.Damage, weapon.Type, equipped))
	}
	
	ui.PressEnterToContinue(player)
}

// ✅ NOUVEAU : Obtenir les équipements disponibles
func getAvailableEquipments(player *Character) []string {
	var equipments []string
	for _, item := range player.Inventory {
		if isEquipmentItem(item) {
			equipments = append(equipments, item)
		}
	}
	return equipments
}

// ✅ NOUVEAU : Vérifier si un objet est un équipement
func isEquipmentItem(item string) bool {
	equipmentItems := []string{
		"Chapeau de l'Errant", "Chapeau de l'aventurier",
		"Tunique des Songes", "Tunique de l'aventurier", 
		"Bottes de l'Oublié", "Bottes de l'aventurier",
		"Amulette de Protection",
	}
	
	for _, equipment := range equipmentItems {
		if item == equipment {
			return true
		}
	}
	return false
}

// ✅ NOUVEAU : Gestion de l'usage d'objets
func handleItemUsage(player *Character, item string, itemIndex int) {
	switch item {
	case "Potion de vie":
		if player.PvCurr >= player.PvMax {
			ui.PrintError("💖 Vos PV sont déjà au maximum !")
			return
		}
		player.RemoveFromInventory(item)
		player.PvCurr += 30
		if player.PvCurr > player.PvMax {
			player.PvCurr = player.PvMax
		}
		ui.PrintSuccess(fmt.Sprintf("💖 +30 PV ! (%d/%d)", player.PvCurr, player.PvMax))
		
	case "Potion de mana":
		if player.ManaCurr >= player.ManaMax {
			ui.PrintError("🔮 Votre mana est déjà au maximum !")
			return
		}
		player.RemoveFromInventory(item)
		player.ManaCurr += 20
		if player.ManaCurr > player.ManaMax {
			player.ManaCurr = player.ManaMax
		}
		ui.PrintSuccess(fmt.Sprintf("🔮 +20 Mana ! (%d/%d)", player.ManaCurr, player.ManaMax))
		
	case "Potion de poison":
		ui.PrintError("☠️ Vous êtes sûr de vouloir boire cela ? (o/n)")
		var confirm string
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) == "o" || strings.ToLower(confirm) == "oui" {
			player.TakePoison()
		} else {
			ui.PrintInfo("Vous rangez prudemment la potion.")
		}
		
	case "Cristal de Mana":
		player.RemoveFromInventory(item)
		player.ManaCurr += 50
		if player.ManaCurr > player.ManaMax {
			player.ManaCurr = player.ManaMax
		}
		ui.PrintSuccess(fmt.Sprintf("💠 +50 Mana ! (%d/%d)", player.ManaCurr, player.ManaMax))
		
	case "Pierre de Résurrection":
		ui.PrintInfo("💎 Cette pierre pulse d'une énergie mystique...")
		ui.PrintInfo("Elle vous protégera automatiquement en cas de mort.")
		player.HasResurrectionStone = true
		
	default:
		// Vérifier si c'est une arme
		if _, isWeapon := Weapons[item]; isWeapon {
			ui.PrintInfo(fmt.Sprintf("⚔️ %s est une arme. Utilisez le menu 'A' pour l'équiper.", item))
		} else if isEquipmentItem(item) {
			ui.PrintInfo(fmt.Sprintf("🛡️ %s est un équipement. Utilisez le menu 'E' pour l'équiper.", item))
		} else {
			ui.PrintError("❌ Cet objet ne peut pas être utilisé directement.")
		}
	}
}


