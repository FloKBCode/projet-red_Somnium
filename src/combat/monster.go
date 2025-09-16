package combat

import (
	"fmt"
)

type Monster struct {
	Name      string
	Level     int
	MaxHP     int
	CurrentHP int
	Attack    int
	Loot      []string
}

func InitGoblin() Monster {
	return Monster{
		Name:      "Gobelin",
		Level:     1,
		MaxHP:     30,
		CurrentHP: 30,
		Attack:    5,
		Loot:      []string{"Dague rouillée", "Morceau de cuir"},
	}
}
func (m *Monster) AttackTarget(target *Character) {
	target.PvCurr -= m.Attack
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	fmt.Printf("%s attaque %s et inflige %d dégâts !\n", m.Name, target.Name, m.Attack)
}
func (m *Monster) IsDead() bool {
	return m.CurrentHP <= 0
}
func (m *Monster) DisplayInfo() {
	fmt.Println("══════════════ MONSTRE ══════════════")
	fmt.Printf("Nom : %s (Niveau %d)\n", m.Name, m.Level)
	fmt.Printf("Vie : %d/%d\n", m.CurrentHP, m.MaxHP)
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
