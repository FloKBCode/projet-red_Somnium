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

type TurnResult int

const (
	TurnContinue TurnResult = iota
	TurnVictory
	TurnFlee
	TurnDefeat
)

type CombatState struct {
	Turn         int
	PlayerAlive  bool
	MonsterAlive bool
	ShieldTurns  int
	PlayerFirst  bool
}

// ⚔️ Combat générique
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

// 🎯 Tour du joueur
func CharacterTurnNew(player *character.Character, monster *Monster, state *CombatState) TurnResult {
	ui.PrintInfo(fmt.Sprintf("\n⚔️ C'est votre tour, %s !", player.Name))
	ui.PrintInfo(fmt.Sprintf("💖 Vos PV : %d/%d | 🔮 Mana : %d/%d",
		player.PvCurr, player.PvMax, player.ManaCurr, player.ManaMax))
	ui.PrintInfo(fmt.Sprintf("👹 %s : %d/%d PV", monster.Name, monster.PvCurr, monster.PvMax))

	fmt.Println("\n--- Menu de combat ---")
	fmt.Println("1. Attaquer (Coup de poing)")
	fmt.Println("2. Sorts")
	fmt.Println("3. Inventaire")
	fmt.Println("4. Fuir")

	var choice int
	fmt.Print("👉 Votre choix : ")
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

// 👹 Attaque du monstre
func monsterAttackPattern(monster *Monster, player *character.Character, state *CombatState) {
	damage := monster.Attack

	// Attaque spéciale tous les 3 tours
	if state.Turn%3 == 0 {
		damage = int(float64(damage) * 1.5)
		fmt.Printf("⚡ %s concentre ses forces !\n", monster.Name)
	}

	// Réduction par bouclier
	if state.ShieldTurns > 0 {
		damage /= 2
		state.ShieldTurns--
		fmt.Println("🛡️ Bouclier réduit les dégâts de moitié !")
	}

	player.PvCurr -= damage
	if player.PvCurr < 0 {
		player.PvCurr = 0
	}

	ui.PrintError(fmt.Sprintf("👹 %s attaque %s et inflige %d dégâts ! (%d/%d PV restants)\n",
		monster.Name, player.Name, damage, player.PvCurr, player.PvMax))
}

// 🎉 Victoire
func handleVictory(player *character.Character, monster *Monster, isTraining bool, isBoss bool) {
	ui.PrintSuccess(fmt.Sprintf("🎉 Vous avez vaincu %s !", monster.Name))

	if !isTraining {
		// ✅ AJOUTER : Mise à jour des quêtes
		quest.UpdateQuestProgress("kill", monster.Name, 1)

		// XP et loot
		xpGain := 25 + (monster.Level * 10)
		if isBoss {
			xpGain *= 2
		}
		player.GainXP(xpGain)

		// Loot
		if len(monster.Loot) > 0 {
			loot := monster.Loot[0]
			player.AddToInventory(loot)
			ui.PrintSuccess(fmt.Sprintf("🎁 Vous trouvez : %s", loot))

			// ✅ AJOUTER : Mise à jour quête collect
			quest.UpdateQuestProgress("collect", loot, 1)
		}
	} else {
		player.GainXP(25)
		if player.IsDead() {
			player.Resurrect()
		}
	}
}

// 💀 Défaite
func handleDefeat(player *character.Character, monster *Monster, isTraining bool) {
	fmt.Printf("💀 Vous avez été vaincu par %s...\n", monster.Name)
	if isTraining {
		player.Resurrect()
		fmt.Println("✨ Vous revenez à la vie pour continuer l'entraînement.")
	} else {
		player.Resurrect()
		fmt.Println("✨ Vous êtes soigné mais perdez votre progression.")
	}
}

// 🎓 Combat d’entraînement
func TrainingFight(player *character.Character) {
	goblin := InitGoblin()
	Fight(player, &goblin, true, false)
}

// ⚔️ Combat normal
func StartFight(player *character.Character, monster Monster) error {
	victory := Fight(player, &monster, false, false)
	if !victory {
		return fmt.Errorf("combat perdu contre %s", monster.Name)
	}
	return nil
}

// 👑 Combat de boss
func StartBossFight(player *character.Character, boss Monster) bool {
	return Fight(player, &boss, false, true)
}

func handleSpellMenu(player *character.Character, monster *Monster, state *CombatState) {
	fmt.Println("\n--- Sorts disponibles ---")
	fmt.Println("1. Boule de Feu (15 mana)")
	fmt.Println("2. Soin (10 mana)")
	fmt.Println("3. Bouclier (8 mana)")

	var spellChoice int
	fmt.Print("👉 Choix du sort : ")
	fmt.Scanln(&spellChoice)

	switch spellChoice {
	case 1:
		BouleDeFeu(player, monster)
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

func DetermineFirstPlayer(player *character.Character, monster *Monster) bool {
	fmt.Println("\n🎲 ═══ JET D'INITIATIVE ═══ 🎲")

	// Le joueur lance son initiative
	player.RollInitiative()
	fmt.Printf("🎲 %s obtient %d d'initiative !\n", monster.Name, monster.Initiative)

	playerWins := player.Initiative >= monster.Initiative

	if playerWins {
		fmt.Printf("⚡ %s prend l'initiative ! Vous commencez.\n", player.Name)
	} else {
		fmt.Printf("⚡ %s est plus rapide ! Il commence.\n", monster.Name)
	}

	fmt.Println(strings.Repeat("═", 40))
	return playerWins
}

func handleFlee(player *character.Character, monster *Monster) {
	ui.PrintInfo(fmt.Sprintf("💨 Vous fuyez devant %s !", monster.Name))
	ui.PrintInfo("Vous retournez au menu principal sans récompense.")

	// Optionnel : petite pénalité pour la fuite
	if player.PvCurr > 10 {
		player.PvCurr -= 10
		ui.PrintError("💔 La fuite vous coûte 10 PV (stress)")
	}
}
