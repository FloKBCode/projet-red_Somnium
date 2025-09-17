package combat

import (
	"fmt"
	"math/rand"
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
	// Chance de coup critique
	if rand.Intn(100) < 15 {
		damage *= 2
		fmt.Println("💥 Coup critique !")
	}
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

	// Coup critique possible
	if rand.Intn(100) < 15 {
		damage *= 2
		fmt.Println("💥 Coup critique magique !")
	}

	target.PvCurr -= damage
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	fmt.Printf("%s lance Boule de Feu sur %s et inflige %d dégâts !\n",
		caster.Name, target.Name, damage)
	return damage
}

// Soin : restaure des PV
func Heal(caster *character.Character) {
	cost := 10
	heal := 20

	if !caster.ConsumeMP(cost) {
		fmt.Printf("%s n'a pas assez de mana pour se soigner !\n", caster.Name)
		return
	}

	caster.PvCurr += heal
	if caster.PvCurr > caster.PvMax {
		caster.PvCurr = caster.PvMax
	}
	fmt.Printf("✨ %s se soigne et regagne %d PV (%d/%d)!\n",
		caster.Name, heal, caster.PvCurr, caster.PvMax)
}

// Bouclier : réduit les dégâts subis pendant 1 tour
func Shield(caster *character.Character) {
	cost := 8
	if !caster.ConsumeMP(cost) {
		fmt.Printf("%s n'a pas assez de mana pour activer Bouclier !\n", caster.Name)
		return
	}
	caster.IsShielded = true
	fmt.Printf("🛡️ %s se protège avec un bouclier magique pour ce tour !\n", caster.Name)
}

// ----------------- OUTILS -----------------

// 🔋 Restaurer du mana
func RestoreMana(c *character.Character, amount int) {
	c.ManaCurr += amount
	if c.ManaCurr > c.ManaMax {
		c.ManaCurr = c.ManaMax
	}
	fmt.Printf("🔮 %s regagne %d mana ! (%d/%d)\n", c.Name, amount, c.ManaCurr, c.ManaMax)
}

func ConsumeMana(c *character.Character, spellName string) bool {
	cost := ManaCost(c, spellName)
	if c.ManaCurr >= cost {
		c.ManaCurr -= cost
		return true
	}
	return false
}

func ManaCost(c *character.Character, spellName string) int {
	if cost, ok := SpellCosts[spellName]; ok {
		return cost
	}
	return 0 // Sort inconnu 
	// → pas de coût
}

var SpellCosts = map[string]int{
	"Coup de poing": 5,
	"Boule de feu":  15,
	"Soin":          10, // Mission 3
	"Bouclier":      8,  // Mission 3
}