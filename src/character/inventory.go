package character

import (
	"fmt"
	"somnium/ui"
)

// AddToInventory ajoute un objet si de la place est dispo
func (c *Character) AddToInventory(item string) bool {
	if len(c.Inventory) >= c.InventorySize {
		ui.PrintError("🎒 Inventaire plein !")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	ui.PrintSuccess("➕ " + item + " ajouté à l’inventaire.")
	return true
}

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

// CountItem retourne combien d’objets du même type sont présents
func (c *Character) CountItem(itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

// TakePot consomme une potion de vie
func (c *Character) TakePot() bool {
	if !c.RemoveFromInventory("Potion de vie") {
		ui.PrintError("❌ Pas de potion de vie !")
		return false
	}
	c.PvCurr += 50
	if c.PvCurr > c.PvMax {
		c.PvCurr = c.PvMax
	}
	ui.PrintSuccess(fmt.Sprintf("💖 Vous récupérez 50 PV ! (%d/%d)", c.PvCurr, c.PvMax))
	return true
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

// UpgradeInventorySlot agrandit le sac (3 fois max)
func (c *Character) UpgradeInventorySlot() bool {
	upgradeCost := 30
	maxUpgrades := 3

	if c.XPUpgrades >= maxUpgrades {
		ui.PrintError("🚫 Votre sac ne peut pas être agrandi davantage.")
		return false
	}
	if c.Money < upgradeCost {
		ui.PrintError("💰 Pas assez d'or pour améliorer votre sac.")
		return false
	}

	c.Money -= upgradeCost
	c.InventorySize += 10
	c.XPUpgrades++
	ui.PrintSuccess(fmt.Sprintf("🎒 Votre sac s’élargit (+10 emplacements). Capacité : %d", c.InventorySize))
	return true
}

// AccessInventory permet d’utiliser un objet depuis le sac
func AccessInventory(player *Character) {
	for {
		fmt.Printf("\n=== Inventaire (%d/%d) ===\n", len(player.Inventory), player.InventorySize)
		if len(player.Inventory) == 0 {
			fmt.Println("Votre sac est vide.")
			return
		}

		for i, item := range player.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Println("0. Retour")

		var choice int
		fmt.Print("Utiliser quel objet ? ")
		fmt.Scanln(&choice)

		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(player.Inventory) {
			ui.PrintError("❌ Choix invalide")
			continue
		}

		item := player.Inventory[choice-1]
		switch item {
		case "Potion de vie":
			player.TakePot()
		case "Potion de poison":
			player.TakePoison()
		default:
			player.UseItem(item)
		}
	}
}
