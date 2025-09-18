package combat

import (
	"fmt"
	"somnium/character"
	"somnium/ui"
	"math/rand"
	"time"
)

// Type Monster représente un monstre dans le jeu
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

// Attaque la cible et réduit ses PV
func (m *Monster) AttackTarget(target *character.Character) {
	target.PvCurr -= m.Attack
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	ui.PrintError(fmt.Sprintf("%s attaque %s et inflige %d dégâts !", m.Name, target.Name, m.Attack))
}

// Vérifie si le monstre est mort
func (m *Monster) IsDead() bool {
	return m.PvCurr <= 0
}

// Affiche les informations du monstre
func (m *Monster) DisplayInfo() {
	ui.PrintError("══════════════ MONSTRE ══════════════")
	ui.PrintError(fmt.Sprintf("Nom : %s (Niveau %d)", m.Name, m.Level))
	ui.PrintError(fmt.Sprintf("Vie : %d/%d", m.PvCurr, m.PvMax))
	ui.PrintError(fmt.Sprintf("Attaque : %d", m.Attack))
	ui.PrintError(fmt.Sprintf("Initiative : %d", m.Initiative))

	if len(m.Loot) > 0 {
		ui.PrintInfo(fmt.Sprintf("Butin possible : %v", m.Loot))
	} else {
		ui.PrintInfo("Ce monstre ne laisse rien derrière lui...")
	}
	if m.IsDead() {
		ui.PrintError("⚰️  Ce monstre est mort.")
	}
	ui.PrintError("═════════════════════════════════════")
}

// ✅ RENFORCÉ : GenerateDefaultMonster plus fort
func GenerateDefaultMonster(level int, difficulty Difficulty) Monster {
	baseHP := 45 + (level * 20)    // ✅ AUGMENTÉ : 30→45 base, 10→20 par niveau
	baseAttack := 8 + (level * 4)  // ✅ AUGMENTÉ : 5→8 base, 2→4 par niveau

	monster := Monster{
		Name:   "Démon Intérieur",
		Level:  level,
		PvMax:  baseHP,
		PvCurr: baseHP,
		Attack: baseAttack,
		Loot:   []string{"Essence spirituelle", "Fragment d'âme", "Dague rouillée"},
		Phase:  1,
	}

	AdjustMonsterStats(&monster, difficulty)
	return monster
}

// RollInitiative génère un score d'initiative aléatoire
func RollInitiative() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1 // Random 1-20
}

// ✅ RENFORCÉ : Gobelins plus forts
func InitWeakGoblin() Monster {
	return Monster{
		Name:       "Gobelin Craintif",
		Level:      1,
		PvMax:      35,     // ✅ AUGMENTÉ : 20→35
		PvCurr:     35,
		Attack:     6,      // ✅ AUGMENTÉ : 3→6
		Loot:       []string{"Petit fragment", "Poussière d'espoir", "Dague rouillée"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitStrongGoblin() Monster {
	return Monster{
		Name:       "Gobelin des Abysses",
		Level:      3,
		PvMax:      80,     // ✅ AUGMENTÉ : 50→80
		PvCurr:     80,
		Attack:     15,     // ✅ AUGMENTÉ : 8→15
		Loot:       []string{"Fragment sombre", "Essence de peur", "Épée de fer"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitNightmareSpider() Monster {
	return Monster{
		Name:       "Araignée du Cauchemar",
		Level:      2,
		PvMax:      55,     // ✅ AUGMENTÉ : 35→55
		PvCurr:     55,
		Attack:     12,     // ✅ AUGMENTÉ : 6→12
		Loot:       []string{"Soie empoisonnée", "Crocs veineux", "Arc spectral"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitShadowWraith() Monster {
	return Monster{
		Name:       "Spectre des Ombres",
		Level:      4,
		PvMax:      100,    // ✅ AUGMENTÉ : 60→100
		PvCurr:     100,
		Attack:     20,     // ✅ AUGMENTÉ : 12→20
		Loot:       []string{"Essence spectrale", "Voile d'ombre", "Bâton magique"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitMemoryPhantom() Monster {
	return Monster{
		Name:       "Fantôme des Souvenirs",
		Level:      3,
		PvMax:      70,     // ✅ AUGMENTÉ : 40→70
		PvCurr:     70,
		Attack:     14,     // ✅ AUGMENTÉ : 7→14
		Loot:       []string{"Fragment de mémoire", "Larme cristalline", "Lame des cauchemars"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitGoblin() Monster {
	return Monster{
		Name:   "Gobelin",
		Level:  1,
		PvMax:  50,    // ✅ AUGMENTÉ : 30→50
		PvCurr: 50,
		Attack: 10,    // ✅ AUGMENTÉ : 5→10
		Loot:   []string{"Dague rouillée", "Morceau de cuir", "Fragment d'âme"},
		Phase:  1,
		Initiative: RollInitiative(),
	}
}

// ✅ NOUVEAU : Monstres de boss intermédiaires
func InitMiniBoss() Monster {
	return Monster{
		Name:       "Gardien des Regrets",
		Level:      5,
		PvMax:      150,
		PvCurr:     150,
		Attack:     25,
		Loot:       []string{"Sceptre du trauma", "Cristal de puissance", "Armure spectrale"},
		Phase:      1,
		Initiative: RollInitiative(),
	}
}

func InitNightmareBeast() Monster {
	return Monster{
		Name:       "Bête du Cauchemar",
		Level:      6,
		PvMax:      200,
		PvCurr:     200,
		Attack:     30,
		Loot:       []string{"Griffe de la peur", "Cœur de ténèbres", "Lame maudite"},
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

	// ✅ Ajouter mini-boss pour niveaux élevés
	if level >= 4 {
		monsters = append(monsters, InitMiniBoss)
	}
	if level >= 5 {
		monsters = append(monsters, InitNightmareBeast)
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

// ✅ RENFORCÉ : GenerateBoss encore plus puissant
func GenerateBoss(level int) Monster {
	monster := GenerateDefaultMonster(level+3, Hard) // ✅ +3 niveaux au lieu de +2
	monster.Name = "Trauma Primordial"
	monster.PvMax *= 3    // ✅ x3 au lieu de x2
	monster.PvCurr = monster.PvMax
	monster.Attack = int(float64(monster.Attack) * 2.0) // ✅ x2 au lieu de x1.5
	monster.Loot = append(monster.Loot, "Cristal de Trauma", "Couronne de la Libération")
	return monster
}