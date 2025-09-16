package character

import "somnium/ui"

func (c *Character) UpgradeInventorySlot() bool {
	upgradeCost := 20 // coût en or, tu peux le mettre configurable

	if c.Money < upgradeCost {
		ui.PrintError("💰 Pas assez d'or pour améliorer votre sac.")
		return false
	}

	// On consomme l'or et on augmente la capacité
	c.Money -= upgradeCost
	c.InventorySize += 5 // +5 slots par amélioration, par exemple

	ui.PrintSuccess("🎒 Votre sac s’élargit, comme si les songes pliaient l’espace (+5 emplacements).")
	return true
}