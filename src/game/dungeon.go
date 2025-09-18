package game

import (
	"errors"
	"fmt"
	"math/rand"
	"somnium/character"
	"somnium/combat"
	"somnium/ui"
	"time"
	"somnium/quest"
)

const (
	MaxLayer      = 4
	InvalidChoice = -1
)

var (
	ErrInvalidLayer = errors.New("couche invalide")
	ErrNilPlayer    = errors.New("joueur invalide")
	ErrExploration  = errors.New("erreur d'exploration")
)

type Material struct {
	Name     string
	Rarity   int
	MinLayer int
}

var Materials = []Material{
	{"Cuir de Sanglier", 1, 1},
	{"Plume de Corbeau", 1, 1},
	{"Fourrure de Loup", 2, 2},
	{"Peau de Troll", 3, 3},
}

// Layer représente une couche de conscience
type Layer struct {
	Level       int
	Name        string
	Description string
	Choice1     LayerChoice
	Choice2     LayerChoice
	IsBoss      bool
}

// LayerChoice représente un choix dans une couche
type LayerChoice struct {
	Text       string
	Risk       int    // 1 = sûr, 2 = risqué
	Reward     int    // Multiplicateur de récompenses
	NextLayer  int    // Couche suivante
	FlavorText string // Phrase d'ambiance
}

// Définition des couches
var Layers = []Layer{
	{
		Level:       1,
		Name:        "Surface des Rêves",
		Description: "Les premières brumes de votre inconscient se dessinent...",
		Choice1: LayerChoice{
			Text:       "Explorer les souvenirs familiers (Sûr)",
			Risk:       1,
			Reward:     1,
			NextLayer:  1, // Reste au même niveau
			FlavorText: "Vous restez dans la zone de confort, mais vos démons grandissent dans l'ombre...",
		},
		Choice2: LayerChoice{
			Text:       "Plonger vers les émotions enfouies (Risqué)",
			Risk:       2,
			Reward:     3,
			NextLayer:  2,
			FlavorText: "Votre courage illumine les profondeurs. Vous sentez votre esprit se renforcer.",
		},
		IsBoss: false,
	},
	{
		Level:       2,
		Name:        "Vallée des Regrets",
		Description: "Les échos de vos choix passés résonnent dans l'obscurité...",
		Choice1: LayerChoice{
			Text:       "Éviter les souvenirs douloureux (Sûr)",
			Risk:       1,
			Reward:     1,
			NextLayer:  2,
			FlavorText: "Vous détournez le regard, mais les blessures restent ouvertes...",
		},
		Choice2: LayerChoice{
			Text:       "Affronter vos regrets (Risqué)",
			Risk:       2,
			Reward:     3,
			NextLayer:  3,
			FlavorText: "Chaque regret accepté devient une leçon. Votre âme se fortifie.",
		},
		IsBoss: false,
	},
	{
		Level:       3,
		Name:        "Gouffre des Peurs Profondes",
		Description: "Ici résident vos terreurs les plus anciennes, celles qui ont façonné qui vous êtes...",
		Choice1: LayerChoice{
			Text:       "Se réfugier dans le déni (Sûr)",
			Risk:       1,
			Reward:     1,
			NextLayer:  3,
			FlavorText: "Vous fermez les yeux, mais vos peurs se nourrissent de votre faiblesse...",
		},
		Choice2: LayerChoice{
			Text:       "Regarder vos peurs en face (Risqué)",
			Risk:       2,
			Reward:     4,
			NextLayer:  4,
			FlavorText: "En nommant vos peurs, vous leur retirez leur pouvoir. Vous êtes presque prêt.",
		},
		IsBoss: false,
	},
	{
		Level:       4,
		Name:        "Le Cœur du Trauma",
		Description: "Vous voilà face à la source de toute votre souffrance. C'est ici que tout se joue.",
		Choice1: LayerChoice{
			Text:       "Fuir vers la surface (Abandon)",
			Risk:       0,
			Reward:     0,
			NextLayer:  0, // Fin du jeu - échec
			FlavorText: "Vous remontez vers la lumière, mais elle s'éteint à jamais...",
		},
		Choice2: LayerChoice{
			Text:       "Affronter le Boss du Trauma (Courage)",
			Risk:       3,
			Reward:     10,
			NextLayer:  5, // Victoire - éveil
			FlavorText: "Vous levez la tête. Cette fois, vous êtes assez fort. Le combat final commence.",
		},
		IsBoss: true,
	},
}

// ExploreLayer gère l'exploration d'une couche
func ExploreLayer(player *character.Character) error {
	if player == nil {
		return ErrNilPlayer
	}

	currentLayerIndex := GetPlayerLayer(player)
	if currentLayerIndex < 0 || currentLayerIndex >= len(Layers) {
		ui.PrintError("❌ Couche invalide")
		return ErrInvalidLayer
	}

	layer := Layers[currentLayerIndex]

	ui.PrintInfo(fmt.Sprintf("\n🌀 === %s ===", layer.Name))
	ui.PrintInfo(layer.Description)
	ui.PrintInfo(fmt.Sprintf("Couche actuelle : %d/%d", player.CurrentLayer, MaxLayer))

	fmt.Printf("\n1. %s\n", layer.Choice1.Text)
	fmt.Printf("2. %s\n", layer.Choice2.Text)

	choice := InvalidChoice
	for choice == InvalidChoice {
		fmt.Print("\nVotre choix (1-2): ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 2 {
			ui.PrintError("Choix invalide. Veuillez entrer 1 ou 2.")
			choice = InvalidChoice
		}
	}

	selectedChoice := layer.Choice1
	if choice == 2 {
		selectedChoice = layer.Choice2
	}

	ui.PrintInfo(fmt.Sprintf("\n%s", selectedChoice.FlavorText))

	if layer.IsBoss {
		ui.PrintInfo("💀 Vous sentez une présence terrifiante...")
		return handleBossLayer(player)
	}

	// Combat selon le risque
	if selectedChoice.Risk > 0 {
		if err := handleCombat(player, selectedChoice); err != nil {
			ui.PrintError(fmt.Sprintf("Erreur de combat : %v", err))
			return err
		}
	}

	// Progression de couche
	if selectedChoice.NextLayer != player.CurrentLayer {
		setPlayerLayer(player, selectedChoice.NextLayer)
		if selectedChoice.NextLayer > player.CurrentLayer {
			ui.PrintSuccess(fmt.Sprintf("🌟 Vous avez progressé vers la couche %d !", selectedChoice.NextLayer))
		}
	}

	// Mettre à jour les quêtes
	quest.UpdateQuestProgress("explore", layer.Name, 1)

	return nil
}

func handleCombat(player *character.Character, choice LayerChoice) error {
	if err := generateCombatForRisk(player, choice.Risk); err != nil {
		return err
	}
	return dropCraftMaterials(player, choice.Reward)
}

// GetPlayerLayer retourne la couche actuelle du joueur
func GetPlayerLayer(player *character.Character) int {
	if player.CurrentLayer < 1 || player.CurrentLayer > MaxLayer {
		return 1 // Retourne à la première couche si invalide
	}
	return player.CurrentLayer - 1 // Index 0-based pour le tableau Layers
}

// setPlayerLayer définit la couche du joueur
func setPlayerLayer(player *character.Character, layer int) {
	if layer < 0 || layer > MaxLayer {
		return
	}
	player.CurrentLayer = layer
}

// generateCombatForRisk génère un combat basé sur le risque choisi
func generateCombatForRisk(player *character.Character, risk int) error {
	if player == nil {
		return ErrNilPlayer
	}

	difficulty := combat.Normal
	if risk == 2 {
		difficulty = combat.Hard
	}

	monster := combat.GenerateMonster(player.Level, difficulty)
	return combat.StartFight(player, monster)
}

func dropCraftMaterials(player *character.Character, multiplier int) error {
	if player == nil {
		return ErrNilPlayer
	}
	if multiplier <= 0 {
		return nil
	}

	availableMaterials := filterMaterialsByLayer(player.CurrentLayer)
	for i := 0; i < multiplier && i < len(availableMaterials); i++ {
		material := availableMaterials[i]
		player.AddToInventory(material.Name)
		fmt.Printf("🎁 Vous trouvez : %s\n", material.Name)
	}
	return nil
}

func filterMaterialsByLayer(layer int) []Material {
	var filtered []Material
	for _, m := range Materials {
		if m.MinLayer <= layer {
			filtered = append(filtered, m)
		}
	}
	return filtered
}

// handleBossLayer gère la couche boss
func handleBossLayer(player *character.Character) error {
	if player == nil {
		return ErrNilPlayer
	}

	boss := combat.GenerateBoss(player.Level)
	victory := combat.StartBossFight(player, boss)

	if !victory {
		return gameOver(player)
	}

	fmt.Println("\n🌟 Félicitations ! Vous avez vaincu vos démons intérieurs !")
	return giveBossRewards(player)
}

func giveBossRewards(player *character.Character) error {
	// Implémentation des récompenses spéciales
	return nil
}

// gameOver gère la fin du jeu (échec)
func gameOver(player *character.Character) error {
	if player == nil {
		return ErrNilPlayer
	}

	fmt.Println("\n💀 Votre esprit sombre dans les ténèbres...")
	player.CurrentLayer = 1
	return nil
}

func hasExploredEnough(player *character.Character) bool {
	if player == nil {
		return false
	}
	currentLayer := GetPlayerLayer(player)
	if currentLayer < 0 || currentLayer >= len(Layers) {
		return false
	}
	return player.Level >= Layers[currentLayer].Level*2
}

func unlockMerchantItems(player *character.Character) {
	if player == nil {
		return
	}
	// TODO: Implémenter le déblocage d'items selon le niveau de couche
}

// GenerateLoot sélectionne un matériau de loot en fonction de la couche
func GenerateLoot(layer Layer) Material {
	var possible []Material
	for _, m := range Materials {
		if m.MinLayer <= layer.Level {
			possible = append(possible, m)
		}
	}
	if len(possible) == 0 {
		return Material{Name: "Objet inconnu", Rarity: 0, MinLayer: 0}
	}
	rand.Seed(time.Now().UnixNano())
	return possible[rand.Intn(len(possible))]
}
