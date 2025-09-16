package character

import "fmt"
// ----------------- MÉTHODES PERSONNAGE -----------------

// Méthodes liées à la magie / sorts sur le type Character.
// Ces méthodes doivent être définies dans le package 'character'.

// CanCastSpell vérifie si le personnage connaît le sort.
func (c *Character) CanCastSpell(spellName string) bool {
	for _, skill := range c.Skills {
		if skill == spellName {
			return true
		}
	}
	return false
}

// ConsumeMP consomme du mana si disponible.
// IMPORTANT : receiver pointeur pour modifier le personnage original.
func (c *Character) ConsumeMP(cost int) bool {
	if c.ManaCurr < cost {
		return false
	}
	c.ManaCurr -= cost
	return true
}

// LearnSpell ajoute un sort aux compétences si pas déjà appris.
func (c *Character) LearnSpell(spellName string) bool {
	if c.CanCastSpell(spellName) {
		return false
	}
	c.Skills = append(c.Skills, spellName)
	fmt.Printf("%s apprend le sort %s !\n", c.Name, spellName)
	return true
}
