package game

import (
	"fmt"
	"math/rand"
	"somnium/character"
	"somnium/combat"
	"somnium/ui"
	"strings"
	"time"
)

// RoomType définit les types de salles disponibles
type RoomType int

const (
	RoomCombat RoomType = iota
	RoomTreasure
	RoomTrap
	RoomRiddle
	RoomShrine
	RoomMerchant
	RoomHeal
	RoomEvent
	RoomBoss
)

// Room représente une salle du donjon
type Room struct {
	Type        RoomType
	Name        string
	Description string
	Difficulty  int // 1-3
	Completed   bool
}

// Event représente un événement aléatoire
type Event struct {
	Name        string
	Description string
	Choices     []EventChoice
}

type EventChoice struct {
	Text   string
	Effect func(*character.Character) string
	Risk   int // 0-3
}

// =============== GÉNÉRATION DE SALLES ===============

// GenerateRoom génère une salle aléatoire selon la couche
func GenerateRoom(layer int, player *character.Character) Room {
	rand.Seed(time.Now().UnixNano())

	// Probabilités par couche
	var roomTypes []RoomType

	switch layer {
	case 1: // Surface - salles faciles
		roomTypes = []RoomType{
			RoomCombat, RoomCombat, RoomCombat, // 30%
			RoomTreasure, RoomTreasure, // 20%
			RoomHeal, RoomHeal, // 20%
			RoomRiddle,           // 10%
			RoomEvent, RoomEvent, // 20%
		}
	case 2: // Vallée des Regrets
		roomTypes = []RoomType{
			RoomCombat, RoomCombat, RoomCombat, // 30%
			RoomTrap, RoomTrap, // 20%
			RoomTreasure,           // 10%
			RoomRiddle, RoomRiddle, // 20%
			RoomEvent,  // 10%
			RoomShrine, // 10%
		}
	case 3: // Gouffre des Peurs
		roomTypes = []RoomType{
			RoomCombat, RoomCombat, // 20%
			RoomTrap, RoomTrap, RoomTrap, // 30%
			RoomRiddle,           // 10%
			RoomEvent, RoomEvent, // 20%
			RoomShrine,   // 10%
			RoomMerchant, // 10%
		}
	default:
		roomTypes = []RoomType{RoomCombat}
	}

	selectedType := roomTypes[rand.Intn(len(roomTypes))]
	return createRoom(selectedType, layer)
}

// createRoom crée une salle spécifique
func createRoom(roomType RoomType, layer int) Room {
	switch roomType {
	case RoomCombat:
		return Room{
			Type:        RoomCombat,
			Name:        "Arène des Ombres",
			Description: "Des créatures oniriques rodent dans cette salle sombre...",
			Difficulty:  layer,
		}
	case RoomTreasure:
		return Room{
			Type:        RoomTreasure,
			Name:        "Chambre au Trésor",
			Description: "Un coffre ancien brille faiblement dans les ténèbres.",
			Difficulty:  1,
		}
	case RoomTrap:
		return Room{
			Type:        RoomTrap,
			Name:        "Salle Piégée",
			Description: "Le sol craque sous vos pas... Quelque chose ne va pas ici.",
			Difficulty:  layer,
		}
	case RoomRiddle:
		return Room{
			Type:        RoomRiddle,
			Name:        "Sanctuaire des Énigmes",
			Description: "Une voix éthérée résonne : 'Résolvez mon mystère pour passer...'",
			Difficulty:  layer,
		}
	case RoomShrine:
		return Room{
			Type:        RoomShrine,
			Name:        "Autel de Conscience",
			Description: "Un autel mystique dégage une aura apaisante.",
			Difficulty:  1,
		}
	case RoomMerchant:
		return Room{
			Type:        RoomMerchant,
			Name:        "Échoppe Fantôme",
			Description: "Un marchand spectral vous accueille d'un sourire inquiétant.",
			Difficulty:  1,
		}
	case RoomHeal:
		return Room{
			Type:        RoomHeal,
			Name:        "Source de Sérénité",
			Description: "Une fontaine cristalline émane une énergie curative.",
			Difficulty:  1,
		}
	case RoomEvent:
		return Room{
			Type:        RoomEvent,
			Name:        "Nexus des Possibles",
			Description: "La réalité semble fluctuer ici... Plusieurs chemins s'ouvrent.",
			Difficulty:  layer,
		}
	case RoomBoss:
		return Room{
			Type:        RoomBoss,
			Name:        "Salle du Boss Final",
			Description: "Le gardien ultime du donjon vous attend dans les ténèbres absolues...",
			Difficulty:  3,
		}
	default:
		return createRoom(RoomCombat, layer)
	}
}

// =============== GESTION DES SALLES ===============

// ExploreRoom gère l'exploration d'une salle
func ExploreRoom(room Room, player *character.Character) error {
	ui.PrintInfo(fmt.Sprintf("\n🏛️  === %s ===", room.Name))
	ui.PrintInfo(room.Description)

	switch room.Type {
	case RoomCombat:
		return handleCombatRoom(player, room.Difficulty)
	case RoomTreasure:
		return handleTreasureRoom(player, room.Difficulty)
	case RoomTrap:
		return handleTrapRoom(player, room.Difficulty)
	case RoomRiddle:
		return handleRiddleRoom(player, room.Difficulty)
	case RoomShrine:
		return handleShrineRoom(player)
	case RoomMerchant:
		return handleMerchantRoom(player)
	case RoomHeal:
		return handleHealRoom(player)
	case RoomEvent:
		return handleEventRoom(player, room.Difficulty)
	case RoomBoss:
		return handleCombatRoom(player, room.Difficulty) // Traite comme un combat spécial
	default:
		return handleCombatRoom(player, room.Difficulty)
	}
}

// =============== TYPES DE SALLES SPÉCIFIQUES ===============

// 🗡️ SALLE DE COMBAT
func handleCombatRoom(player *character.Character, difficulty int) error {
	ui.PrintInfo("⚔️ Des créatures surgissent des ombres !")

	var combatDifficulty combat.Difficulty
	switch difficulty {
	case 1:
		combatDifficulty = combat.Easy
	case 2:
		combatDifficulty = combat.Normal
	default:
		combatDifficulty = combat.Hard
	}

	monster := combat.GenerateMonster(player.Level, combatDifficulty)
	return combat.StartFight(player, monster)
}

// 💰 SALLE AU TRÉSOR
func handleTreasureRoom(player *character.Character, difficulty int) error {
	ui.PrintInfo("✨ Un coffre ancien vous fait face...")
	fmt.Println("1. Ouvrir prudemment le coffre")
	fmt.Println("2. Forcer l'ouverture")
	fmt.Println("3. Ignorer et continuer")

	var choice int
	fmt.Print("👉 Votre choix : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1: // Prudent
		ui.PrintInfo("Vous examinez le coffre avec attention...")
		time.Sleep(2 * time.Second)

		if rand.Intn(100) < 80 { // 80% de réussite
			reward := generateTreasure(difficulty)
			ui.PrintSuccess(fmt.Sprintf("🎁 Vous trouvez : %s !", reward.Name))
			applyTreasureReward(player, reward)
		} else {
			ui.PrintError("💥 Le coffre était piégé ! Vous perdez 15 PV.")
			player.PvCurr -= 15
			if player.PvCurr < 0 {
				player.PvCurr = 0
			}
		}

	case 2: // Forcer
		ui.PrintInfo("Vous forcez brutalement l'ouverture...")

		if rand.Intn(100) < 60 { // 60% de réussite
			reward := generateTreasure(difficulty)
			ui.PrintSuccess(fmt.Sprintf("🎁 Vous trouvez : %s !", reward.Name))
			applyTreasureReward(player, reward)
		} else {
			ui.PrintError("💥 PIÈGE ! Une explosion vous blesse gravement (-25 PV).")
			player.PvCurr -= 25
			if player.PvCurr < 0 {
				player.PvCurr = 0
			}
		}

	case 3: // Ignorer
		ui.PrintInfo("Vous ignorez le coffre et continuez votre chemin.")
		ui.PrintInfo("Parfois, la prudence est la meilleure stratégie...")
	}

	return nil
}

// 🕳️ SALLE PIÉGÉE
func handleTrapRoom(player *character.Character, difficulty int) error {
	traps := []string{
		"Lames Tournoyantes",
		"Fosse aux Piques",
		"Gaz Toxique",
		"Sol qui s'effondre",
		"Flèches Empoisonnées",
	}

	selectedTrap := traps[rand.Intn(len(traps))]
	ui.PrintError(fmt.Sprintf("🕳️ PIÈGE ACTIVÉ : %s !", selectedTrap))

	fmt.Println("Comment réagissez-vous ?")
	fmt.Println("1. Esquiver rapidement (Agilité)")
	fmt.Println("2. Encaisser et foncer (Endurance)")
	fmt.Println("3. Analyser et contourner (Intelligence)")

	var choice int
	fmt.Print("👉 Votre choix : ")
	fmt.Scanln(&choice)

	// Calcul de réussite selon la classe
	successChance := 50
	switch choice {
	case 1: // Agilité
		if player.Class == "Voleur" {
			successChance = 80
		}
	case 2: // Endurance
		if player.Class == "Guerrier" {
			successChance = 80
		}
	case 3: // Intelligence
		if player.Class == "Mage" || player.Class == "Occultiste" {
			successChance = 80
		}
	}

	if rand.Intn(100) < successChance {
		ui.PrintSuccess("✅ Vous évitez habilement le piège !")
		// Bonus pour réussir
		bonus := 10 + (difficulty * 5)
		player.Money += bonus
		ui.PrintSuccess(fmt.Sprintf("💰 Vous trouvez %d fragments dans les débris !", bonus))
	} else {
		damage := 15 + (difficulty * 5)
		ui.PrintError(fmt.Sprintf("💥 Le piège vous touche ! Vous perdez %d PV.", damage))
		player.PvCurr -= damage
		if player.PvCurr < 0 {
			player.PvCurr = 0
		}
	}

	return nil
}

// 🧩 SALLE D'ÉNIGME
func handleRiddleRoom(player *character.Character, difficulty int) error {
	riddles := []Riddle{
		{
			Question: "Je grandis quand vous me nourrissez, mais je meurs quand vous me donnez à boire. Que suis-je ?",
			Answer:   "feu",
			Hint:     "Je danse dans l'obscurité...",
		},
		{
			Question: "Plus j'ai de gardiens, moins je suis gardé. Que suis-je ?",
			Answer:   "secret",
			Hint:     "Chuchoté entre amis...",
		},
		{
			Question: "Je n'ai ni commencement ni fin, mais je contiens tout. Que suis-je ?",
			Answer:   "cercle",
			Hint:     "Forme parfaite sans angles...",
		},
	}

	selectedRiddle := riddles[rand.Intn(len(riddles))]

	ui.PrintInfo("🧩 Une voix mystérieuse résonne :")
	ui.PrintInfo(fmt.Sprintf("'%s'", selectedRiddle.Question))

	fmt.Println("1. Répondre à l'énigme")
	fmt.Println("2. Demander un indice (-10 fragments)")
	fmt.Println("3. Forcer le passage (combat)")

	var choice int
	fmt.Print("👉 Votre choix : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		return solveRiddle(player, selectedRiddle, difficulty)
	case 2:
		if player.Money >= 10 {
			player.Money -= 10
			ui.PrintInfo(fmt.Sprintf("💡 Indice : %s", selectedRiddle.Hint))
			return solveRiddle(player, selectedRiddle, difficulty)
		} else {
			ui.PrintError("❌ Pas assez de fragments pour un indice.")
			return solveRiddle(player, selectedRiddle, difficulty)
		}
	case 3:
		ui.PrintError("⚔️ Votre refus déclenche la colère du gardien !")
		return handleCombatRoom(player, difficulty+1)
	}

	return nil
}

// ⛩️ AUTEL DE CONSCIENCE
func handleShrineRoom(player *character.Character) error {
	ui.PrintInfo("⛩️ Vous vous approchez de l'autel mystique...")
	ui.PrintInfo("Une énergie apaisante vous enveloppe.")

	fmt.Println("Que souhaitez-vous offrir ?")
	fmt.Println("1. 20 fragments (Bénédiction mineure)")
	fmt.Println("2. 50 fragments (Bénédiction majeure)")
	fmt.Println("3. Une partie de votre essence - 10 PV (Bénédiction divine)")
	fmt.Println("4. Rien - continuer")

	var choice int
	fmt.Print("👉 Votre choix : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		if player.Money >= 20 {
			player.Money -= 20
			player.PvMax += 5
			ui.PrintSuccess("✨ Bénédiction mineure : +5 PV maximum permanents !")
		} else {
			ui.PrintError("❌ Pas assez de fragments.")
		}
	case 2:
		if player.Money >= 50 {
			player.Money -= 50
			player.PvMax += 10
			player.ManaMax += 5
			ui.PrintSuccess("🌟 Bénédiction majeure : +10 PV max et +5 Mana max !")
		} else {
			ui.PrintError("❌ Pas assez de fragments.")
		}
	case 3:
		if player.PvCurr > 10 {
			player.PvCurr -= 10
			// Bonus aléatoire puissant
			bonuses := []func(){
				func() {
					player.PvMax += 20
					ui.PrintSuccess("🔥 Bénédiction divine : +20 PV maximum !")
				},
				func() {
					player.ManaMax += 15
					ui.PrintSuccess("🔮 Bénédiction divine : +15 Mana maximum !")
				},
				func() {
					newSkill := "Régénération"
					player.Skills = append(player.Skills, newSkill)
					ui.PrintSuccess("💫 Bénédiction divine : Sort de Régénération appris !")
				},
			}
			bonuses[rand.Intn(len(bonuses))]()
		} else {
			ui.PrintError("❌ Pas assez de vitalité pour ce sacrifice.")
		}
	case 4:
		ui.PrintInfo("Vous continuez respectueusement votre chemin.")
	}

	return nil
}

// 💚 SOURCE DE GUÉRISON
func handleHealRoom(player *character.Character) error {
	ui.PrintInfo("💚 Une source cristalline pulse d'une énergie curative...")

	fmt.Println("1. Boire l'eau (Restauration complète)")
	fmt.Println("2. Méditer près de la source (Restauration partielle + bonus)")
	fmt.Println("3. Remplir une gourde (Item pour plus tard)")
	fmt.Println("4. Continuer sans toucher")

	var choice int
	fmt.Print("👉 Votre choix : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		player.PvCurr = player.PvMax
		player.ManaCurr = player.ManaMax
		ui.PrintSuccess("💚 Vous êtes complètement restauré !")
	case 2:
		heal := player.PvMax / 2
		manaRestore := player.ManaMax / 2
		player.PvCurr += heal
		player.ManaCurr += manaRestore
		if player.PvCurr > player.PvMax {
			player.PvCurr = player.PvMax
		}
		if player.ManaCurr > player.ManaMax {
			player.ManaCurr = player.ManaMax
		}
		// Bonus temporaire
		ui.PrintSuccess(fmt.Sprintf("🧘 Guérison partielle : +%d PV, +%d Mana", heal, manaRestore))
		ui.PrintSuccess("✨ Bonus : Votre prochain combat sera plus facile !")
		// TODO: Implémenter bonus combat
	case 3:
		if player.AddToInventory("Gourde de Source Curative") {
			ui.PrintSuccess("🧪 Gourde de Source Curative ajoutée à l'inventaire !")
		} else {
			ui.PrintError("🎒 Inventaire plein !")
		}
	case 4:
		ui.PrintInfo("Vous respectez la source et continuez.")
	}

	return nil
}

// 🎲 ÉVÉNEMENT ALÉATOIRE
func handleEventRoom(player *character.Character, difficulty int) error {
	events := []Event{
		{
			Name:        "Le Choix du Miroir",
			Description: "Un miroir fêlé vous montre deux reflets différents...",
			Choices: []EventChoice{
				{
					Text: "Toucher le reflet courageux",
					Effect: func(p *character.Character) string {
						p.PvMax += 10
						return "Votre courage grandit ! +10 PV maximum."
					},
					Risk: 1,
				},
				{
					Text: "Toucher le reflet sage",
					Effect: func(p *character.Character) string {
						p.ManaMax += 8
						return "Votre sagesse s'approfondit ! +8 Mana maximum."
					},
					Risk: 1,
				},
				{
					Text: "Briser le miroir",
					Effect: func(p *character.Character) string {
						if rand.Intn(100) < 50 {
							p.Money += 100
							return "Les fragments révèlent un trésor ! +100 fragments."
						} else {
							p.PvCurr -= 20
							if p.PvCurr < 0 {
								p.PvCurr = 0
							}
							return "Les éclats vous blessent... -20 PV."
						}
					},
					Risk: 3,
				},
			},
		},
		{
			Name:        "L'Offre de l'Ombre",
			Description: "Une ombre vous propose un marché suspect...",
			Choices: []EventChoice{
				{
					Text: "Accepter l'offre (Risqué)",
					Effect: func(p *character.Character) string {
						if rand.Intn(100) < 60 {
							p.GainXP(100)
							return "L'ombre tient parole ! +100 XP."
						} else {
							p.PvMax -= 10
							return "Vous avez été trompé ! -10 PV maximum."
						}
					},
					Risk: 3,
				},
				{
					Text: "Refuser poliment",
					Effect: func(p *character.Character) string {
						p.ManaCurr += 10
						if p.ManaCurr > p.ManaMax {
							p.ManaCurr = p.ManaMax
						}
						return "Votre sagesse est récompensée. +10 Mana."
					},
					Risk: 0,
				},
			},
		},
	}

	selectedEvent := events[rand.Intn(len(events))]

	ui.PrintInfo(fmt.Sprintf("🎲 === %s ===", selectedEvent.Name))
	ui.PrintInfo(selectedEvent.Description)

	fmt.Println("\nQue faites-vous ?")
	for i, choice := range selectedEvent.Choices {
		riskText := ""
		switch choice.Risk {
		case 0:
			riskText = " (Sûr)"
		case 1:
			riskText = " (Peu risqué)"
		case 2:
			riskText = " (Risqué)"
		case 3:
			riskText = " (Très risqué)"
		}
		fmt.Printf("%d. %s%s\n", i+1, choice.Text, riskText)
	}

	var choice int
	fmt.Print("👉 Votre choix : ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(selectedEvent.Choices) {
		ui.PrintError("❌ Choix invalide, vous ignorez l'événement.")
		return nil
	}

	selectedChoice := selectedEvent.Choices[choice-1]
	result := selectedChoice.Effect(player)
	ui.PrintInfo(fmt.Sprintf("🎲 %s", result))

	return nil
}

// =============== STRUCTURES UTILITAIRES ===============

type Riddle struct {
	Question string
	Answer   string
	Hint     string
}

type Treasure struct {
	Name     string
	Type     string // "gold", "item", "equipment", "xp"
	Value    int
	ItemName string
}

func solveRiddle(player *character.Character, riddle Riddle, difficulty int) error {
	var answer string
	fmt.Print("👉 Votre réponse : ")
	fmt.Scanln(&answer)

	if strings.ToLower(strings.TrimSpace(answer)) == riddle.Answer {
		reward := 50 + (difficulty * 25)
		xpReward := 75 + (difficulty * 25)

		ui.PrintSuccess("🎉 Excellente réponse !")
		ui.PrintSuccess(fmt.Sprintf("💰 Récompense : %d fragments", reward))
		ui.PrintSuccess(fmt.Sprintf("✨ Bonus XP : %d", xpReward))

		player.Money += reward
		player.GainXP(xpReward)

		// Chance d'objet bonus
		if rand.Intn(100) < 30 {
			items := []string{"Livre de Sort: Soin", "Potion de Mana", "Fragment d'Éternité"}
			bonusItem := items[rand.Intn(len(items))]
			player.AddToInventory(bonusItem)
			ui.PrintSuccess(fmt.Sprintf("🎁 Objet bonus : %s !", bonusItem))
		}
	} else {
		ui.PrintError("❌ Mauvaise réponse...")
		ui.PrintError("Le passage se referme, vous devez combattre !")
		return handleCombatRoom(player, difficulty)
	}

	return nil
}

func generateTreasure(difficulty int) Treasure {
	treasures := []Treasure{
		{"Bourse de Fragments", "gold", 30 + (difficulty * 20), ""},
		{"Potion de Vie Majeure", "item", 0, "Potion de Vie Majeure"},
		{"Cristal de Mana", "item", 0, "Cristal de Mana"},
		{"Parchemin d'XP", "xp", 100 + (difficulty * 50), ""},
	}

	return treasures[rand.Intn(len(treasures))]
}

func applyTreasureReward(player *character.Character, treasure Treasure) {
	switch treasure.Type {
	case "gold":
		player.Money += treasure.Value
	case "item":
		player.AddToInventory(treasure.ItemName)
	case "xp":
		player.GainXP(treasure.Value)
	}
}

func handleMerchantRoom(player *character.Character) error {
	ui.PrintInfo("👻 Un marchand fantôme vous accueille...")
	ui.PrintInfo("'Des objets rares pour les âmes courageuses !'")

	// Marchand spécialisé avec objets uniques
	specialItems := map[string]int{
		"Élixir de Courage":      75,  // +20 PV max permanent
		"Tome de Maîtrise":       100, // Apprend un sort aléatoire
		"Amulette de Protection": 120, // Résistance aux pièges
		"Pierre de Résurrection": 200, // Résurrection automatique à la mort
	}

	for {
		ui.PrintInfo(fmt.Sprintf("\n💰 Vos fragments : %d", player.Money))

		i := 1
		itemList := make([]string, 0, len(specialItems))
		for item, price := range specialItems {
			fmt.Printf("%d. %s - %d fragments\n", i, item, price)
			itemList = append(itemList, item)
			i++
		}
		fmt.Println("0. Quitter")

		var choice int
		fmt.Print("👉 Votre choix : ")
		fmt.Scanln(&choice)

		if choice == 0 {
			ui.PrintInfo("Vous quittez le marchand spectral...")
			return nil
		}
		if choice < 1 || choice > len(itemList) {
			ui.PrintError("❌ Choix invalide")
			continue
		}

		selectedItem := itemList[choice-1]
		price := specialItems[selectedItem]

		if player.Money < price {
			ui.PrintError(fmt.Sprintf("💰 Pas assez de fragments ! Il vous faut %d.", price))
			continue
		}

		// Application des effets
		switch selectedItem {
		case "Élixir de Courage":
			player.PvMax += 20
			ui.PrintSuccess("💪 Élixir de Courage bu : +20 PV maximum permanents !")

		case "Tome de Maîtrise":
			spells := []string{"Chaîne d'éclairs", "Mur de glace", "Soin majeur"}
			newSpell := spells[rand.Intn(len(spells))]
			player.LearnSpell(newSpell)
			ui.PrintSuccess(fmt.Sprintf("📖 Vous apprenez un sort rare : %s !", newSpell))

		case "Amulette de Protection":
			if player.AddToInventory("Amulette de Protection") {
				ui.PrintSuccess("🛡️ Amulette ajoutée : Vous résistez mieux aux pièges !")
				// TODO : Implémenter la réduction de dégâts de pièges
			} else {
				ui.PrintError("🎒 Inventaire plein !")
				continue
			}

		case "Pierre de Résurrection":
			if player.AddToInventory("Pierre de Résurrection") {
				ui.PrintSuccess("💎 Pierre ajoutée : Vous reviendrez à la vie une fois en cas de mort !")
				// TODO : Implémenter l'effet de résurrection automatique
			} else {
				ui.PrintError("🎒 Inventaire plein !")
				continue
			}
		}

		player.Money -= price
		ui.PrintInfo(fmt.Sprintf("✅ Achat de %s pour %d fragments.", selectedItem, price))
	}
}