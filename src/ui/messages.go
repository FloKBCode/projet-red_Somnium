package ui

import "github.com/fatih/color"

const (
	ErrNotEnoughMoney     = "💰 Pas assez d'argent! Il vous faut %d pièces d'or."
	ErrInventoryFull      = "🎒 Votre sac déborde déjà des échos des songes."
	ErrNotEnoughMaterials = "⚒️ Les fragments manquent pour façonner %s."
	ErrNotEnoughMana      = "🔮 L’énergie onirique est insuffisante (%d requis)."
	ErrItemNotFound       = "❌ La relique '%s' ne se trouve pas dans votre sac."
)

func PrintError(message string) {
	color.New(color.FgRed).Add(color.Bold).Println(message)
}

func PrintSuccess(message string) {
	color.New(color.FgGreen).Add(color.Bold).Println(message)
}

func PrintInfo(message string) {
	color.New(color.FgCyan).Println(message)
}
