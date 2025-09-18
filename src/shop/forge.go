package shop

import (
	"fmt"
	"somnium/character"
	"somnium/ui"
)

// Recipe définit une recette de forge maudite.
type Recipe struct {
	Name      string
	Materials map[string]int // "matériau": quantité
	Cost      int
	Result    string
}

// Recipes liste des artefacts disponibles à la forge.
var Recipes = []Recipe{
	{
		Name:      "Chapeau de l'Errant",
		Materials: map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1},
		Cost:      5,
		Result:    "Chapeau de l'Errant",
	},
	{
		Name:      "Tunique des Songes",
		Materials: map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1},
		Cost:      5,
		Result:    "Tunique des Songes",
	},
	{
		Name:      "Bottes de l’Oublié",
		Materials: map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1},
		Cost:      5,
		Result:    "Bottes de l’Oublié",
	},
}

// ForgeMenu affiche le menu principal de la forge et gère les interactions
// Permet au joueur de forger des objets avec les matériaux récoltés
func ForgeMenu(player *character.Character) {
	for {
		fmt.Println("\n═══════════════════════════════")
		fmt.Println(" 🔥 La Forge des Cauchemars 🔥 ")
		fmt.Println("═══════════════════════════════")
		fmt.Println("1. Contempler les artefacts forgés dans les flammes oniriques")
		fmt.Println("0. Quitter ce lieu hanté")
		fmt.Print("→ Votre choix: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displayCraftableItems()
			fmt.Print("\n→ Choisissez un artefact à forger (1-3) ou 0 pour reculer dans l’ombre: ")
			var itemChoice int
			fmt.Scanln(&itemChoice)
			if itemChoice >= 1 && itemChoice <= len(Recipes) {
				if craftItem(player, itemChoice-1) {
					ui.PrintSuccess("⚒️  Les flammes s’élèvent... un nouvel artefact prend vie entre vos mains.")
				}
			} else if itemChoice != 0 {
				ui.PrintError("❌ Les ombres ne comprennent pas ce choix...")
			}
		case 0:
			ui.PrintInfo("Vous quittez la forge, laissant derrière vous l’écho des marteaux.")
			return
		default:
			ui.PrintError("❌ Les flammes se moquent de votre hésitation.")
		}
	}
}

// displayCraftableItems affiche la liste des objets pouvant être forgés
// Montre le nom, le coût et les matériaux requis pour chaque recette
func displayCraftableItems() {
	ui.PrintInfo("\nArtefacts que vous pouvez forger :")
	for i, recipe := range Recipes {
		fmt.Printf("%d. %s (💰 %d or | ⚒️ %v)\n", i+1, recipe.Name, recipe.Cost, recipe.Materials)
	}
}

// craftItem tente de forger un objet spécifique
// Vérifie les ressources du joueur et crée l'objet si possible
// Retourne true si le craft réussit, false sinon
func craftItem(player *character.Character, itemChoice int) bool {
	if itemChoice < 0 || itemChoice >= len(Recipes) {
		ui.PrintError("❌ Le choix s’efface dans le néant.")
		return false
	}

	recipe := Recipes[itemChoice]
	if !hasRequiredMaterials(player, recipe) {
		return false
	}

	// Consomme les matériaux et l’or
	for material, qty := range recipe.Materials {
		for i := 0; i < qty; i++ {
			player.RemoveFromInventory(material)
		}
	}
	player.Money -= recipe.Cost

	// Équipe l’artefact forgé
	player.EquipItem(recipe.Result)

	ui.PrintSuccess(fmt.Sprintf("%s a été forgé dans les flammes oniriques !", recipe.Result))
	return true
}

// hasRequiredMaterials vérifie si le joueur possède les ressources nécessaires
// Contrôle l'or et les matériaux requis pour une recette donnée
// Retourne true si le joueur a tout ce qu'il faut, false sinon
func hasRequiredMaterials(player *character.Character, recipe Recipe) bool {
	if player.Money < recipe.Cost {
		ui.PrintError(fmt.Sprintf(ui.ErrNotEnoughMoney, recipe.Cost))
		return false
	}
	for material, qty := range recipe.Materials {
		if player.CountItem(material) < qty {
			ui.PrintError(fmt.Sprintf(ui.ErrNotEnoughMaterials, recipe.Name))
			return false
		}
	}
	return true
}
