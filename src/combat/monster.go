package combat

import (
	"fmt"
	"somnium/character"
	"math/rand"
	"time"
)

type Monster struct {
	Name   string
	Level  int
	PvMax  int
	PvCurr int
	Attack int
	Loot   []string
	Phase  int
	Initiative int
}

func InitGoblin() Monster {
	return Monster{
		Name:   "Gobelin",
		Level:  1,
		PvMax:  30,
		PvCurr: 30,
		Attack: 5,
		Loot:   []string{"Dague rouillée", "Morceau de cuir"},
		Phase:  1,
		Initiative: RollInitiative(),
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
	fmt.Printf("Initiative : %d\n", m.Initiative)

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

// GenerateDefaultMonster crée un nouveau monstre adapté au niveau
func GenerateDefaultMonster(level int, difficulty Difficulty) Monster {
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
	monster := GenerateDefaultMonster(level+2, Hard)
	monster.Name = "Trauma Primordial"
	monster.PvMax *= 2
	monster.PvCurr = monster.PvMax
	monster.Attack = int(float64(monster.Attack) * 1.5)
	monster.Loot = append(monster.Loot, "Cristal de Trauma")
	return monster
}

func RollInitiative() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1 // Random 1-20
}

func InitWeakGoblin() Monster {
	return Monster{
		Name:       "Gobelin Craintif",
		Level:      1,
		PvMax:      20,
		PvCurr:     20,
		Attack:     3,
		Loot:       []string{"Petit fragment", "Poussière d'espoir"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitStrongGoblin() Monster {
	return Monster{
		Name:       "Gobelin des Abysses",
		Level:      3,
		PvMax:      50,
		PvCurr:     50,
		Attack:     8,
		Loot:       []string{"Fragment sombre", "Essence de peur"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitNightmareSpider() Monster {
	return Monster{
		Name:       "Araignée du Cauchemar",
		Level:      2,
		PvMax:      35,
		PvCurr:     35,
		Attack:     6,
		Loot:       []string{"Soie empoisonnée", "Crocs veineux"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitShadowWraith() Monster {
	return Monster{
		Name:       "Spectre des Ombres",
		Level:      4,
		PvMax:      60,
		PvCurr:     60,
		Attack:     12,
		Loot:       []string{"Essence spectrale", "Voile d'ombre"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitMemoryPhantom() Monster {
	return Monster{
		Name:       "Fantôme des Souvenirs",
		Level:      3,
		PvMax:      40,
		PvCurr:     40,
		Attack:     7,
		Loot:       []string{"Fragment de mémoire", "Larme cristalline"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

// ✅ AMÉLIORER GenerateMonster pour plus de variété
func GenerateMonster(level int, difficulty Difficulty) Monster {
	monsters := []func() Monster{
		InitGoblin,
		InitWeakGoblin,
		InitStrongGoblin,
		InitNightmareSpider,
		InitShadowWraith,
		InitMemoryPhantom,
	}

	// Choisir un monstre aléatoire selon le niveau
	var availableMonsters []func() Monster
	for _, monsterFunc := range monsters {
		testMonster := monsterFunc()
		if testMonster.Level <= level+1 { // Monstres adaptés au niveau
			availableMonsters = append(availableMonsters, monsterFunc)
		}
	}

	if len(availableMonsters) == 0 {
		availableMonsters = []func() Monster{InitGoblin} // Fallback
	}

	// Générer monstre aléatoire
	rand.Seed(time.Now().UnixNano())
	selectedMonster := availableMonsters[rand.Intn(len(availableMonsters))]()
	
	AdjustMonsterStats(&selectedMonster, difficulty)
	return selectedMonster
}