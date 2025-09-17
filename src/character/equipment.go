package character

import "somnium/ui"

// EquipItem équipe un objet et renvoie true si succès
func (c *Character) EquipItem(itemName string) bool {
	switch itemName {
	case "Chapeau de l'Errant", "Chapeau de l'aventurier":
		// si déjà un chapeau, on le remet dans l'inventaire
		if c.Equipment.Head != "" {
			c.Inventory = append(c.Inventory, c.Equipment.Head)
			ui.PrintInfo("⛑ " + c.Equipment.Head + " a été rangé dans l'inventaire.")
		}
		c.Equipment.Head = itemName
		ui.PrintSuccess(itemName + " équipé (+10 PV).")
		c.UpdateStatsFromEquipment()
		return true

	case "Tunique des Songes", "Tunique de l'aventurier":
		if c.Equipment.Chest != "" {
			c.Inventory = append(c.Inventory, c.Equipment.Chest)
			ui.PrintInfo("👕 " + c.Equipment.Chest + " a été rangé dans l'inventaire.")
		}
		c.Equipment.Chest = itemName
		ui.PrintSuccess(itemName + " équipé (+25 PV).")
		c.UpdateStatsFromEquipment()
		return true

	case "Bottes de l’Oublié", "Bottes de l'aventurier":
		if c.Equipment.Feet != "" {
			c.Inventory = append(c.Inventory, c.Equipment.Feet)
			ui.PrintInfo("👢 " + c.Equipment.Feet + " a été rangé dans l'inventaire.")
		}
		c.Equipment.Feet = itemName
		ui.PrintSuccess(itemName + " équipé (+15 PV).")
		c.UpdateStatsFromEquipment()
		return true
	}

	ui.PrintError("❌ Item invalide ou non reconnu.")
	return false
}

// UnequipItem enlève l’objet du slot et le range dans l’inventaire
func (c *Character) UnequipItem(slot string) bool {
	switch slot {
	case "Head":
		if c.Equipment.Head != "" {
			ui.PrintInfo("⛑ Retrait de " + c.Equipment.Head)
			c.Inventory = append(c.Inventory, c.Equipment.Head)
			c.Equipment.Head = ""
			c.UpdateStatsFromEquipment()
			return true
		}
	case "Chest":
		if c.Equipment.Chest != "" {
			ui.PrintInfo("👕 Retrait de " + c.Equipment.Chest)
			c.Inventory = append(c.Inventory, c.Equipment.Chest)
			c.Equipment.Chest = ""
			c.UpdateStatsFromEquipment()
			return true
		}
	case "Feet":
		if c.Equipment.Feet != "" {
			ui.PrintInfo("👢 Retrait de " + c.Equipment.Feet)
			c.Inventory = append(c.Inventory, c.Equipment.Feet)
			c.Equipment.Feet = ""
			c.UpdateStatsFromEquipment()
			return true
		}
	}
	ui.PrintError("❌ Aucun item à déséquiper dans ce slot.")
	return false
}

// GetEquipmentBonus calcule le bonus total PV selon l’équipement
func (c *Character) GetEquipmentBonus() int {
	bonus := 0
	if c.Equipment.Head != "" {
		bonus += 10
	}
	if c.Equipment.Chest != "" {
		bonus += 25
	}
	if c.Equipment.Feet != "" {
		bonus += 15
	}
	return bonus
}

// UpdateStatsFromEquipment recalcule PvMax et ajuste PvCurr
func (c *Character) UpdateStatsFromEquipment() {
	baseHP := (c.Level-1)*20 + 100 // Exemple : 100 PV de base + 20 par niveau
	// 👉 Tu peux remplacer cette formule par ta logique exacte
	c.PvMax = baseHP + c.GetEquipmentBonus()

	if c.PvCurr > c.PvMax {
		c.PvCurr = c.PvMax
	}
}

// DisplayEquipment affiche l’équipement actuel
func (c *Character) DisplayEquipment() {
	ui.PrintInfo("\n=== Équipement ===")
	if c.Equipment.Head == "" {
		ui.PrintInfo("Tête: (vide)")
	} else {
		ui.PrintInfo("Tête: " + c.Equipment.Head)
	}
	if c.Equipment.Chest == "" {
		ui.PrintInfo("Torse: (vide)")
	} else {
		ui.PrintInfo("Torse: " + c.Equipment.Chest)
	}
	if c.Equipment.Feet == "" {
		ui.PrintInfo("Pieds: (vide)")
	} else {
		ui.PrintInfo("Pieds: " + c.Equipment.Feet)
	}
}