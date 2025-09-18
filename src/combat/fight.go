package combat

import (
	"fmt"
	"math/rand"
	"somnium/character"
	"somnium/quest"
	"somnium/ui"
	"strings"
	"time"
)

// Résultats possibles d'un tour
type TurnResult int

// Résultats possibles d'un tour
const (
	TurnContinue TurnResult = iota
	TurnVictory
	TurnFlee
	TurnDefeat
)

// État du combat
type CombatState struct {
	Turn         int
	PlayerAlive  bool
	MonsterAlive bool
	ShieldTurns  int
	PlayerFirst  bool
}

//  Combat générique
func Fight(player *character.Character, monster *Monster, isTraining bool, isBoss bool) bool {
	rand.Seed(time.Now().UnixNano())

	playerFirst := DetermineFirstPlayer(player, monster)

	state := CombatState{
		Turn:         1,
		PlayerAlive:  true,
		MonsterAlive: true,
		ShieldTurns:  0,
		PlayerFirst:  playerFirst,
	}

	ui.PrintInfo(fmt.Sprintf("\n⚔️ %s apparaît ! (%d/%d PV)", monster.Name, monster.PvCurr, monster.PvMax))
	monster.DisplayInfo()

	for state.PlayerAlive && state.MonsterAlive {
		ui.PrintInfo(fmt.Sprintf("\n=== Tour %d ===", state.Turn))

		var result TurnResult

		if state.PlayerFirst {
			// Joueur commence
			result = CharacterTurnNew(player, monster, &state)

			switch result {
			case TurnVictory:
				handleVictory(player, monster, isTraining, isBoss)
				return true
			case TurnFlee:
				handleFlee(player, monster)
				return false // ✅ FUITE = ÉCHEC, pas victoire
			case TurnDefeat:
				handleDefeat(player, monster, isTraining)
				return false
			}

			// Tour du monstre si combat continue
			if result == TurnContinue {
				monsterAttackPattern(monster, player, &state)
				if player.IsDead() {
					handleDefeat(player, monster, isTraining)
					return false
				}
			}
		} else {
			// Monstre commence
			monsterAttackPattern(monster, player, &state)
			if player.IsDead() {
				handleDefeat(player, monster, isTraining)
				return false
			}

			// Tour du joueur
			result = CharacterTurnNew(player, monster, &state)

			switch result {
			case TurnVictory:
				handleVictory(player, monster, isTraining, isBoss)
				return true
			case TurnFlee:
				handleFlee(player, monster)
				return false // ✅ FUITE = ÉCHEC
			case TurnDefeat:
				handleDefeat(player, monster, isTraining)
				return false
			}
		}

		state.Turn++
	}

	return false
}

//  Tour
func CharacterTurnNew(player *character.Character, monster *Monster, state *CombatState) TurnResult {
	ui.PrintInfo(fmt.Sprintf("\n⚔️ C'est votre tour, %s !", player.Name))
	ui.PrintInfo(fmt.Sprintf("💖 Vos PV : %d/%d | 🔮 Mana : %d/%d",
		player.PvCurr, player.PvMax, player.ManaCurr, player.ManaMax))
	ui.PrintInfo(fmt.Sprintf("👹 %s : %d/%d PV", monster.Name, monster.PvCurr, monster.PvMax))

	ui.PrintInfo("\n--- Menu de combat ---")
	ui.PrintInfo("1. Attaquer (Coup de poing)")
	ui.PrintInfo("2. Sorts")
	ui.PrintInfo("3. Inventaire")
	ui.PrintInfo("4. Fuir")

	var choice int
	ui.PrintInfo("👉 Votre choix : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		CoupDePoing(player, monster)
	case 2:
		handleSpellMenu(player, monster, state)
	case 3:
		character.AccessInventory(player)
		var invChoice int
		fmt.Scanln(&invChoice)
		if invChoice < 1 || invChoice > len(player.Inventory) {
			ui.PrintError("❌ Choix invalide")
			break
		}
	case 4:
		// ✅ FUITE CORRECTE
		ui.PrintInfo("💨 Vous tentez de fuir le combat...")
		return TurnFlee // ✅ Retourner FUITE, pas false
	default:
		ui.PrintError("❌ Choix invalide, vous perdez votre tour !")
		return TurnContinue
	}

	// Vérifier l'état du monstre après l'action
	if monster.IsDead() {
		ui.PrintSuccess(fmt.Sprintf("💀 %s est vaincu !", monster.Name))
		return TurnVictory
	} else {
		ui.PrintInfo(fmt.Sprintf("👹 %s : %d/%d PV restants", monster.Name, monster.PvCurr, monster.PvMax))
		return TurnContinue
	}
}

// ✅ AMÉLIORÉ : Attaque du monstre plus forte
func monsterAttackPattern(monster *Monster, player *character.Character, state *CombatState) {
	damage := monster.Attack

	// ✅ RENFORCÉ : Attaques spéciales plus fréquentes et plus fortes
	if state.Turn%2 == 0 { // Attaque spéciale tous les 2 tours au lieu de 3
		damage = int(float64(damage) * 1.8) // 80% bonus au lieu de 50%
		ui.PrintError(fmt.Sprintf("⚡ %s concentre toute sa rage !", monster.Name))
	}

	// Réduction par bouclier
	if state.ShieldTurns > 0 {
		damage /= 2
		state.ShieldTurns--
		ui.PrintInfo("🛡️ Bouclier réduit les dégâts de moitié !")
	}

	player.PvCurr -= damage
	if player.PvCurr < 0 {
		player.PvCurr = 0
	}

	ui.PrintError(fmt.Sprintf("👹 %s attaque %s et inflige %d dégâts ! (%d/%d PV restants)",
		monster.Name, player.Name, damage, player.PvCurr, player.PvMax))
}

//  Victoire
func handleVictory(player *character.Character, monster *Monster, isTraining bool, isBoss bool) {
	ui.PrintSuccess(fmt.Sprintf("🎉 Vous avez vaincu %s !", monster.Name))

	// ✅ NOUVEAU : Débloquer succès première victoire
	if !isTraining {
		player.UnlockAchievement("first_victory")
	}

	if !isTraining {
		// ✅ AJOUTER : Mise à jour des quêtes
		quest.UpdateQuestProgress("kill", monster.Name, 1)

		// XP et loot
		xpGain := 35 + (monster.Level * 15) // ✅ RENFORCÉ : Plus d'XP
		if isBoss {
			xpGain *= 3 // ✅ RENFORCÉ : Boss donne 3x XP au lieu de 2x
			player.UnlockAchievement("boss_slayer")
		}
		player.GainXP(xpGain)

		// Loot amélioré
		if len(monster.Loot) > 0 {
			// ✅ CORRIGÉ : Loot aléatoire plus intéressant
			lootIndex := rand.Intn(len(monster.Loot))
			loot := monster.Loot[lootIndex]
			
			if player.AddToInventory(loot) {
				ui.PrintSuccess(fmt.Sprintf("🎁 Vous trouvez : %s", loot))
				
				// ✅ AMÉLIORATION : Si c'est une arme, proposer de l'équiper
				if weapon, isWeapon := character.Weapons[loot]; isWeapon {
					ui.PrintInfo(fmt.Sprintf("⚔️ %s est une arme ! (+%d dégâts)", loot, weapon.Damage))
					ui.PrintInfo("Voulez-vous l'équiper maintenant ? (o/n)")
					var equipChoice string
					fmt.Scanln(&equipChoice)
					if strings.ToLower(equipChoice) == "o" || strings.ToLower(equipChoice) == "oui" {
						player.EquipWeapon(loot)
					}
				}

				// ✅ AJOUTER : Mise à jour quête collect
				quest.UpdateQuestProgress("collect", loot, 1)
			} else {
				ui.PrintError("🎒 Inventaire plein ! Objet perdu...")
			}
		}

		// ✅ AMÉLIORATION : Chance d'or bonus
		bonusGold := 10 + rand.Intn(monster.Level * 5)
		player.Money += bonusGold
		ui.PrintInfo(fmt.Sprintf("💰 Vous trouvez aussi %d fragments !", bonusGold))
	} else {
		player.GainXP(25)
		if player.IsDead() {
			player.Resurrect()
		}
	}
}

// Défaite
func handleDefeat(player *character.Character, monster *Monster, isTraining bool) {
	ui.PrintError(fmt.Sprintf("💀 Vous avez été vaincu par %s...", monster.Name))
	if isTraining {
		player.Resurrect()
		ui.PrintSuccess("✨ Vous revenez à la vie pour continuer l'entraînement.")
	} else {
		player.Resurrect()
		ui.PrintError("✨ Vous êtes soigné mais perdez votre progression.")
	}
}

// 🎓 Combat d'entraînement
func TrainingFight(player *character.Character) {
	goblin := InitGoblin()
	Fight(player, &goblin, true, false)
}

//  Combat normal
func StartFight(player *character.Character, monster Monster) error {
	victory := Fight(player, &monster, false, false)
	if !victory {
		return fmt.Errorf("combat perdu contre %s", monster.Name)
	}
	return nil
}

//  Combat de boss
func StartBossFight(player *character.Character, boss Monster) bool {
	return Fight(player, &boss, false, true)
}

// Gestion du menu des sorts
func handleSpellMenu(player *character.Character, monster *Monster, state *CombatState) {
	ui.PrintInfo("\n--- Sorts disponibles ---")
	ui.PrintInfo("1. Boule de Feu (15 mana)")
	ui.PrintInfo("2. Soin (10 mana)")
	ui.PrintInfo("3. Bouclier (8 mana)")

	var spellChoice int
	ui.PrintInfo("👉 Choix du sort : ")
	fmt.Scanln(&spellChoice)

	switch spellChoice {
	case 1:
		if player.CanCastSpell("Boule de feu") {
			BouleDeFeu(player, monster)
		} else {
			ui.PrintError("❌ Vous ne connaissez pas ce sort !")
		}
	case 2:
		if ConsumeMana(player, "Soin") {
			heal := 20
			player.PvCurr += heal
			if player.PvCurr > player.PvMax {
				player.PvCurr = player.PvMax
			}
			ui.PrintSuccess(fmt.Sprintf("💖 %s se soigne de %d PV (%d/%d)",
				player.Name, heal, player.PvCurr, player.PvMax))
		} else {
			ui.PrintError("❌ Pas assez de mana !")
		}
	case 3:
		if ConsumeMana(player, "Bouclier") {
			state.ShieldTurns = 3
			ui.PrintSuccess(fmt.Sprintf("🛡️ %s active un bouclier pour 3 tours !", player.Name))
		} else {
			ui.PrintError("❌ Pas assez de mana !")
		}
	default:
		ui.PrintError("❌ Sort invalide.")
	}
}

//	Détermine qui commence
func DetermineFirstPlayer(player *character.Character, monster *Monster) bool {
	ui.ClearScreen(player)
	ui.PrintInfo("\n🎲 ═══ JET D'INITIATIVE ═══ 🎲")

	// Le joueur lance son initiative
	player.RollInitiative()
	ui.PrintInfo(fmt.Sprintf("🎲 %s obtient %d d'initiative !", monster.Name, monster.Initiative))

	playerWins := player.Initiative >= monster.Initiative

	if playerWins {
		ui.PrintSuccess(fmt.Sprintf("⚡ %s prend l'initiative ! Vous commencez.", player.Name))
	} else {
		ui.PrintError(fmt.Sprintf("⚡ %s est plus rapide ! Il commence.", monster.Name))
	}

	ui.PrintInfo(strings.Repeat("═", 40))
	return playerWins
}

// Gestion de la fuite
func handleFlee(player *character.Character, monster *Monster) {
	ui.PrintInfo(fmt.Sprintf("💨 Vous fuyez devant %s !", monster.Name))
	ui.PrintInfo("Vous retournez au menu principal sans récompense.")

	// Optionnel : petite pénalité pour la fuite
	if player.PvCurr > 10 {
		player.PvCurr -= 10
		ui.PrintError("💔 La fuite vous coûte 10 PV (stress)")
	}
}