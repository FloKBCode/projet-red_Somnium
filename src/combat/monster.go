package combat

import (
	"fmt"
	"somnium/character"
)

type Monster struct {
	Name   string
	Level  int
	PvMax  int
	PvCurr int
	Attack int
	Loot   []string
	Phase  int
}

func InitGoblin() Monster {
	return Monster{
		Name:   "Gobelin",
		Level:  1,
		PvMax:  30,
		PvCurr: 30,
		Attack: 5,
		Loot:   []string{"Dague rouillée", "Morceau de cuir"},
	}
}
func (m *Monster) AttackTarget(target *character.Character) {
	target.PvCurr -= m.Attack
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	fmt.Printf("%s attaque %s et inflige %d dégâts !\n", m.Name, target.Name, m.Attack)
}
func (m *Monster) IsDead() bool {
	return m.PvCurr <= 0
}
func (m *Monster) DisplayInfo() {
	fmt.Println("══════════════ MONSTRE ══════════════")
	fmt.Printf("Nom : %s (Niveau %d)\n", m.Name, m.Level)
	fmt.Printf("Vie : %d/%d\n", m.PvCurr, m.PvMax)
	fmt.Printf("Attaque : %d\n", m.Attack)

	if len(m.Loot) > 0 {
		fmt.Printf("Butin possible : %v\n", m.Loot)
	} else {
		fmt.Println("Ce monstre ne laisse rien derrière lui...")
	}
	if m.IsDead() {
		fmt.Println("⚰️  Ce monstre est mort.")
	}
	fmt.Println("═════════════════════════════════════")
}

// GenerateMonster crée un nouveau monstre adapté au niveau
func GenerateMonster(level int, difficulty Difficulty) Monster {
	baseHP := 30 + (level * 10)
	baseAttack := 5 + (level * 2)

	monster := Monster{
		Name:   "Démon Intérieur",
		Level:  level,
		PvMax:  baseHP,
		PvCurr: baseHP,
		Attack: baseAttack,
		Loot:   []string{"Essence spirituelle", "Fragment d'âme"},
		Phase:  1,
	}

	AdjustMonsterStats(&monster, difficulty)
	return monster
}

// GenerateBoss crée un boss puissant
func GenerateBoss(level int) Monster {
	monster := GenerateMonster(level+2, Hard)
	monster.Name = "Trauma Primordial"
	monster.PvMax *= 2
	monster.PvCurr = monster.PvMax
	monster.Attack = int(float64(monster.Attack) * 1.5)
	monster.Loot = append(monster.Loot, "Cristal de Trauma")
	return monster
}
