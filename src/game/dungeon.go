package game


import (
	"fmt"
	"strings"
	"somnium/character"
	"somnium/combat"
)

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
	Text        string
	Risk        int    // 1 = sûr, 2 = risqué
	Reward      int    // Multiplicateur de récompenses
	NextLayer   int    // Couche suivante
	FlavorText  string // Phrase d'ambiance
}

// Définition des couches
var Layers = []Layer{
	{
		Level:       1,
		Name:        "Surface des Rêves",
		Description: "Les premières brumes de votre inconscient se dessinent...",
		Choice1: LayerChoice{
			Text:        "Explorer les souvenirs familiers (Sûr)",
			Risk:        1,
			Reward:      1,
			NextLayer:   1, // Reste au même niveau
			FlavorText:  "Vous restez dans la zone de confort, mais vos démons grandissent dans l'ombre...",
		},
		Choice2: LayerChoice{
			Text:        "Plonger vers les émotions enfouies (Risqué)", 
			Risk:        2,
			Reward:      3,
			NextLayer:   2,
			FlavorText:  "Votre courage illumine les profondeurs. Vous sentez votre esprit se renforcer.",
		},
		IsBoss: false,
	},
	{
		Level:       2,
		Name:        "Vallée des Regrets",
		Description: "Les échos de vos choix passés résonnent dans l'obscurité...",
		Choice1: LayerChoice{
			Text:        "Éviter les souvenirs douloureux (Sûr)",
			Risk:        1,
			Reward:      1,
			NextLayer:   2,
			FlavorText:  "Vous détournez le regard, mais les blessures restent ouvertes...",
		},
		Choice2: LayerChoice{
			Text:        "Affronter vos regrets (Risqué)",
			Risk:        2,
			Reward:      3,
			NextLayer:   3,
			FlavorText:  "Chaque regret accepté devient une leçon. Votre âme se fortifie.",
		},
		IsBoss: false,
	},
	{
		Level:       3,
		Name:        "Gouffre des Peurs Profondes",
		Description: "Ici résident vos terreurs les plus anciennes, celles qui ont façonné qui vous êtes...",
		Choice1: LayerChoice{
			Text:        "Se réfugier dans le déni (Sûr)",
			Risk:        1,
			Reward:      1,
			NextLayer:   3,
			FlavorText:  "Vous fermez les yeux, mais vos peurs se nourrissent de votre faiblesse...",
		},
		Choice2: LayerChoice{
			Text:        "Regarder vos peurs en face (Risqué)",
			Risk:        2,
			Reward:      4,
			NextLayer:   4,
			FlavorText:  "En nommant vos peurs, vous leur retirez leur pouvoir. Vous êtes presque prêt.",
		},
		IsBoss: false,
	},
	{
		Level:       4,
		Name:        "Le Cœur du Trauma",
		Description: "Vous voilà face à la source de toute votre souffrance. C'est ici que tout se joue.",
		Choice1: LayerChoice{
			Text:        "Fuir vers la surface (Abandon)",
			Risk:        0,
			Reward:      0,
			NextLayer:   0, // Fin du jeu - échec
			FlavorText:  "Vous remontez vers la lumière, mais elle s'éteint à jamais...",
		},
		Choice2: LayerChoice{
			Text:        "Affronter le Boss du Trauma (Courage)",
			Risk:        3,
			Reward:      10,
			NextLayer:   5, // Victoire - éveil
			FlavorText:  "Vous levez la tête. Cette fois, vous êtes assez fort. Le combat final commence.",
		},
		IsBoss: true,
	},
}

// ExploreLayer gère l'exploration d'une couche
func ExploreLayer(player *character.Character) {
	currentLayer := GetPlayerLayer(player)
	layer := Layers[currentLayer-1]
	
	fmt.Println(strings.Repeat("═",60))
	fmt.Printf("🌀 %s - Niveau %d 🌀\n", layer.Name, layer.Level)
	fmt.Println(strings.Repeat("═",60))
	fmt.Println(layer.Description)
	
	if layer.IsBoss {
		handleBossLayer(player, layer)
		return
	}
	
	// Afficher les choix
	fmt.Println("\n💭 Que choisit votre esprit ?")
	fmt.Printf("1. %s\n", layer.Choice1.Text)
	fmt.Printf("2. %s\n", layer.Choice2.Text)
	
	var choice int
	fmt.Print("👉 Votre décision : ")
	fmt.Scanln(&choice)
	
	var selectedChoice LayerChoice
	switch choice {
	case 1:
		selectedChoice = layer.Choice1
	case 2:
		selectedChoice = layer.Choice2
	default:
		fmt.Println("❌ Dans l'hésitation, votre esprit choisit la prudence...")
		selectedChoice = layer.Choice1
	}
	
	// Afficher le texte d'ambiance
	fmt.Printf("\n🌙 %s\n", selectedChoice.FlavorText)
	
	// Générer combat en fonction du risque
	generateCombatForRisk(player, selectedChoice.Risk, selectedChoice.Reward)
	
	// Progression vers la couche suivante
	setPlayerLayer(player, selectedChoice.NextLayer)
	
	// Débloquer nouveaux items marchand
	unlockMerchantItems(player, selectedChoice.NextLayer)
}

// GetPlayerLayer retourne la couche actuelle du joueur
func GetPlayerLayer(player *character.Character) int {
	// Utiliser le niveau du personnage comme couche
	if player.Level > 4 {
		return 4 // Boss level
	}
	return player.Level
}

// setPlayerLayer définit la couche du joueur
func setPlayerLayer(player *character.Character, newLayer int) {
	if newLayer > player.Level {
		player.Level = newLayer
		fmt.Printf("🌟 Vous accédez à une nouvelle couche de conscience : Niveau %d\n", newLayer)
	}
}

// generateCombatForRisk génère un combat basé sur le risque choisi
func generateCombatForRisk(player *character.Character, risk, reward int) {
	fmt.Printf("\n⚔️ Les ombres de cette couche prennent forme...\n")
	
	var monster combat.Monster
	switch risk {
	case 1: // Sûr - monstre faible
		monster = combat.InitWeakGoblin()
		fmt.Println("Un petit gobelin craintif vous fait face...")
	case 2: // Risqué - monstre normal  
		monster = combat.InitGoblin()
		fmt.Println("Un gobelin des profondeurs surgit des ténèbres...")
	case 3: // Très risqué - monstre fort
		monster = combat.InitStrongGoblin() 
		fmt.Println("Une créature terrifiante se matérialise...")
	}
	
	// Combat
	victory := combat.LayerFight(player, &monster)
	
	if victory {
		// Récompenses multipliées
		goldReward := 20 * reward
		xpReward := 25 * reward
		
		player.Money += goldReward
		player.GainXP(xpReward)
		
		fmt.Printf("💰 Vous récupérez %d fragments de mémoire !\n", goldReward)
		
		// Drop d'objets de craft selon le niveau
		dropCraftMaterials(player, risk)
	}
}

// dropCraftMaterials fait dropper des matériaux
func dropCraftMaterials(player *character.Character, risk int) {
	materials := []string{"Cuir de Sanglier", "Plume de Corbeau"}
	if risk >= 2 {
		materials = append(materials, "Fourrure de Loup")
	}
	if risk >= 3 {
		materials = append(materials, "Peau de Troll")
	}
	
	// Drop aléatoire
	if len(materials) > 0 {
		dropped := materials[0] // Simplification - drop le premier
		player.AddToInventory(dropped)
		fmt.Printf("🎁 Vous trouvez : %s\n", dropped)
	}
}

// handleBossLayer gère la couche boss
func handleBossLayer(player *character.Character, layer Layer) {
	fmt.Println("\n💀 Vous sentez une présence immense et terrifiante...")
	fmt.Println("💀 C'est LUI. Votre trauma originel.")
	
	// Vérifier si le joueur a assez exploré
	if hasExploredEnough(player) {
		fmt.Println("✨ Mais vous n'êtes plus le même. Vous avez la force de le vaincre.")
		
		fmt.Printf("1. %s\n", layer.Choice1.Text)
		fmt.Printf("2. %s\n", layer.Choice2.Text)
		
		var choice int
		fmt.Print("👉 Votre choix final : ")
		fmt.Scanln(&choice)
		
		if choice == 2 {
			// Combat contre le boss
			fmt.Printf("\n%s\n", layer.Choice2.FlavorText)
			bossVictory := combat.FinalBossFight(player)
			
			if bossVictory {
				fmt.Println("\n🌟 ═══ ÉVEIL COMPLET ═══ 🌟")
				fmt.Println("Vous ouvrez les yeux dans la vraie vie.")
				fmt.Println("Vos traumatismes n'ont plus de pouvoir sur vous.")
				fmt.Println("Vous avez GAGNÉ. Félicitations.")
				fmt.Println(strings.Repeat("═", 40))
			} else {
				gameOver(player, "Vous n'étiez pas encore assez fort...")
			}
		} else {
			fmt.Printf("\n%s\n", layer.Choice1.FlavorText)
			gameOver(player, "Vous avez abandonné face à vos démons.")
		}
	} else {
		fmt.Println("💀 Vous réalisez avec horreur que vous n'êtes pas prêt...")
		fmt.Println("💀 Vous n'avez pas assez exploré vos profondeurs.")
		gameOver(player, "Vos démons vous submergent.")
	}
}

// hasExploredEnough vérifie si le joueur a assez exploré
func hasExploredEnough(player *character.Character) bool {
	// Critères : Level 4+, certain équipement, argent suffisant
	return player.Level >= 4 && 
		   player.Money >= 100 && 
		   (player.Equipment.Head != "" || player.Equipment.Chest != "" || player.Equipment.Feet != "")
}

// gameOver gère la fin du jeu (échec)
func gameOver(player *character.Character, reason string) {
	fmt.Println("\n💀 ═══ GAME OVER ═══ 💀")
	fmt.Printf("%s\n", reason)
	fmt.Println("Dans votre monde onirique ET dans la réalité,")
	fmt.Println("votre esprit s'éteint à jamais...")
	fmt.Println(strings.Repeat("═",30))
	
	fmt.Println("\n🔄 Voulez-vous recommencer avec un nouvel esprit ?")
	fmt.Println("1. Oui - Nouvelle tentative")
	fmt.Println("2. Non - Accepter l'échec")
	
	var choice int
	fmt.Scanln(&choice)
	
	if choice == 1 {
		// Redémarrer le jeu
		MainMenu()
	}
}

// unlockMerchantItems débloque de nouveaux items selon le niveau
func unlockMerchantItems(player *character.Character, layer int) {
	switch layer {
	case 2:
		fmt.Println("🏪 Le marchand a de nouveaux souvenirs à vous proposer...")
	case 3:
		fmt.Println("🏪 Des artefacts plus puissants apparaissent chez le marchand...")
	case 4:
		fmt.Println("🏪 Le marchand vous regarde avec respect. Des reliques légendaires vous attendent...")
	}
}



