package combat

import (
	"fmt"
	"math/rand"
	"somnium/character"
	"time"
)

// Type Boss représente le boss 
type Boss struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
	Phase     int 	
}

// InitBoss initialise le boss
func InitBoss(name string) Boss {
	return Boss{
		Name:      name,
		MaxHP:     100,
		CurrentHP: 100,
		Attack:    15,
		Phase:     1,
	}
}

// Convertit Boss en Monster pour réutiliser Fight()
func BossToMonster(boss Boss) Monster {
	return Monster{
		Name:   boss.Name,
		PvMax:  boss.MaxHP,
		PvCurr: boss.CurrentHP,
		Attack: boss.Attack,
		Level:  10,
		Loot:   []string{},
		Phase:  boss.Phase,
	}
}

// StartFinalBossFight lance le combat final
func StartFinalBossFight(player *character.Character) bool {
	rand.Seed(time.Now().UnixNano())

	boss := InitBoss("Trauma")
	monster := BossToMonster(boss)

	fmt.Println("💀 ═══ COMBAT FINAL ═══ 💀")
	fmt.Printf("Face à vous se dresse : %s\n", boss.Name)
	fmt.Println("Il représente tout ce qui vous a brisé...")

	victory := Fight(player, &monster, false, true)

	if !victory {
		fmt.Println("💀 Vos forces vous abandonnent...")
		player.CurrentLayer = 1 // retour à la première couche
		return false
	}

	fmt.Println("🌟 Le trauma se dissout dans la lumière...")
	fmt.Println("🌟 Vous avez libéré votre esprit !")
	return true
}

// Fonction spéciale pour adapter l'attaque du boss avec phases
func monsterAttackPatternBoss(boss *Monster, player *character.Character, turn int) {
	damage := boss.Attack

	// Phase 2 si HP <= 50%
	if boss.PvCurr <= boss.PvMax/2 && boss.Phase == 1 {
		boss.Attack = int(float64(boss.Attack) * 1.33)
		boss.Phase = 2
		fmt.Println("💀 Le trauma révèle sa vraie forme ! Ses attaques se renforcent !")
	}

	// Attaque du boss
	player.PvCurr -= damage
	if player.PvCurr < 0 {
		player.PvCurr = 0
	}

	switch boss.Phase {
	case 1:
		fmt.Printf("💀 %s murmure vos échecs... (%d dégâts)\n", boss.Name, damage)
	case 2:
		fmt.Printf("💀 %s hurle votre désespoir... (%d dégâts)\n", boss.Name, damage)
	}
}

