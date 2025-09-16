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

