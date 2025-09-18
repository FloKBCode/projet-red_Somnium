package game

import (
	"fmt"
	"somnium/character"
	"somnium/ui"
)

type Achievement struct {
	ID          string
	Name        string
	Description string
	Unlocked    bool
	Reward      AchievementReward
}

type AchievementReward struct {
	Gold int
	XP   int
	Item string
}

var gameAchievements = []Achievement{
	{
		ID:          "first_victory",
		Name:        "Premier Pas",
		Description: "Gagner votre premier combat",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 50, XP: 100, Item: ""},
	},
	{
		ID:          "collector",
		Name:        "Collectionneur",
		Description: "Obtenir 10 objets différents",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 100, XP: 200, Item: "Sac magique"},
	},
	{
		ID:          "warrior",
		Name:        "Guerrier Éveillé",
		Description: "Atteindre le niveau 5",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 200, XP: 300, Item: "Élixir de Puissance"},
	},
	{
		ID:          "explorer",
		Name:        "Explorateur des Profondeurs",
		Description: "Visiter toutes les couches du Labyrinthe",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 300, XP: 500, Item: "Carte des Souvenirs"},
	},
	{
		ID:          "boss_slayer",
		Name:        "Tueur de Trauma",
		Description: "Vaincre le Boss final",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 500, XP: 1000, Item: "Couronne de la Libération"},
	},
	{
		ID:          "survivor",
		Name:        "Survivant",
		Description: "Survivre à 10 combats sans mourir",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 150, XP: 250, Item: "Amulette de Vie"},
	},
	{
		ID:          "wealthy",
		Name:        "Riche en Fragments",
		Description: "Posséder 1000 fragments",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 0, XP: 400, Item: "Coffre doré"},
	},
	{
		ID:          "master_crafter",
		Name:        "Maître Artisan",
		Description: "Forger 5 objets différents",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 250, XP: 350, Item: "Marteau enchanté"},
	},
	{
		ID:          "spell_master",
		Name:        "Maître des Arcanes",
		Description: "Apprendre 5 sorts différents",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 200, XP: 400, Item: "Bâton de maîtrise"},
	},
	{
		ID:          "nightmare_walker",
		Name:        "Marcheur des Cauchemars",
		Description: "Terminer 50 combats",
		Unlocked:    false,
		Reward:      AchievementReward{Gold: 400, XP: 600, Item: "Bottes spectrales"},
	},
}

// GetAchievement retourne un succès par son ID
func GetAchievement(id string) *Achievement {
	for i := range gameAchievements {
		if gameAchievements[i].ID == id {
			return &gameAchievements[i]
		}
	}
	return nil
}

// CheckAndUnlockAchievement vérifie et débloque un succès
func CheckAndUnlockAchievement(player *character.Character, achievementID string) {
	// Vérifier si déjà débloqué
	for _, unlockedID := range player.Achievements {
		if unlockedID == achievementID {
			return
		}
	}

	achievement := GetAchievement(achievementID)
	if achievement == nil {
		return
	}

	// Vérifier les conditions spécifiques
	canUnlock := false
	switch achievementID {
	case "first_victory":
		canUnlock = true // Débloqué manuellement dans le combat
	case "collector":
		canUnlock = countUniqueItems(player) >= 10
	case "warrior":
		canUnlock = player.Level >= 5
	case "explorer":
		canUnlock = player.CurrentLayer >= 4
	case "boss_slayer":
		canUnlock = true // Débloqué manuellement
	case "survivor":
		canUnlock = true // À implémenter avec un compteur
	case "wealthy":
		canUnlock = player.Money >= 1000
	case "master_crafter":
		canUnlock = true // À implémenter avec compteur de craft
	case "spell_master":
		canUnlock = len(player.Skills) >= 5
	case "nightmare_walker":
		canUnlock = true // À implémenter avec compteur de combats
	}

	if canUnlock {
		// Débloquer le succès
		player.Achievements = append(player.Achievements, achievementID)
		achievement.Unlocked = true

		// Afficher le message et donner les récompenses
		ui.PrintSuccess(fmt.Sprintf("🏆 SUCCÈS DÉBLOQUÉ : %s", achievement.Name))
		ui.PrintSuccess(fmt.Sprintf("📜 %s", achievement.Description))

		// Donner les récompenses
		if achievement.Reward.Gold > 0 {
			player.Money += achievement.Reward.Gold
			ui.PrintSuccess(fmt.Sprintf("💰 +%d fragments", achievement.Reward.Gold))
		}
		if achievement.Reward.XP > 0 {
			player.GainXP(achievement.Reward.XP)
		}
		if achievement.Reward.Item != "" {
			if player.AddToInventory(achievement.Reward.Item) {
				ui.PrintSuccess(fmt.Sprintf("🎁 Objet reçu : %s", achievement.Reward.Item))
			}
		}
	}
}

// ShowAchievementsMenu affiche le menu des succès
func ShowAchievementsMenu(player *character.Character) {
	ui.PrintInfo("\n🏆 === Salle des Trophées ===")
	ui.PrintInfo("Vos exploits dans le Labyrinthe des Cauchemars")

	unlockedCount := len(player.Achievements)
	totalCount := len(gameAchievements)
	
	ui.PrintInfo(fmt.Sprintf("Progression : %d/%d succès débloqués", unlockedCount, totalCount))

	ui.PrintInfo("\n--- Succès Débloqués ---")
	if unlockedCount == 0 {
		ui.PrintInfo("Aucun succès débloqué pour l'instant...")
	} else {
		for _, achievementID := range player.Achievements {
			achievement := GetAchievement(achievementID)
			if achievement != nil {
				ui.PrintSuccess(fmt.Sprintf("✅ %s - %s", achievement.Name, achievement.Description))
			}
		}
	}

	ui.PrintInfo("\n--- Succès À Débloquer ---")
	for _, achievement := range gameAchievements {
		if !isAchievementUnlocked(player, achievement.ID) {
			progress := getAchievementProgress(player, achievement.ID)
			ui.PrintError(fmt.Sprintf("❌ %s - %s %s", achievement.Name, achievement.Description, progress))
			
			// Afficher les récompenses
			if achievement.Reward.Gold > 0 || achievement.Reward.XP > 0 || achievement.Reward.Item != "" {
				rewards := "Récompense: "
				if achievement.Reward.Gold > 0 {
					rewards += fmt.Sprintf("%d fragments ", achievement.Reward.Gold)
				}
				if achievement.Reward.XP > 0 {
					rewards += fmt.Sprintf("%d XP ", achievement.Reward.XP)
				}
				if achievement.Reward.Item != "" {
					rewards += fmt.Sprintf("+ %s", achievement.Reward.Item)
				}
				ui.PrintInfo(fmt.Sprintf("   🎁 %s", rewards))
			}
		}
	}

	ui.PressEnterToContinue(player)
}

// Fonctions utilitaires
func isAchievementUnlocked(player *character.Character, achievementID string) bool {
	for _, unlockedID := range player.Achievements {
		if unlockedID == achievementID {
			return true
		}
	}
	return false
}

func getAchievementProgress(player *character.Character, achievementID string) string {
	switch achievementID {
	case "collector":
		current := countUniqueItems(player)
		return fmt.Sprintf("(%d/10)", current)
	case "warrior":
		return fmt.Sprintf("(Niveau %d/5)", player.Level)
	case "explorer":
		return fmt.Sprintf("(Couche %d/4)", player.CurrentLayer)
	case "wealthy":
		return fmt.Sprintf("(%d/1000 fragments)", player.Money)
	case "spell_master":
		return fmt.Sprintf("(%d/5 sorts)", len(player.Skills))
	default:
		return ""
	}
}

func countUniqueItems(player *character.Character) int {
	itemSet := make(map[string]bool)
	for _, item := range player.Inventory {
		itemSet[item] = true
	}
	
	// Ajouter l'équipement
	if player.Equipment.Head != "" {
		itemSet[player.Equipment.Head] = true
	}
	if player.Equipment.Chest != "" {
		itemSet[player.Equipment.Chest] = true
	}
	if player.Equipment.Feet != "" {
		itemSet[player.Equipment.Feet] = true
	}
	if player.Weapon.Name != "" {
		itemSet[player.Weapon.Name] = true
	}
	
	return len(itemSet)
}

// CheckAllAchievements vérifie tous les succès automatiques
func CheckAllAchievements(player *character.Character) {
	CheckAndUnlockAchievement(player, "collector")
	CheckAndUnlockAchievement(player, "warrior")
	CheckAndUnlockAchievement(player, "explorer")
	CheckAndUnlockAchievement(player, "wealthy")
	CheckAndUnlockAchievement(player, "spell_master")
}