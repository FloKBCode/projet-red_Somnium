package combat

import (
	"fmt"
	"somnium/character"
)

// État d’un combat
type CombatState struct {
	Turn         int
	PlayerAlive  bool
	MonsterAlive bool
}

// Combat d’entrainement contre un gobelin
func TrainingFight(player *character.Character) {
	goblin := InitGoblin()
	state := CombatState{
		Turn:         1,
		PlayerAlive:  true,
		MonsterAlive: true,
	}

	fmt.Println("⚔️ Début du combat d'entraînement contre un Gobelin !")

	for state.PlayerAlive && state.MonsterAlive {
		fmt.Printf("\n=== Tour %d ===\n", state.Turn)

		// Tour du joueur
		state.MonsterAlive = CharacterTurn(player, &goblin, state.Turn)
		if !state.MonsterAlive {
			fmt.Println("🎉 Le joueur a vaincu le Gobelin !")
			break
		}

		// Tour du gobelin
		GoblinPattern(&goblin, player, state.Turn)
		if player.IsDead() {
			state.PlayerAlive = false
			fmt.Println("💀 Le joueur a été vaincu par le Gobelin...")
			break
		}

		state.Turn++
	}
}

// Tour du joueur
func CharacterTurn(player *character.Character, monster *Monster, turn int) bool {
	// Pour l’instant, attaque de base → Coup de Poing
	damage := CoupDePoing(player, monster)

	fmt.Printf("%s attaque le Gobelin (Tour %d) et inflige %d dégâts. Vie restante du Gobelin: %d/%d\n",
		player.Name, turn, damage, monster.CurrentHP, monster.MaxHP)

	// Vérifier si le monstre est mort
	return !monster.IsDead()
}

// Pattern d’attaque du gobelin
func GoblinPattern(goblin *Monster, player *character.Character, turn int) {
	damage := goblin.Attack

	// Tous les 3 tours → attaque renforcée
	if turn%3 == 0 {
		damage *= 2
		fmt.Println("⚡ Le Gobelin concentre ses forces pour une attaque puissante !")
	}

	player.PvCurr -= damage
	if player.PvCurr < 0 {
		player.PvCurr = 0
	}

	fmt.Printf("👹 %s attaque %s et inflige %d dégâts ! (%d/%d PV restants)\n",
		goblin.Name, player.Name, damage, player.PvCurr, player.PvMax)
}
