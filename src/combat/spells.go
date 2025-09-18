package combat

import (
	"fmt"
	"math/rand"
	"somnium/character"
	"somnium/ui"
)

// Spell représente un sort avec ses caractéristiques
type Spell struct {
	Name     string
	Damage   int
	ManaCost int
	Effect   string
}

// ✅ AMÉLIORÉ : CoupDePoing avec dégâts d'arme
func CoupDePoing(caster *character.Character, target *Monster) int {
	damage := caster.GetAttackDamage() // ✅ Utilise les dégâts avec arme
	
	if rand.Intn(100) < 15 {
		damage *= 2
		ui.PrintSuccess("💥 Coup critique !")
	}
	
	target.PvCurr -= damage
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	
	// Message personnalisé selon l'arme
	if caster.Weapon.Name != "" {
		ui.PrintInfo(fmt.Sprintf("%s attaque avec %s et inflige %d dégâts !",
			caster.Name, caster.Weapon.Name, damage))
	} else {
		ui.PrintInfo(fmt.Sprintf("%s utilise Coup de Poing et inflige %d dégâts !",
			caster.Name, damage))
	}
	
	return damage
}

// BouleDeFeu : attaque magique à distance
func BouleDeFeu(caster *character.Character, target *Monster) int {
	cost := 15
	damage := 18

	if !caster.ConsumeMP(cost) {
		ui.PrintError(fmt.Sprintf("%s n'a pas assez de mana pour lancer Boule de Feu !", caster.Name))
		return 0
	}

	if rand.Intn(100) < 15 {
		damage *= 2
		ui.PrintSuccess("💥 Coup critique magique !")
	}

	target.PvCurr -= damage
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	ui.PrintInfo(fmt.Sprintf("%s lance Boule de Feu sur %s et inflige %d dégâts !",
		caster.Name, target.Name, damage))
	return damage
}

// Soin : restaure des PV
func Heal(caster *character.Character) {
	cost := 10
	heal := 20

	if !caster.ConsumeMP(cost) {
		ui.PrintError(fmt.Sprintf("%s n'a pas assez de mana pour se soigner !", caster.Name))
		return
	}

	caster.PvCurr += heal
	if caster.PvCurr > caster.PvMax {
		caster.PvCurr = caster.PvMax
	}
	ui.PrintSuccess(fmt.Sprintf("✨ %s se soigne et regagne %d PV (%d/%d)!",
		caster.Name, heal, caster.PvCurr, caster.PvMax))
}

// Bouclier : réduit les dégâts subis pendant 1 tour
func Shield(caster *character.Character) {
	cost := 8
	if !caster.ConsumeMP(cost) {
		ui.PrintError(fmt.Sprintf("%s n'a pas assez de mana pour activer Bouclier !", caster.Name))
		return
	}
	caster.IsShielded = true
	ui.PrintSuccess(fmt.Sprintf("🛡️ %s se protège avec un bouclier magique pour ce tour !", caster.Name))
}

// 🔋 Restaurer du mana
func RestoreMana(c *character.Character, amount int) {
	c.ManaCurr += amount
	if c.ManaCurr > c.ManaMax {
		c.ManaCurr = c.ManaMax
	}
	ui.PrintSuccess(fmt.Sprintf("🔮 %s regagne %d mana ! (%d/%d)", c.Name, amount, c.ManaCurr, c.ManaMax))
}

// Consomme le mana pour un sort, retourne vrai si réussi
func ConsumeMana(c *character.Character, spellName string) bool {
	cost := ManaCost(c, spellName)
	if c.ManaCurr >= cost {
		c.ManaCurr -= cost
		return true
	}
	return false
}

// Récupère le coût en mana d'un sort
func ManaCost(c *character.Character, spellName string) int {
	if cost, ok := SpellCosts[spellName]; ok {
		return cost
	}
	return 0 // Sort inconnu → pas de coût
}

// Chaîne d'éclairs : attaque électrique avec chance de critique
func ChaineLightning(caster *character.Character, target *Monster) int {
	cost := 20
	damage := 25

	if !caster.ConsumeMP(cost) {
		ui.PrintError(fmt.Sprintf("%s n'a pas assez de mana pour lancer Chaîne d'éclairs !", caster.Name))
		return 0
	}

	if rand.Intn(100) < 20 { // 20% de critique
		damage = int(float64(damage) * 1.5)
		ui.PrintSuccess("⚡ Foudre dévastatrice !")
	}

	target.PvCurr -= damage
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	ui.PrintInfo(fmt.Sprintf("⚡ %s invoque une chaîne d'éclairs sur %s et inflige %d dégâts !",
		caster.Name, target.Name, damage))
	return damage
}

// Mur de glace : réduit les dégâts du prochain tour de 75%
func MurDeGlace(caster *character.Character) {
	cost := 18
	if !caster.ConsumeMP(cost) {
		ui.PrintError(fmt.Sprintf("%s n'a pas assez de mana pour créer un Mur de glace !", caster.Name))
		return
	}
	
	// Réduit les dégâts du prochain tour de 75%
	caster.IsShielded = true // Réutilise le système existant mais avec effet renforcé
	ui.PrintSuccess(fmt.Sprintf("🧊 %s érige un mur de glace protecteur !", caster.Name))
}

// Soin ++
func SoinMajeur(caster *character.Character) {
	cost := 25
	heal := 40

	if !caster.ConsumeMP(cost) {
		ui.PrintError(fmt.Sprintf("%s n'a pas assez de mana pour lancer Soin majeur !", caster.Name))
		return
	}

	caster.PvCurr += heal
	if caster.PvCurr > caster.PvMax {
		caster.PvCurr = caster.PvMax
	}
	ui.PrintSuccess(fmt.Sprintf("💚 %s se soigne puissamment et regagne %d PV (%d/%d)!",
		caster.Name, heal, caster.PvCurr, caster.PvMax))
}

// DraineSoul : récupère de la vie égale à la moitié des dégâts
func DraineSoul(caster *character.Character, target *Monster) int {
	cost := 12
	damage := 15

	if !caster.ConsumeMP(cost) {
		ui.PrintError(fmt.Sprintf("%s n'a pas assez de mana pour lancer Draine-âme !", caster.Name))
		return 0
	}

	target.PvCurr -= damage
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}
	
	heal := damage / 2
	caster.PvCurr += heal
	if caster.PvCurr > caster.PvMax {
		caster.PvCurr = caster.PvMax
	}

	ui.PrintInfo(fmt.Sprintf("🌀 %s draine l'essence de %s : %d dégâts et +%d PV récupérés !",
		caster.Name, target.Name, damage, heal))
	return damage
}

// ExplosionPsychique : sort puissant avec malus : coûte aussi de la vie
func ExplosionPsychique(caster *character.Character, target *Monster) int {
	cost := 35
	damage := 45

	if !caster.ConsumeMP(cost) {
		ui.PrintError(fmt.Sprintf("%s n'a pas assez de mana pour déclencher une Explosion psychique !", caster.Name))
		return 0
	}
	
	caster.PvCurr -= 5
	if caster.PvCurr < 1 {
		caster.PvCurr = 1
	}

	target.PvCurr -= damage
	if target.PvCurr < 0 {
		target.PvCurr = 0
	}

	ui.PrintError(fmt.Sprintf("💥 %s libère une explosion psychique dévastatrice sur %s ! (%d dégâts, -5 PV pour le lanceur)",
		caster.Name, target.Name, damage))
	return damage
}

// Menu des sorts
func SpellMenu(player *character.Character, monster *Monster, state *CombatState) {
	availableSpells := getAvailableSpells(player)
	
	if len(availableSpells) == 0 {
		ui.PrintError("❌ Vous ne connaissez aucun sort !")
		return
	}

	ui.PrintInfo("\n--- Sorts disponibles ---")
	for i, spell := range availableSpells {
		cost := SpellCosts[spell]
		manaStatus := "✅"
		if player.ManaCurr < cost {
			manaStatus = "❌"
		}
		ui.PrintInfo(fmt.Sprintf("%d. %s (%d mana) %s", i+1, spell, cost, manaStatus))
	}

	var spellChoice int
	ui.PrintInfo("👉 Choix du sort : ")
	fmt.Scanln(&spellChoice)

	if spellChoice < 1 || spellChoice > len(availableSpells) {
		ui.PrintError("❌ Sort invalide.")
		return
	}

	selectedSpell := availableSpells[spellChoice-1]
	castSpell(player, monster, selectedSpell, state)
}

func getAvailableSpells(player *character.Character) []string {
	var available []string
	for _, skill := range player.Skills {
		if _, exists := SpellCosts[skill]; exists {
			available = append(available, skill)
		}
	}
	return available
}

func castSpell(player *character.Character, monster *Monster, spellName string, state *CombatState) {
	switch spellName {
	case "Boule de feu":
		BouleDeFeu(player, monster)
	case "Soin":
		if ConsumeMana(player, "Soin") {
			heal := 20
			player.PvCurr += heal
			if player.PvCurr > player.PvMax {
				player.PvCurr = player.PvMax
			}
			ui.PrintSuccess(fmt.Sprintf("💖 %s se soigne de %d PV (%d/%d)", 
				player.Name, heal, player.PvCurr, player.PvMax))
		}
	case "Bouclier":
		if ConsumeMana(player, "Bouclier") {
			state.ShieldTurns = 3
			ui.PrintSuccess(fmt.Sprintf("🛡️ %s active un bouclier pour 3 tours !", player.Name))
		}
	case "Chaîne d'éclairs":
		ChaineLightning(player, monster)
	case "Mur de glace":
		MurDeGlace(player)
	case "Soin majeur":
		SoinMajeur(player)
	case "Draine-âme":
		DraineSoul(player, monster)
	case "Explosion psychique":
		ExplosionPsychique(player, monster)
	default:
		ui.PrintError("❌ Sort non implémenté : " + spellName)
	}
}

var SpellCosts = map[string]int{
	"Coup de poing":      0,   
	"Boule de feu":       15,
	"Soin":               10,
	"Bouclier":           8,
	"Chaîne d'éclairs":   20,
	"Mur de glace":       18,
	"Soin majeur":        25,
	"Régénération":       30,
	"Explosion psychique": 35,
	"Draine-âme":         12,
}