package shop

import (
	"fmt"
	"somnium/character"
)

type Item struct {
	Name   string
	Price  int
	Effect func(*character.Character)
}

var inventories = map[int][]Item{
	1: {
		{"Potion de Rêve", 10, func(c *character.Character) {
			c.PvCurr += 20
			if c.PvCurr > c.PvMax {
				c.PvCurr = c.PvMax
			}
		}},
		{"Amulette du Souvenir", 15, func(c *character.Character) {
			c.ManaCurr += 10
			if c.ManaCurr > c.ManaMax {
				c.ManaCurr = c.ManaMax
			}
		}},
	},
	2: {
		{"Clé des Couches", 25, func(c *character.Character) { c.Level++ }},
		{"Potion de Clarté", 20, func(c *character.Character) {
			c.PvCurr += 30
			if c.PvCurr > c.PvMax {
				c.PvCurr = c.PvMax
			}
			c.ManaCurr += 10
			if c.ManaCurr > c.ManaMax {
				c.ManaCurr = c.ManaMax
			}
		}},
	},
	3: {
		{"Pierre de l'Esprit", 50, func(c *character.Character) {
			c.PvCurr = c.PvMax
			c.ManaCurr = c.ManaMax
		}},
	},
}

func MerchantMenu(player *character.Character) {
	inventory, exists := inventories[player.Level]
	if !exists || len(inventory) == 0 {
		fmt.Println("Le marchand n'a rien à vendre pour cette couche...")
		return
	}
	displayMerchantItems(player, inventory)
	var choice int
	fmt.Print("Choisissez un objet : ")
	fmt.Scanln(&choice)
	if choice == 0 {
		fmt.Println("Vous quittez le marchand...")
		return
	}
	if choice < 1 || choice > len(inventory) {
		fmt.Println("Objet invalide.")
		return
	}
	if processPurchase(player, choice-1) {
		fmt.Printf("Vous avez acheté %s !\n", inventory[choice-1].Name)
	}
}

func displayMerchantItems(player *character.Character, inventory []Item) {
	fmt.Println("=== Marchand des Cauchemars ===")
	fmt.Printf("Couches explorées: %d | Fragments de mémoire: %d\n", player.Level, player.Money)
	for i, item := range inventory {
		fmt.Printf("%d) %s - %d fragments\n", i+1, item.Name, item.Price)
	}
	fmt.Println("0) Quitter")
}

func processPurchase(player *character.Character, itemChoice int) bool {
	inventory := inventories[player.Level]
	selected := inventory[itemChoice]
	if !canAfford(player, selected.Price) {
		fmt.Println("Vous n'avez pas assez de fragments...")
		return false
	}
	player.Money -= selected.Price
	selected.Effect(player)
	return true
}

func canAfford(player *character.Character, price int) bool {
	return player.Money >= price
}
