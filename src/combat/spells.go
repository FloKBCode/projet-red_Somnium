package combat

import (
	"fmt"
	"somnium/character"
)

// Spell représente un sort avec ses caractéristiques
type Spell struct {
	Name     string
	Damage   int
	ManaCost int
	Effect   string
}

// ----------------- SORTS -----------------

// CoupDePoing : attaque de base gratuite
func CoupDePoing(caster *character.Character, target *Monster) int {
	damage := 8
	target.PvCurr -= damage
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	fmt.Printf("%s utilise Coup de Poing sur %s et inflige %d dégâts !\n",
		caster.Name, target.Name, damage)
	return damage
}

// BouleDeFeu : sort qui consomme du mana
func BouleDeFeu(caster *character.Character, target *Monster) int {
	cost := 15
	damage := 18

	if !caster.ConsumeMP(cost) {
		fmt.Printf("%s n'a pas assez de mana pour lancer Boule de Feu !\n", caster.Name)
		return 0
	}

	target.PvCurr -= damage
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	fmt.Printf("%s lance Boule de Feu sur %s et inflige %d dégâts !\n",
		caster.Name, target.Name, damage)
	return damage
}

// Coûts en mana des sorts
var SpellCosts = map[string]int{
	"Coup de poing": 5,
	"Boule de feu":  15,
	"Soin":          10, // Mission 3
	"Bouclier":      8,  // Mission 3
}

// 🔋 Restaurer du mana
func RestoreMana(c *character.Character, amount int) {
	c.ManaCurr += amount
	if c.ManaCurr > c.ManaMax {
		c.ManaCurr = c.ManaMax
	}
	fmt.Printf("🔮 %s regagne %d mana ! (%d/%d)\n", c.Name, amount, c.ManaCurr, c.ManaMax)
}

// 🔮 Obtenir le coût d’un sort
func ManaCost(c *character.Character, spellName string) int {
	if cost, ok := SpellCosts[spellName]; ok {
		return cost
	}
	return 0 // Sort inconnu → pas de coût
}

// Vérifier si on peut lancer un sort
func CanCastSpell(c *character.Character, spellName string) bool {
	cost := ManaCost(c, spellName)
	return c.ManaCurr >= cost
}

// Consommer du mana
func ConsumeMana(c *character.Character, spellName string) bool {
	cost := ManaCost(c, spellName)
	if c.ManaCurr >= cost {
		c.ManaCurr -= cost
		return true
	}
	return false
}
