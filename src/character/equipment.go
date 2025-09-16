package character

import "somnium/ui"


func (c *Character) EquipItem(itemName string) bool {
	switch itemName {
	case "Chapeau de l'Errant", "Chapeau de l'aventurier":
		if c.Equipment.Head == "" {
			c.Equipment.Head = itemName
			c.PvMax += 10
			ui.PrintSuccess(itemName + " équipé (+10 PV).")
			return true
		}
	case "Tunique des Songes", "Tunique de l'aventurier":
		if c.Equipment.Chest == "" {
			c.Equipment.Chest = itemName
			c.PvMax += 25
			ui.PrintSuccess(itemName + " équipé (+25 PV).")
			return true
		}
	case "Bottes de l’Oublié", "Bottes de l'aventurier":
		if c.Equipment.Feet == "" {
			c.Equipment.Feet = itemName
			c.PvMax += 15
			ui.PrintSuccess(itemName + " équipé (+15 PV).")
			return true
		}
	}
	ui.PrintError("❌ Slot déjà occupé ou item invalide.")
	return false
}

func (c *Character) UnequipItem(slot string) bool {
	switch slot {
	case "Head":
		if c.Equipment.Head != "" {
			ui.PrintInfo("⛑ Retrait de " + c.Equipment.Head)
			c.PvMax -= 10
			c.Equipment.Head = ""
			return true
		}
	case "Chest":
		if c.Equipment.Chest != "" {
			ui.PrintInfo("👕 Retrait de " + c.Equipment.Chest)
			c.PvMax -= 25
			c.Equipment.Chest = ""
			return true
		}
	case "Feet":
		if c.Equipment.Feet != "" {
			ui.PrintInfo("👢 Retrait de " + c.Equipment.Feet)
			c.PvMax -= 15
			c.Equipment.Feet = ""
			return true
		}
	}
	ui.PrintError("❌ Aucun item à déséquiper dans ce slot.")
	return false
}

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

func (c *Character) UpdateStatsFromEquipment() {
	currentBonus := c.GetEquipmentBonus()
	if c.PvCurr > c.PvMax-currentBonus {
		c.PvCurr = c.PvMax - currentBonus
	}
}

func (c *Character) DisplayEquipment() {
	ui.PrintInfo("\n=== Équipement ===")
	ui.PrintInfo("Tête: " + c.Equipment.Head)
	ui.PrintInfo("Torse: " + c.Equipment.Chest)
	ui.PrintInfo("Pieds: " + c.Equipment.Feet)
}
