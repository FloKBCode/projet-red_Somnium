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

func CoupDePoing(caster *character.Character, target *monster.Monster) int {
	damage := 8 + caster.Attack
	target.CurrentHP -= damage
	if target.CurrentHP < 0 {
		target.CurrentHP = 0
	}
	fmt.Printf("%s utilise Coup de Poing sur %s et inflige %d dégâts !\n", caster.Name, target.Name, damage)
	return damage
}

func BouleDeFeu(caster *character.Character, target *monster.Monster) int {
	damage := 18 + caster.MagicPower
	target.CurrentHP -= damage
	if target.CurrentHP < 0 {
		target.CurrentHP = 0
	}
	fmt.Printf("%s lance Boule de Feu sur %s et inflige %d dégâts !\n", caster.Name, target.Name, damage)
	return damage
}

// ----------------- GESTION DU SPELLBOOK -----------------

func (c *character.Character) CanCastSpell(spellName string) bool {
	for _, spell := range c.SpellBook {
		if spell == spellName {
			return true
		}
	}
	return false
}

// Même si on ne gère pas le mana, on garde la fonction pour la compatibilité
func (c *character.Character) ConsumeMP(cost int) bool {
	// On suppose que le personnage peut toujours lancer le sort
	return true
}

// Apprend un sort si le personnage ne le connaît pas déjà
func (c *character.Character) LearnSpell(spellName string) bool {
	if c.CanCastSpell(spellName) {
		return false
	}
	c.SpellBook = append(c.SpellBook, spellName)
	fmt.Printf("%s apprend le sort %s !\n", c.Name, spellName)
	return true
}
