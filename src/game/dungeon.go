package game


import (
	"fmt"
	"somnium/character"
)


func (c *character.Character) GainXP(amount int) {
	c.XPCurr += amount
	fmt.Printf("%s gagne %d points d’expérience. (Total: %d)\n", c.Name, amount, c.XPCurr)

	// Vérifie si un level up est possible
	for c.CheckLevelUp() {
		c.LevelUp()
	}
}

// CheckLevelUp renvoie true si le personnage peut monter de niveau.
func (c *character.Character) CheckLevelUp() bool {
	xpNeeded := c.CalculateXPNeeded(c.Level)
	return c.XPCurr >= xpNeeded
}

// LevelUp fait passer le personnage au niveau suivant et augmente ses stats.
func (c *character.Character) LevelUp() {
	c.Level++
	c.XPCurr -= c.CalculateXPNeeded(c.Level - 1)

	// Bonus de stats à chaque level up
	c.PvMax += 20
	c.ManaMax += 10

	// Restaure les PV/Mana au max après level up
	c.PvCurr = c.PvMax
	c.ManaCurr = c.ManaMax

	fmt.Printf("✨ %s passe au niveau %d !\n", c.Name, c.Level)
	fmt.Printf("Stats → PV: %d | Mana: %d | STR: %d | AGI: %d | INT: %d\n",
		c.PvMax, c.ManaMax)
}

// CalculateXPNeeded calcule l’XP nécessaire pour passer du level donné au suivant.
func (c *character.Character) CalculateXPNeeded(level int) int {
	return level * 100
}


