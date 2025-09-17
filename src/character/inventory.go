package character

import (
	"fmt"
	"somnium/ui"
)

func (c *Character) AddToInventory(item string) bool {
	if len(c.Inventory) >= c.InventorySize {
		fmt.Println("🎒 Inventaire plein !")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func (c *Character) RemoveFromInventory(item string) bool {
	for i, invItem := range c.Inventory {
		if invItem == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Character) CountItem(itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

func (c *Character) TakePot() bool {
	if c.CountItem("Potion de vie") == 0 {
		fmt.Println("❌ Pas de potion de vie !")
		return false
	}
	
	c.RemoveFromInventory("Potion de vie")
	c.PvCurr += 50
	if c.PvCurr > c.PvMax {
		c.PvCurr = c.PvMax
	}
	
	fmt.Printf("💖 Vous récupérez 50 PV ! (%d/%d)\n", c.PvCurr, c.PvMax)
	return true
}

func (c *Character) TakePoison() bool {
	if c.CountItem("Potion de poison") == 0 {
		fmt.Println("❌ Pas de potion de poison !")
		return false
	}
	
	c.RemoveFromInventory("Potion de poison")
	fmt.Println("☠️ Vous buvez la potion de poison...")
	c.PoisonEffect()
	return true
}

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

func AccessInventory(player *Character) {
	for {
		fmt.Println("\n=== Inventaire ===")
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
			fmt.Println("❌ Choix invalide")
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