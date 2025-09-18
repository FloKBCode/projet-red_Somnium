package shop

import (
	"fmt"
	"somnium/character"
	"somnium/ui"
	"strings"
)

var MerchantItems = map[string]int{
	"Potion de vie":                   3,
	"Potion de mana":                  2,   
	"Potion de poison":                6,
	"Livre de Sort: Boule de feu":     25,
	"Livre de Sort: Chaîne d'éclairs": 40,  
	"Livre de Sort: Mur de glace":     35,  
	"Livre de Sort: Soin majeur":      50,  
	"Livre de Sort: Draine-âme":       30,  
	"Livre de Sort: Explosion psychique": 75, 
	"Fourrure de Loup":                4,
	"Peau de Troll":                   7,
	"Cuir de Sanglier":                3,
	"Plume de Corbeau":                1,
	"Augmentation d'inventaire":       30,
}

func MerchantMenu(player *character.Character) {
	for {
		ui.PrintInfo("\n🏪 === Marchand des Fragments ===")
		ui.PrintInfo(fmt.Sprintf("💰 Vos fragments : %d", player.Money))
		
		i := 1
		itemList := make([]string, 0, len(MerchantItems))
		
		for item, price := range MerchantItems {
			fmt.Printf("%d. %s - %d fragments\n", i, item, price)
			itemList = append(itemList, item)
			i++
		}
		fmt.Println("0. Quitter")
		
		var choice int
		fmt.Print("👉 Votre choix : ")
		fmt.Scanln(&choice)
		
		if choice == 0 {
			ui.PrintInfo("Vous quittez le marchand...")
			return
		}
		
		if choice < 1 || choice > len(itemList) {
			ui.PrintError("❌ Choix invalide")
			continue
		}
		
		selectedItem := itemList[choice-1]
		price := MerchantItems[selectedItem]
		
		if player.Money < price {
			ui.PrintError(fmt.Sprintf("💰 Pas assez de fragments ! Il vous faut %d.", price))
			continue
		}
		
		// Gestion spéciale pour les livres de sorts
		if strings.HasPrefix(selectedItem, "Livre de Sort:") {
			spellName := strings.TrimPrefix(selectedItem, "Livre de Sort: ")
			if player.CanCastSpell(spellName) {
				ui.PrintError(fmt.Sprintf("📖 Vous connaissez déjà le sort %s !", spellName))
				continue
			}
			
			player.Money -= price
			player.LearnSpell(spellName)
			ui.PrintSuccess(fmt.Sprintf("📚✨ Vous apprenez le sort %s !", spellName))
			continue
		}
		
		// Gestion spéciale pour l'augmentation d'inventaire
		if selectedItem == "Augmentation d'inventaire" {
			if player.UpgradeInventorySlot() {
				player.Money -= price
			}
			continue
		}
		
		// Objets normaux
		if !player.AddToInventory(selectedItem) {
			ui.PrintError("🎒 Inventaire plein !")
			continue
		}
		
		player.Money -= price
		ui.PrintSuccess(fmt.Sprintf("✅ Vous achetez %s pour %d fragments !", selectedItem, price))
	}
}

