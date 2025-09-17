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
	goblin.DisplayInfo() // Afficher infos monstre

	for state.PlayerAlive && state.MonsterAlive {
		fmt.Printf("\n=== Tour %d ===\n", state.Turn)

		// Tour du joueur avec menu
		state.MonsterAlive = CharacterTurn(player, &goblin, state.Turn)
		if !state.MonsterAlive {
			fmt.Println("🎉 Le joueur a vaincu le Gobelin !")
			
			// ✅ GAIN D'XP À LA VICTOIRE
			player.GainXP(25) // GoblinXP = 25
			
			// Gestion de la mort
			if player.IsDead() {
				fmt.Println("💀 Mais vous succombez aussi à vos blessures...")
				player.Resurrect()
			}
			break
		}

		// Tour du gobelin
		GoblinPattern(&goblin, player, state.Turn)
		if player.IsDead() {
			state.PlayerAlive = false
			fmt.Println("💀 Le joueur a été vaincu par le Gobelin...")
			player.Resurrect() // Auto-résurrection
			break
		}

		state.Turn++
	}
}

// Tour du joueur
func CharacterTurn(player *character.Character, monster *Monster, turn int) bool {
	fmt.Printf("\n⚔️ C'est votre tour, %s !\n", player.Name)
	fmt.Printf("💖 Vos PV : %d/%d | 🔮 Mana : %d/%d\n", 
		player.PvCurr, player.PvMax, player.ManaCurr, player.ManaMax)
	
	fmt.Println("1. Attaquer (Coup de Poing - gratuit)")
	if player.CanCastSpell("Boule de feu") && player.ManaCurr >= 15 {
		fmt.Println("2. Boule de Feu (18 dégâts - 15 mana)")
	}
	fmt.Println("3. Utiliser inventaire")
	fmt.Println("4. Fuir le combat")
	
	var choice int
	fmt.Print("👉 Votre choix : ")
	fmt.Scanln(&choice)
	
	switch choice {
	case 1:
		damage := CoupDePoing(player, monster)
		fmt.Printf("💥 %s inflige %d dégâts ! (%d/%d PV restants)\n", 
			player.Name, damage, monster.CurrentHP, monster.MaxHP)
	
	case 2:
		if player.CanCastSpell("Boule de feu") && player.ManaCurr >= 15 {
			damage := BouleDeFeu(player, monster)
			if damage > 0 {
				fmt.Printf("🔥 %s inflige %d dégâts magiques ! (%d/%d PV restants)\n", 
					player.Name, damage, monster.CurrentHP, monster.MaxHP)
			}
		} else {
			fmt.Println("❌ Sort indisponible ou pas assez de mana !")
			return true // Le tour continue
		}
	
	case 3:
		fmt.Println("🎒 Accès à l'inventaire...")
		// Utilisation simple d'une potion
		if player.CountItem("Potion de vie") > 0 {
			player.TakePot()
		} else {
			fmt.Println("❌ Aucune potion disponible !")
		}
	
	case 4:
		fmt.Println("💨 Vous fuyez le combat...")
		return false // Monstre "survit", combat terminé
	
	default:
		fmt.Println("❌ Action invalide, vous perdez votre tour !")
	}
	
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

