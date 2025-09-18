package character

import (
	"fmt"
	"somnium/ui"
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
			if player.PvCurr >= player.PvMax {
				ui.PrintError("💖 Vos PV sont déjà au maximum !")
				continue
			}
			ui.PrintSuccess("💖 Vos PV augmentent !")
		case "Potion de mana":
			if player.ManaCurr >= player.ManaMax {
				ui.PrintError("🔮 Votre mana est déjà au maximum !")
				continue
			}
			ui.PrintSuccess("🔮 Votre mana augmente !")
		}

		player.UseItem(item)
	}
}
