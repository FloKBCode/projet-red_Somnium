package combat

import (
	"fmt"
	"somnium/character"
)

// TraumaBoss représente le boss final
type TraumaBoss struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
	Phase     int // Phase du combat
}

func InitTraumaBoss() TraumaBoss {
	return TraumaBoss{
		Name:      "Le Trauma Originel",
		MaxHP:     100,
		CurrentHP: 100,
		Attack:    15,
		Phase:     1,
	}
}

func FinalBossFight(player *character.Character) bool {
	boss := InitTraumaBoss()

	fmt.Println("💀 ═══ COMBAT FINAL ═══ 💀")
	fmt.Printf("Face à vous se dresse : %s\n", boss.Name)
	fmt.Println("Il représente tout ce qui vous a brisé...")

	for !player.IsDead() && boss.CurrentHP > 0 {
		// Phase du boss change selon ses PV
		if boss.CurrentHP <= 50 && boss.Phase == 1 {
			boss.Phase = 2
			boss.Attack = 20
			fmt.Println("💀 Le trauma révèle sa vraie forme ! Ses attaques se renforcent !")
		}

		// Tour du joueur (menu complet)
		fmt.Printf("\n⚔️ Boss PV: %d/%d | Vos PV: %d/%d\n",
			boss.CurrentHP, boss.MaxHP, player.PvCurr, player.PvMax)

		damage := finalBossPlayerTurn(player, &boss)
		boss.CurrentHP -= damage

		if boss.CurrentHP <= 0 {
			fmt.Println("🌟 Le trauma se dissout dans la lumière...")
			fmt.Println("🌟 Vous avez libéré votre esprit !")
			return true
		}

		// Tour du boss
		finalBossAttack(&boss, player)

		if player.IsDead() {
			fmt.Println("💀 Vos forces vous abandonnent...")
			return false
		}
	}

	return false
}

func finalBossPlayerTurn(player *character.Character, boss *TraumaBoss) int {
	fmt.Println("\n--- Actions disponibles ---")
	fmt.Println("1. Attaque normale (10 dégâts)")
	fmt.Println("2. Sort puissant (25 dégâts, coûte 20 mana)")
	fmt.Println("3. Utiliser une potion")

	var choice int
	fmt.Print("Votre choix (1-3): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		return 10
	case 2:
		if player.ManaCurr >= 20 {
			player.ManaCurr -= 20
			return 25
		}
		fmt.Println("❌ Pas assez de mana !")
		return 0
	case 3:
		if player.CountItem("Potion de vie") > 0 {
			player.TakePot()
		} else {
			fmt.Println("❌ Aucune potion disponible !")
		}
		return 0
	default:
		fmt.Println("❌ Action invalide !")
		return 0
	}
}

func finalBossAttack(boss *TraumaBoss, player *character.Character) {
	damage := boss.Attack

	switch boss.Phase {
	case 1:
		fmt.Printf("💀 %s murmure vos échecs... (%d dégâts)\n", boss.Name, damage)
	case 2:
		fmt.Printf("💀 %s hurle votre désespoir... (%d dégâts)\n", boss.Name, damage)
	}

	player.PvCurr -= damage
	if player.PvCurr < 0 {
		player.PvCurr = 0
	}
}
