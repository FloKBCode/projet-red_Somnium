package game

import (
	"errors"
	"fmt"
	"somnium/character"
	"somnium/combat"
	"somnium/ui"
	"time"
	"strings"
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

// ✅ NOUVEAU : Types de fins possibles
type EndingType int

const (
	EndingGood EndingType = iota  // Réveil du coma - libération
	EndingBad                     // Mort définitive - succombé aux démons
	EndingNeutral                 // Retour forcé au début
)

// ✅ NOUVEAU : Structure pour une fin
type DungeonEnding struct {
	Type        EndingType
	Title       string
	Description []string
	Condition   string
}

// ✅ NOUVEAU : Fins disponibles
var PossibleEndings = []DungeonEnding{
	{
		Type:  EndingGood,
		Title: "🌟 L'ÉVEIL LIBÉRATEUR 🌟",
		Description: []string{
			"Vos yeux s'ouvrent lentement... La lumière du jour filtre à travers les rideaux.",
			"Vous êtes dans un lit d'hôpital, mais cette fois, vous êtes vraiment éveillé.",
			"Les médecins parlent de 'miracle'... Vous savez que c'est bien plus que cela.",
			"Vous avez affronté vos démons les plus profonds et vous en êtes sorti vainqueur.",
			"Le trauma qui vous emprisonnait s'est dissous dans la lumière de votre courage.",
			"Votre esprit est libre. Votre nouvelle vie commence maintenant.",
		},
		Condition: "Vaincre le boss final avec plus de 50% de PV",
	},
	{
		Type:  EndingBad,
		Title: "💀 L'ABSORPTION ÉTERNELLE 💀",
		Description: []string{
			"L'obscurité vous engloutit complètement...",
			"Vos démons intérieurs ont pris le contrôle de votre esprit.",
			"Dans le monde réel, les médecins constatent que votre état s'aggrave.",
			"Votre corps respire encore, mais votre âme s'est perdue dans les ténèbres.",
			"Le trauma a gagné. Vous devenez une partie du Labyrinthe pour l'éternité.",
			"Votre conscience sombre à jamais dans les profondeurs de votre psyché brisée.",
		},
		Condition: "Mourir trop souvent ou échouer au boss final",
	},
	{
		Type:  EndingNeutral,
		Title: "🔄 LE CYCLE RECOMMENCE 🔄",
		Description: []string{
			"Vous sentez votre esprit vaciller et retourner vers la surface...",
			"Les souvenirs de vos récentes aventures s'estompent comme un rêve.",
			"Vous n'êtes pas encore prêt à affronter la vérité finale.",
			"Votre subconscient vous ramène au point de départ, plus fort qu'avant.",
			"Chaque tentative vous rapproche de la libération finale.",
			"Le Labyrinthe vous attend, patient, pour votre prochain voyage intérieur.",
		},
		Condition: "Fuir du boss final ou abandonner",
	},
}

// ExploreLayer gère l'exploration d'une couche avec fins multiples
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

	// ✅ NARRATIVE IMMERSIVE PAR COUCHE
	displayLayerNarrative(layer, player)

	// ✅ CHOIX D'EXPLORATION
	ui.PrintInfo("\nComment voulez-vous explorer cette couche ?")
	ui.PrintInfo("1. Explorer une salle aléatoire (Système de salles avancé)")
	ui.PrintInfo("2. Faire un choix de progression (Système classique)")

	var explorationChoice int
	ui.PrintInfo("👉 Votre choix (1-2): ")
	fmt.Scanln(&explorationChoice)

	if explorationChoice == 1 {
		// Utiliser le nouveau système de salles
		room := GenerateRoom(player.CurrentLayer, player)
		ui.PrintInfo(fmt.Sprintf("\n🚪 Vous pénétrez dans : %s", room.Name))
		return ExploreRoom(room, player)
	}

	// Système classique avec fins multiples
	ui.PrintInfo(fmt.Sprintf("\n🌀 === %s ===", layer.Name))
	ui.PrintInfo(layer.Description)
	ui.PrintInfo(fmt.Sprintf("Couche actuelle : %d/%d", player.CurrentLayer, MaxLayer))

	ui.PrintInfo(fmt.Sprintf("\n1. %s", layer.Choice1.Text))
	ui.PrintInfo(fmt.Sprintf("2. %s", layer.Choice2.Text))

	choice := InvalidChoice
	for choice == InvalidChoice {
		ui.PrintInfo("\nVotre choix (1-2): ")
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

	// ✅ GESTION SPÉCIALE DE LA COUCHE 4 (Boss final)
	if layer.IsBoss {
		return handleFinalBossLayer(player, selectedChoice)
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
			
			// ✅ Vérifier succès explorateur
			CheckAndUnlockAchievement(player, "explorer")
		}
	}

	return nil
}

// ✅ NOUVEAU : Gestion de la couche finale avec fins multiples
func handleFinalBossLayer(player *character.Character, choice LayerChoice) error {
	if choice.NextLayer == 0 { // Choix de fuite
		ui.PrintError("💀 Vous choisissez de fuir face à votre trauma...")
		time.Sleep(2 * time.Second)
		return triggerEnding(player, EndingNeutral)
	}

	// Combat final contre le boss
	ui.PrintError("💀 Vous sentez une présence terrifiante...")
	ui.PrintError("Le Trauma Primordial émerge des ténèbres...")
	
	boss := combat.GenerateBoss(player.Level)
	
	// Sauvegarder l'état avant le combat final
	
	victory := combat.StartBossFight(player, boss)

	if !victory {
		// Échec - fin mauvaise
		ui.PrintError("💀 Vos forces vous abandonnent face à vos démons les plus profonds...")
		time.Sleep(2 * time.Second)
		return triggerEnding(player, EndingBad)
	}

	// Victoire - déterminer le type de fin selon l'état du joueur
	if player.PvCurr > player.PvMax/2 {
		// Victoire avec plus de 50% PV - fin parfaite
		ui.PrintSuccess("🌟 Le trauma se dissout dans la lumière de votre courage !")
		time.Sleep(2 * time.Second)
		return triggerEnding(player, EndingGood)
	} else {
		// Victoire difficile - fin neutre (cycle recommence)
		ui.PrintInfo("🔄 Vous avez survécu, mais à quel prix...")
		time.Sleep(2 * time.Second)
		return triggerEnding(player, EndingNeutral)
	}
}

// ✅ NOUVEAU : Déclencher une fin spécifique
func triggerEnding(player *character.Character, endingType EndingType) error {
	// Trouver la fin correspondante
	var selectedEnding *DungeonEnding
	for i := range PossibleEndings {
		if PossibleEndings[i].Type == endingType {
			selectedEnding = &PossibleEndings[i]
			break
		}
	}

	if selectedEnding == nil {
		return fmt.Errorf("fin introuvable")
	}

	// Afficher la cinématique de fin
	displayEnding(*selectedEnding, player)

	// Appliquer les conséquences selon le type de fin
	switch endingType {
	case EndingGood:
		handleGoodEnding(player)
	case EndingBad:
		handleBadEnding(player)
	case EndingNeutral:
		handleNeutralEnding(player)
	}

	return nil
}

// ✅ NOUVEAU : Affichage cinématique de la fin
func displayEnding(ending DungeonEnding, player *character.Character) {
	ui.ClearScreen(player)
	
	ui.PrintError("\n" + strings.Repeat("═", 60))
	ui.PrintError(ending.Title)
	ui.PrintError(strings.Repeat("═", 60))
	
	for _, line := range ending.Description {
		time.Sleep(2500 * time.Millisecond)
		
		switch ending.Type {
		case EndingGood:
			ui.PrintSuccess(line)
		case EndingBad:
			ui.PrintError(line)
		case EndingNeutral:
			ui.PrintInfo(line)
		}
	}
	
	time.Sleep(3 * time.Second)
	ui.PrintError(strings.Repeat("═", 60))
	
	// Afficher les statistiques finales
	displayFinalStats(player)
	
	time.Sleep(2 * time.Second)
}

// ✅ NOUVEAU : Affichage des statistiques finales
func displayFinalStats(player *character.Character) {
	ui.PrintInfo("\n📊 === BILAN DE VOTRE VOYAGE ===")
	ui.PrintInfo(fmt.Sprintf("Nom : %s (%s %s)", player.Name, player.Race, player.Class))
	ui.PrintInfo(fmt.Sprintf("Niveau atteint : %d", player.Level))
	ui.PrintInfo(fmt.Sprintf("Couche la plus profonde : %d/%d", player.CurrentLayer, MaxLayer))
	ui.PrintInfo(fmt.Sprintf("Fragments collectés : %d", player.Money))
	ui.PrintInfo(fmt.Sprintf("Sorts appris : %d", len(player.Skills)))
	ui.PrintInfo(fmt.Sprintf("Succès débloqués : %d", len(player.Achievements)))
	
	if player.Weapon.Name != "" {
		ui.PrintInfo(fmt.Sprintf("Arme finale : %s", player.Weapon.Name))
	}
}

// ✅ NOUVEAU : Conséquences de la bonne fin
func handleGoodEnding(player *character.Character) {
	// Le joueur a vaincu ses démons - succès ultime
	CheckAndUnlockAchievement(player, "boss_slayer")
	
	// Récompenses spéciales pour la fin parfaite
	player.GainXP(1000)
	player.Money += 500
	player.AddToInventory("Couronne de la Libération")
	
	ui.PrintSuccess("\n🎉 FÉLICITATIONS ! Vous avez atteint la fin parfaite !")
	ui.PrintSuccess("🏆 Votre courage a payé - vous êtes libre !")
	
	ui.PressEnterToContinue(player)
	
	// Proposer de recommencer en New Game+
	ui.PrintInfo("Voulez-vous recommencer avec vos acquis ? (New Game+) (o/n)")
	var choice string
	fmt.Scanln(&choice)
	
	if choice == "o" || choice == "oui" {
		startNewGamePlus(player)
	} else {
		ui.PrintInfo("Merci d'avoir joué au Labyrinthe des Cauchemars !")
	}
}

// ✅ NOUVEAU : Conséquences de la mauvaise fin
func handleBadEnding(player *character.Character) {
	ui.PrintError("\n💀 Votre esprit a été consumé par les ténèbres...")
	ui.PrintError("😢 Cette fin tragique marque la victoire de vos démons intérieurs.")
	
	ui.PressEnterToContinue(player)
	
	// Proposer de recommencer
	ui.PrintInfo("Voulez-vous recommencer votre voyage ? (o/n)")
	var choice string
	fmt.Scanln(&choice)
	
	if choice == "o" || choice == "oui" {
		resetPlayerForRestart(player)
		ui.PrintInfo("🔄 Une nouvelle chance vous est offerte...")
	} else {
		ui.PrintInfo("Parfois, accepter l'échec fait partie du voyage...")
	}
}

// ✅ NOUVEAU : Conséquences de la fin neutre
func handleNeutralEnding(player *character.Character) {
	ui.PrintInfo("\n🔄 Votre voyage n'est pas terminé...")
	ui.PrintInfo("💪 Chaque tentative vous rend plus fort.")
	
	// Bonus pour la prochaine tentative
	player.PvMax += 20
	player.ManaMax += 10
	player.PvCurr = player.PvMax
	player.ManaCurr = player.ManaMax
	
	// Retour à la couche 1 avec bonus
	player.CurrentLayer = 1
	player.Money += 100
	
	ui.PrintSuccess("✨ Bonus pour votre prochain essai : +20 PV Max, +10 Mana Max, +100 fragments !")
	
	ui.PressEnterToContinue(player)
}

// ✅ NOUVEAU : New Game+ avec avantages
func startNewGamePlus(player *character.Character) {
	ui.PrintSuccess("🌟 === MODE NEW GAME+ ACTIVÉ ===")
	
	// Conserver certains avantages
	bonusHP := player.Level * 10
	bonusMana := player.Level * 5
	bonusMoney := player.Money / 2
	keepWeapon := player.Weapon
	keepAchievements := player.Achievements
	
	// Réinitialiser le joueur mais avec bonus
	*player = character.InitCharacter(player.Name, player.Race, player.Class, player.PvMax+bonusHP, player.ManaMax+bonusMana)
	
	// Restaurer les avantages
	player.Money += bonusMoney
	player.Weapon = keepWeapon
	player.Achievements = keepAchievements
	
	ui.PrintSuccess(fmt.Sprintf("💪 Bonus New Game+ : +%d PV Max, +%d Mana Max, +%d fragments", bonusHP, bonusMana, bonusMoney))
	
	if keepWeapon.Name != "" {
		ui.PrintSuccess(fmt.Sprintf("⚔️ Arme conservée : %s", keepWeapon.Name))
	}
}

// ✅ NOUVEAU : Reset pour recommencer
func resetPlayerForRestart(player *character.Character) {
	// Garder le nom et la race/classe mais reset stats
	name := player.Name
	race := player.Race
	class := player.Class
	
	*player = character.InitCharacter(name, race, class, 100, 50)
}

// GetPlayerLayer retourne la couche actuelle du joueur
func GetPlayerLayer(player *character.Character) int {
	if player.CurrentLayer < 1 || player.CurrentLayer > MaxLayer {
		return 0 // Index pour couche 1
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

// handleCombat gère le combat basé sur le choix du joueur
func handleCombat(player *character.Character, choice LayerChoice) error {
	if err := generateCombatForRisk(player, choice.Risk); err != nil {
		return err
	}
	return dropCraftMaterials(player, choice.Reward)
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

// dropCraftMaterials gère la récupération des matériaux de craft
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
		if player.AddToInventory(material.Name) {
			ui.PrintSuccess(fmt.Sprintf("🎁 Vous trouvez : %s", material.Name))
		}
	}
	return nil
}

// filterMaterialsByLayer filtre les matériaux disponibles selon la couche
func filterMaterialsByLayer(layer int) []Material {
	var filtered []Material
	materials := []Material{
		{"Cuir de Sanglier", 1, 1},
		{"Plume de Corbeau", 1, 1},
		{"Fourrure de Loup", 2, 2},
		{"Peau de Troll", 3, 3},
	}
	
	for _, m := range materials {
		if m.MinLayer <= layer {
			filtered = append(filtered, m)
		}
	}
	return filtered
}

// Type Material
type Material struct {
	Name     string
	Rarity   int
	MinLayer int
}

// displayLayerNarrative affiche la narration spécifique à la couche
func displayLayerNarrative(layer Layer, player *character.Character) {
	ui.PrintInfo(fmt.Sprintf("\n🌀 === %s ===", layer.Name))

	// Récits spécifiques par couche
	switch layer.Level {
	case 1:
		ui.PrintInfo("🌫️ Les brumes de la surface ondulent autour de vous...")
		time.Sleep(1 * time.Second)
		ui.PrintInfo("Ici flottent vos souvenirs les plus récents, encore flous et malléables.")
		time.Sleep(1 * time.Second)
		ui.PrintInfo("Vous entendez l'écho lointain de votre voix consciente qui vous appelle...")

	case 2:
		ui.PrintError("🥀 L'air devient plus lourd, chargé de remords...")
		time.Sleep(1 * time.Second)
		ui.PrintInfo("Dans cette vallée résonnent tous vos 'si seulement' et vos 'j'aurais dû'.")
		time.Sleep(1 * time.Second)
		ui.PrintInfo("Les ombres ici ont la forme de vos choix passés.")

	case 3:
		ui.PrintError("🕳️ Un froid glacial remonte de l'abîme sous vos pieds...")
		time.Sleep(1 * time.Second)
		ui.PrintInfo("Vous êtes maintenant face aux terreurs qui ont façonné votre personnalité.")
		time.Sleep(1 * time.Second)
		ui.PrintError("Chaque pas résonne comme un battement de cœur affolé.")

	case 4:
		ui.PrintError("💀 L'atmosphère devient suffocante, presque tangible...")
		time.Sleep(1 * time.Second)
		ui.PrintError("Vous approchez du noyau de votre souffrance originelle.")
		time.Sleep(1 * time.Second)
		ui.PrintError("Ici, seuls les plus braves peuvent espérer triompher.")
	}

	time.Sleep(1500 * time.Millisecond)
	ui.PrintInfo(layer.Description)
}

// Définition des couches (existantes)
type Layer struct {
	Level       int
	Name        string
	Description string
	Choice1     LayerChoice
	Choice2     LayerChoice
	IsBoss      bool
}

type LayerChoice struct {
	Text       string
	Risk       int
	Reward     int
	NextLayer  int
	FlavorText string
}

var Layers = []Layer{
	{
		Level:       1,
		Name:        "Surface des Rêves",
		Description: "Les premières brumes de votre inconscient se dessinent...",
		Choice1: LayerChoice{
			Text:       "Explorer les souvenirs familiers (Sûr)",
			Risk:       1,
			Reward:     1,
			NextLayer:  1,
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
			NextLayer:  0, // Fin neutre
			FlavorText: "Vous remontez vers la lumière, mais elle s'éteint à jamais...",
		},
		Choice2: LayerChoice{
			Text:       "Affronter le Boss du Trauma (Courage)",
			Risk:       3,
			Reward:     10,
			NextLayer:  5, // Combat final
			FlavorText: "Vous levez la tête. Cette fois, vous êtes assez fort. Le combat final commence.",
		},
		IsBoss: true,
	},
}