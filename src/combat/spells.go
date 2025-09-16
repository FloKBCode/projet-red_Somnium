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

// CoupDePoing : attaque de base
func CoupDePoing(caster *character.Character, target *Monster) int {
	damage := 8
	target.CurrentHP -= damage
	if target.CurrentHP < 0 {
		target.CurrentHP = 0
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

	target.CurrentHP -= damage
	if target.CurrentHP < 0 {
		target.CurrentHP = 0
	}
	fmt.Printf("%s lance Boule de Feu sur %s et inflige %d dégâts !\n",
		caster.Name, target.Name, damage)
	return damage
}

// Vérifie si le personnage connaît un sort
func (c *character.Character) CanCastSpell(spellName string) bool {
	for _, skill := range c.Skills {
		if skill == spellName {
			return true
		}
	}
	return false
}

// Consomme le mana si possible
func (c *character.Character) ConsumeMP(cost int) bool {
	if c.ManaCurr < cost {
		return false
	}
	c.ManaCurr -= cost
	return true
}

// Ajoute un sort si pas déjà appris
func (c *character.Character) LearnSpell(spellName string) bool {
	if c.CanCastSpell(spellName) {
		return false
	}
	c.Skills = append(c.Skills, spellName)
	fmt.Printf("%s apprend le sort %s !\n", c.Name, spellName)
	return true
}
