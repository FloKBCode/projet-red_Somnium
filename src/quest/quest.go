package quest

import (
	"fmt"
	"somnium/character"
	"somnium/ui"
)

type QuestStatus int

const (
	QuestAvailable QuestStatus = iota
	QuestActive
	QuestCompleted
)

type QuestObjective struct {
	Type        string // "kill", "collect", "explore", "craft"
	Target      string
	Required    int
	Current     int
	Description string
}

type Quest struct {
	ID          int
	Title       string
	Description string
	Status      QuestStatus
	Objectives  []QuestObjective
	RewardGold  int
	RewardExp   int
	RewardItems []string
	MinLevel    int
}

var gameQuests = []Quest{
	{
		ID:          1,
		Title:       "Premiers Pas dans les Ombres",
		Description: "Explorez votre première couche de conscience et affrontez vos premières peurs.",
		Status:      QuestAvailable,
		Objectives: []QuestObjective{
			{"kill", "Gobelin", 3, 0, "Éliminer 3 Gobelins"},
			{"explore", "Surface des Rêves", 1, 0, "Explorer la Surface des Rêves"},
		},
		RewardGold:  50,
		RewardExp:   100,
		RewardItems: []string{"Potion de vie", "Potion de vie"},
		MinLevel:    1,
	},
	{
		ID:          2,
		Title:       "L'Artisan des Cauchemars",
		Description: "Apprenez l'art de la forge onirique en créant votre premier équipement.",
		Status:      QuestAvailable,
		Objectives: []QuestObjective{
			{"collect", "Cuir de Sanglier", 2, 0, "Collecter 2 Cuirs de Sanglier"},
			{"collect", "Plume de Corbeau", 2, 0, "Collecter 2 Plumes de Corbeau"},
			{"craft", "Chapeau de l'Errant", 1, 0, "Forger un Chapeau de l'Errant"},
		},
		RewardGold:  75,
		RewardExp:   150,
		RewardItems: []string{"Fourrure de Loup"},
		MinLevel:    2,
	},
	{
		ID:          3,
		Title:       "Maître des Arcanes Oniriques",
		Description: "Maîtrisez la magie des rêves en apprenant de nouveaux sorts.",
		Status:      QuestAvailable,
		Objectives: []QuestObjective{
			{"learn", "Boule de feu", 1, 0, "Apprendre le sort Boule de feu"},
			{"cast", "Boule de feu", 5, 0, "Lancer 5 Boules de feu en combat"},
		},
		RewardGold:  100,
		RewardExp:   200,
		RewardItems: []string{"Livre de Sort: Soin"},
		MinLevel:    2,
	},
	{
		ID:          4,
		Title:       "Affronter ses Démons",
		Description: "Descendez dans les profondeurs de votre psyché et affrontez des créatures plus dangereuses.",
		Status:      QuestAvailable,
		Objectives: []QuestObjective{
			{"kill", "Araignée du Cauchemar", 2, 0, "Éliminer 2 Araignées du Cauchemar"},
			{"kill", "Spectre des Ombres", 1, 0, "Éliminer 1 Spectre des Ombres"},
			{"explore", "Vallée des Regrets", 1, 0, "Explorer la Vallée des Regrets"},
		},
		RewardGold:  150,
		RewardExp:   300,
		RewardItems: []string{"Peau de Troll", "Pierre de l'Esprit"},
		MinLevel:    3,
	},
	{
		ID:          5,
		Title:       "Le Gardien du Trauma",
		Description: "Préparez-vous pour l'affrontement final contre votre trauma originel.",
		Status:      QuestAvailable,
		Objectives: []QuestObjective{
			{"level", "4", 1, 0, "Atteindre le niveau 4"},
			{"equip", "Full Set", 3, 0, "Équiper 3 pièces d'équipement"},
			{"explore", "Le Cœur du Trauma", 1, 0, "Atteindre le Cœur du Trauma"},
		},
		RewardGold:  300,
		RewardExp:   500,
		RewardItems: []string{"Cristal de Puissance"},
		MinLevel:    4,
	},
}

func ShowQuestMenu(player *character.Character) {
	for {
		ui.PrintInfo("\n📜 === Carnet de Quêtes ===")
		ui.PrintInfo(fmt.Sprintf("Niveau actuel : %d", player.Level))

		availableQuests := getAvailableQuests(player)
		activeQuests := getActiveQuests()

		fmt.Println("\n--- Quêtes Disponibles ---")
		if len(availableQuests) == 0 {
			ui.PrintInfo("Aucune nouvelle quête disponible.")
		} else {
			for i, quest := range availableQuests {
				fmt.Printf("%d. %s (Niveau %d)\n", i+1, quest.Title, quest.MinLevel)
				fmt.Printf("   %s\n", quest.Description)
			}
		}

		fmt.Println("\n--- Quêtes Actives ---")
		if len(activeQuests) == 0 {
			ui.PrintInfo("Aucune quête active.")
		} else {
			for _, quest := range activeQuests {
				fmt.Printf("🔄 %s\n", quest.Title)
				for _, obj := range quest.Objectives {
					progress := "❌"
					if obj.Current >= obj.Required {
						progress = "✅"
					}
					fmt.Printf("   %s %s (%d/%d)\n", progress, obj.Description, obj.Current, obj.Required)
				}
			}
		}

		fmt.Println("\n--- Actions ---")
		if len(availableQuests) > 0 {
			fmt.Println("A. Accepter une quête")
		}
		fmt.Println("0. Retour")

		var choice string
		fmt.Print("👉 Votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case "A", "a":
			if len(availableQuests) > 0 {
				acceptQuest(availableQuests, player)
			}
		case "0":
			return
		default:
			ui.PrintError("❌ Choix invalide")
		}
	}
}

func getAvailableQuests(player *character.Character) []Quest {
	var available []Quest
	for _, quest := range gameQuests {
		if quest.Status == QuestAvailable && quest.MinLevel <= player.Level {
			available = append(available, quest)
		}
	}
	return available
}

func getActiveQuests() []Quest {
	var active []Quest
	for _, quest := range gameQuests {
		if quest.Status == QuestActive {
			active = append(active, quest)
		}
	}
	return active
}

func getCompletedQuests() []Quest {
	var completed []Quest
	for _, quest := range gameQuests {
		if quest.Status == QuestCompleted {
			completed = append(completed, quest)
		}
	}
	return completed
}

func acceptQuest(availableQuests []Quest, player *character.Character) {
	if len(availableQuests) == 0 {
		return
	}

	fmt.Println("\nQuelle quête accepter ?")
	for i, quest := range availableQuests {
		fmt.Printf("%d. %s\n", i+1, quest.Title)
	}

	var choice int
	fmt.Print("👉 Choix : ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(availableQuests) {
		ui.PrintError("❌ Choix invalide")
		return
	}

	selectedQuest := availableQuests[choice-1]
	// Marquer la quête comme active
	for i := range gameQuests {
		if gameQuests[i].ID == selectedQuest.ID {
			gameQuests[i].Status = QuestActive
			break
		}
	}

	ui.PrintSuccess(fmt.Sprintf("✅ Quête acceptée : %s", selectedQuest.Title))
}

// Fonction pour progresser les quêtes (appelée depuis le combat, exploration, etc.)
func UpdateQuestProgress(action, target string, amount int) {
	for i := range gameQuests {
		if gameQuests[i].Status == QuestActive {
			for j := range gameQuests[i].Objectives {
				obj := &gameQuests[i].Objectives[j]
				if obj.Type == action && obj.Target == target {
					obj.Current += amount
					if obj.Current >= obj.Required {
						ui.PrintSuccess(fmt.Sprintf("🎯 Objectif complété : %s", obj.Description))
					}
				}
			}
			
			// Vérifier si la quête est complète
			completed := true
			for _, obj := range gameQuests[i].Objectives {
				if obj.Current < obj.Required {
					completed = false
					break
				}
			}
			
			if completed {
				completeQuest(&gameQuests[i])
			}
		}
	}
}

func completeQuest(quest *Quest) {
	quest.Status = QuestCompleted
	ui.PrintSuccess(fmt.Sprintf("🏆 QUÊTE TERMINÉE : %s", quest.Title))
	ui.PrintSuccess(fmt.Sprintf("💰 Récompense : %d or, %d XP", quest.RewardGold, quest.RewardExp))
}
	// TODO: Donner les récompenses au joueur
