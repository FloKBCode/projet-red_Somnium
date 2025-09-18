package combat

import (
	"fmt"
	"math/rand"
	"time"
	"somnium/character"
	"somnium/ui"
	"strings"
)

// État du combat
type CombatState struct {
	Turn          int
	PlayerAlive   bool
	MonsterAlive  bool
	ShieldTurns   int
	PlayerFirst   bool
}

// ⚔️ Combat générique
func Fight(player *character.Character, monster *Monster, isTraining bool, isBoss bool) bool {
	rand.Seed(time.Now().UnixNano())

	state := CombatState{
		Turn:          1,
		PlayerAlive:   true,
		MonsterAlive:  true,
		ShieldTurns:   0,
	}

	fmt.Printf("\n⚔️ %s apparaît ! (%d/%d PV)\n", monster.Name, monster.PvCurr, monster.PvMax)

	for state.PlayerAlive && state.MonsterAlive {
		fmt.Printf("\n=== Tour %d ===\n", state.Turn)

		// Tour du joueur
		state.MonsterAlive = CharacterTurn(player, monster, &state)
		if !state.MonsterAlive {
			handleVictory(player, monster, isTraining, isBoss)
			return true
		}

		// Tour du monstre
		monsterAttackPattern(monster, player, &state)
		if player.IsDead() {
			state.PlayerAlive = false
			handleDefeat(player, monster, isTraining)
			return false
		}

		state.Turn++
	}

	return false
}

// 🎯 Tour du joueur
func CharacterTurn(player *character.Character, monster *Monster, state *CombatState) bool {
	ui.PrintInfo(fmt.Sprintf("\n⚔️ C'est votre tour, %s !", player.Name))
	ui.PrintInfo(fmt.Sprintf("💖 PV : %d/%d | 🔮 Mana : %d/%d",
		player.PvCurr, player.PvMax, player.ManaCurr, player.ManaMax))

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
		if player.CountItem("Potion de vie") > 0 {
			player.TakePot()
		} else {
			ui.PrintError("❌ Aucune potion disponible !")
		}
	case 4:
		ui.PrintInfo("💨 Vous fuyez le combat...")
		return false
	default:
		ui.PrintError("❌ Choix invalide, vous perdez votre tour !")
	}

	return !monster.IsDead()
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

	fmt.Printf("👹 %s attaque %s et inflige %d dégâts ! (%d/%d PV restants)\n",
		monster.Name, player.Name, damage, player.PvCurr, player.PvMax)
}

// 🎉 Victoire
func handleVictory(player *character.Character, monster *Monster, isTraining bool, isBoss bool) {
	fmt.Printf("🎉 Vous avez vaincu %s !\n", monster.Name)

	if isTraining {
		player.GainXP(25)
		player.Resurrect()
		return
	}

	// XP normal
	xpGain := 25 + (monster.Level * 10)
	if isBoss {
		xpGain *= 2
	}
	player.GainXP(xpGain)

	// Loot
	if len(monster.Loot) > 0 {
		loot := monster.Loot[0]
		player.AddToInventory(loot)
		fmt.Printf("🎁 Vous trouvez : %s\n", loot)
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