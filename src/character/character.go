package character

import (
	"fmt"
	"math"
	"strings"
	"time"
	"somnium/ui"
	"math/rand"
)

const (
	GoblinXP = 25 
)

// Type Character représente le personnage du joueur.
type Character struct {
	Name          string
	Race          string
	Class         string
	Level         int
	PvMax         int
	PvCurr        int
	Inventory     []string
	InventorySize int
	Money         int
	Skills        []string
	ManaMax       int
	ManaCurr      int
	Equipment     Equipment
	XPCurr        int
	XPMax         int
	XPUpgrades    int
	CurrentLayer  int // Niveau actuel dans le labyrinthe
	IsShielded    bool
	Initiative    int
	TrapResistance int 
	HasResurrectionStone bool
	Weapon        Weapon // ✅ Nouveau : système d'armes
	Achievements  []string // ✅ Nouveau : succès débloques
}

// Type Equipment représente l'équipement du personnage.
type Equipment struct {
	Head  string
	Chest string
	Feet  string
}

// ✅ NOUVEAU : Système d'armes
type Weapon struct {
	Name   string
	Damage int
	Type   string // "sword", "staff", "bow", "dagger"
}

// Armes disponibles
var Weapons = map[string]Weapon{
	"Dague rouillée": {Name: "Dague rouillée", Damage: 5, Type: "dagger"},
	"Épée de fer": {Name: "Épée de fer", Damage: 12, Type: "sword"},
	"Bâton magique": {Name: "Bâton magique", Damage: 8, Type: "staff"},
	"Arc spectral": {Name: "Arc spectral", Damage: 10, Type: "bow"},
	"Lame des cauchemars": {Name: "Lame des cauchemars", Damage: 20, Type: "sword"},
	"Sceptre du trauma": {Name: "Sceptre du trauma", Damage: 25, Type: "staff"},
}

// InitCharacter initialise un nouveau personnage avec des valeurs de base.
func InitCharacter(name, race, class string, pvMax, manaMax int) Character {
	return Character{
		Name:          name,
		Race:          race,
		Class:         class,
		Level:         1,
		PvMax:         pvMax,
		PvCurr:        int(math.Round(float64(pvMax) * 0.5)),
		Inventory:     []string{"Potion de vie", "Potion de vie", "Potion de vie", "Potion de mana", "Potion de mana"},
		InventorySize: 10,
		Money:         100,
		Skills:        []string{"Coup de poing","Soin","Bouclier"}, 
		ManaCurr:      manaMax,
		ManaMax:       manaMax,
		Equipment:     Equipment{},
		XPCurr:        0,
		XPMax:         100,
		XPUpgrades:    0, 
		CurrentLayer:  1,
		IsShielded:    false,
		Initiative:    0,
		TrapResistance: 0,
		HasResurrectionStone: false,
		Weapon:        Weapon{}, // Pas d'arme au début
		Achievements:  []string{}, // Pas de succès au début
	}
}

// DisplayInfo affiche les informations détaillées du personnage.
func (c *Character) DisplayInfo() {
	ui.PrintInfo("═══════════════ LABYRINTHE DES CAUCHEMARS ═══════════════")
	ui.PrintInfo(fmt.Sprintf("Esprit : %s, %s %s, errant entre les couches de conscience.",
		c.Name, c.Race, c.Class))
	ui.PrintInfo(fmt.Sprintf("Niveau de conscience : %d — chaque pas pourrait être le dernier.", c.Level))
	c.DisplayXPInfo()
	ui.PrintInfo(fmt.Sprintf("Vitalité : %d/%d — votre essence vacille entre existence et néant.", c.PvCurr, c.PvMax))
	ui.PrintInfo(fmt.Sprintf("Énergie magique : %d/%d — le flux onirique vous soutient.", c.ManaCurr, c.ManaMax))
	ui.PrintInfo(fmt.Sprintf("Fragments de mémoire : %d — précieux pour survivre.", c.Money))
	
	// ✅ Affichage de l'arme équipée
	if c.Weapon.Name != "" {
		ui.PrintInfo(fmt.Sprintf("Arme équipée : %s (+%d dégâts)", c.Weapon.Name, c.Weapon.Damage))
	} else {
		ui.PrintInfo("Aucune arme équipée — vos poings devront suffire.")
	}
	
	if len(c.Skills) > 0 {
		ui.PrintInfo(fmt.Sprintf("Talents de l'esprit éveillé : %v.", c.Skills))
	} else {
		ui.PrintInfo("Aucun talent conscient pour l'instant, le sommeil est encore lourd.")
	}
	if c.Equipment.Head == "" && c.Equipment.Chest == "" && c.Equipment.Feet == "" {
		ui.PrintInfo("Aucun artefact ne protège votre enveloppe spectrale.")
	} else {
		ui.PrintInfo(fmt.Sprintf("Équipements trouvés dans ce rêve — tête: %s, torse: %s, pieds: %s.",
			c.Equipment.Head, c.Equipment.Chest, c.Equipment.Feet))
	}
	if len(c.Inventory) == 0 {
		ui.PrintInfo("Le sac de votre esprit est vide, attendant les reliques des rêves futurs.")
	} else {
		ui.PrintInfo(fmt.Sprintf("Dans votre sac éthéré : %v.", c.Inventory))
	}
	
	// ✅ Affichage des succès
	if len(c.Achievements) > 0 {
		ui.PrintSuccess(fmt.Sprintf("🏆 Succès débloqués : %d", len(c.Achievements)))
	}
	
	if c.IsDead() {
		ui.PrintError("\n⚠️  Votre essence vacille… la mort dans le Labyrinthe est bien réelle !")
	} else {
		ui.PrintInfo("\nVotre esprit flotte dans l'obscurité, prêt pour le prochain niveau.")
	}

	ui.PrintInfo("══════════════════════════════════════════════════════════")
}

// IsDead vérifie si le personnage est mort (PV à 0 ou moins).
func (c *Character) IsDead() bool {
	return c.PvCurr <= 0
}

// Resurrect ressuscite le personnage avec la moitié de ses PV et Mana.
func (c *Character) Resurrect() {
	c.PvCurr = c.PvMax / 2
	c.ManaCurr = c.ManaMax / 2
	ui.PrintSuccess("Le personnage a été ressuscité !")
}

// PoisonEffect applique l'effet de poison au personnage.
func (c *Character) PoisonEffect() {
	for i := 0; i < 3; i++ {
		c.PvCurr -= 10
		if c.PvCurr < 0 {
			c.PvCurr = 0
		}
		ui.PrintError(fmt.Sprintf("Poison : -10 PV (%d/%d)", c.PvCurr, c.PvMax))
		time.Sleep(1 * time.Second)
		if c.IsDead() {
			ui.PrintError("⚠️ Votre esprit succombe au poison !")
			break
		}
	}
}

// DisplayInventory affiche le contenu de l'inventaire du personnage.
func (c *Character) DisplayInventory() {
	if len(c.Inventory) == 0 {
		ui.PrintInfo("Votre sac est vide.")
		return
	}
	ui.PrintInfo("Inventaire :")
	for i, item := range c.Inventory {
		// ✅ Indication si c'est une arme
		if weapon, isWeapon := Weapons[item]; isWeapon {
			ui.PrintInfo(fmt.Sprintf("%d) %s (⚔️ Arme: +%d dégâts)", i+1, item, weapon.Damage))
		} else {
			ui.PrintInfo(fmt.Sprintf("%d) %s", i+1, item))
		}
	}
}

// UseItem utilise un objet de l'inventaire du personnage.
func (c *Character) UseItem(itemName string) bool {
	for i, item := range c.Inventory {
		if strings.EqualFold(item, itemName) {
			switch strings.ToLower(itemName) {
			case "potion de vie":
				c.PvCurr += 20
				if c.PvCurr > c.PvMax {
					c.PvCurr = c.PvMax
				}
			case "potion de mana":
				c.ManaCurr += 10
				if c.ManaCurr > c.ManaMax {
					c.ManaCurr = c.ManaMax
				}
			default:
				ui.PrintInfo("Vous utilisez " + itemName)
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	ui.PrintError("Objet introuvable : " + itemName)
	return false
}

// ✅ NOUVEAU : Équiper une arme
func (c *Character) EquipWeapon(weaponName string) bool {
	weapon, exists := Weapons[weaponName]
	if !exists {
		ui.PrintError("❌ Cette arme n'existe pas !")
		return false
	}

	// Remettre l'ancienne arme dans l'inventaire si elle existe
	if c.Weapon.Name != "" {
		if !c.AddToInventory(c.Weapon.Name) {
			ui.PrintError("🎒 Inventaire plein ! Impossible de ranger l'ancienne arme.")
			return false
		}
		ui.PrintInfo(fmt.Sprintf("⚔️ %s rangée dans l'inventaire.", c.Weapon.Name))
	}

	// Équiper la nouvelle arme
	c.Weapon = weapon
	c.RemoveFromInventory(weaponName)
	ui.PrintSuccess(fmt.Sprintf("⚔️ %s équipée ! (+%d dégâts)", weapon.Name, weapon.Damage))
	return true
}

// ✅ NOUVEAU : Calculer les dégâts avec l'arme
func (c *Character) GetAttackDamage() int {
	baseDamage := 8 // Coup de poing de base
	if c.Weapon.Name != "" {
		return baseDamage + c.Weapon.Damage
	}
	return baseDamage
}

// CountItem compte le nombre d'instances d'un objet dans l'inventaire.
func (c *Character) CountItem(itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

// AddToInventory ajoute un objet à l'inventaire du personnage.
func (c *Character) AddToInventory(item string) bool {
	if len(c.Inventory) >= c.InventorySize {
		ui.PrintError("Inventaire plein!")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

// ✅ NOUVEAU : Système de succès
func (c *Character) UnlockAchievement(achievementID string) {
	// Vérifier si déjà débloqué
	for _, unlocked := range c.Achievements {
		if unlocked == achievementID {
			return
		}
	}
	
	c.Achievements = append(c.Achievements, achievementID)
	
	// Messages selon le succès
	switch achievementID {
	case "first_victory":
		ui.PrintSuccess("🏆 Succès débloqué : Premier Pas - Gagner votre premier combat !")
	case "collector":
		ui.PrintSuccess("🏆 Succès débloqué : Collectionneur - Obtenir 10 objets différents !")
	case "warrior":
		ui.PrintSuccess("🏆 Succès débloqué : Guerrier - Atteindre le niveau 5 !")
	case "explorer":
		ui.PrintSuccess("🏆 Succès débloqué : Explorateur - Visiter toutes les couches !")
	case "boss_slayer":
		ui.PrintSuccess("🏆 Succès débloqué : Tueur de Boss - Vaincre le trauma primordial !")
	case "survivor":
		ui.PrintSuccess("🏆 Succès débloqué : Survivant - Survivre à 10 combats sans mourir !")
	}
}

// TakeItem utilise un objet spécifique (potion de vie ou mana) et applique ses effets.
func (c *Character) TakeItem(itemName string) {
	switch strings.ToLower(itemName) {
	case "potion de vie":
		if c.PvCurr >= c.PvMax {
			ui.PrintError("💊 Vos PV sont déjà au maximum !")
			return
		}
		if c.CountItem("Potion de vie") > 0 {
			c.UseItem("Potion de vie")
			c.PvCurr += 30
			if c.PvCurr > c.PvMax {
				c.PvCurr = c.PvMax
			}
			ui.PrintSuccess(fmt.Sprintf("💖 Vous buvez une potion de vie. PV: %d/%d", c.PvCurr, c.PvMax))
		} else {
			ui.PrintError("❌ Aucune potion de vie disponible !")
		}
	case "potion de mana":
		if c.ManaCurr >= c.ManaMax {
			ui.PrintError("🔮 Votre énergie magique est déjà au maximum !")
			return
		}
		if c.CountItem("Potion de mana") > 0 {
			c.UseItem("Potion de mana")
			c.ManaCurr += 20
			if c.ManaCurr > c.ManaMax {
				c.ManaCurr = c.ManaMax
			}
			ui.PrintSuccess(fmt.Sprintf("✨ Vous buvez une potion de mana. Mana: %d/%d", c.ManaCurr, c.ManaMax))
		} else {
			ui.PrintError("❌ Aucune potion de mana disponible !")
		}
	default:
		ui.PrintError("❌ Cet objet ne peut pas être utilisé directement !")
	}
}

// GainXP fait gagner de l'expérience au personnage
func (c *Character) GainXP(amount int) {
	ui.PrintSuccess(fmt.Sprintf("✨ Vous gagnez %d points d'expérience !", amount))
	c.XPCurr += amount

	for c.CheckLevelUp() {
		c.LevelUp()
	}
}

// CheckLevelUp vérifie si le personnage a assez d'XP pour monter de niveau
func (c *Character) CheckLevelUp() bool {
	return c.XPCurr >= c.XPMax
}

// LevelUp fait monter le personnage de niveau
func (c *Character) LevelUp() {
	if !c.CheckLevelUp() {
		return
	}
	
	excessXP := c.XPCurr - c.XPMax

	c.Level++
	c.XPUpgrades++
	
	oldMaxHP := c.PvMax
	oldMaxMana := c.ManaMax

	c.PvMax += 20   
	c.ManaMax += 10 

	c.PvCurr = c.PvMax
	c.ManaCurr = c.ManaMax

	
	c.XPMax = c.CalculateXPNeeded(c.Level)
	c.XPCurr = excessXP 

	// Message de level up thématique 
	ui.PrintSuccess("\n🌟 ═══ ÉVEIL DE L'ESPRIT ═══ 🌟")
	ui.PrintSuccess(fmt.Sprintf("Votre conscience s'élève... Niveau %d atteint !", c.Level))
	ui.PrintSuccess(fmt.Sprintf("💖 Vitalité : %d → %d (+20)", oldMaxHP, c.PvMax))
	ui.PrintSuccess(fmt.Sprintf("🔮 Essence : %d → %d (+10)", oldMaxMana, c.ManaMax))
	ui.PrintSuccess("Votre esprit et corps sont restaurés par cette révélation !")

	if excessXP > 0 {
		ui.PrintInfo(fmt.Sprintf("📊 XP en excès reportée : %d/%d", c.XPCurr, c.XPMax))
	}
	ui.PrintInfo("═══════════════════════════════════════")
	
	if c.Level >= 5 {
		c.UnlockAchievement("warrior")
	}
}

// CalculateXPNeeded calcule l'XP nécessaire pour un niveau donné
func (c *Character) CalculateXPNeeded(level int) int {
	return level * 100
}

// DisplayXPInfo affiche les informations d'XP (utilitaire)
func (c *Character) DisplayXPInfo() {
	nextLevelXP := c.CalculateXPNeeded(c.Level + 1)
	ui.PrintInfo(fmt.Sprintf("📊 Expérience : %d/%d (Niveau %d)", c.XPCurr, nextLevelXP, c.Level))
	if c.Level < 10 { 
		remaining := nextLevelXP - c.XPCurr
		ui.PrintInfo(fmt.Sprintf("   Prochain niveau dans : %d XP", remaining))
	}
}

// GetName retourne le nom du personnage.
func (c *Character) GetName() string {
	return c.Name
}

// RollInitiative lance l'initiative pour le personnage en fonction de sa classe.
func (c *Character) RollInitiative() {
	// Seed déjà fait dans monster.go
	baseRoll := rand.Intn(20) + 1
	c.Initiative = baseRoll

	// Bonus selon la classe
	bonus := 0
	switch c.Class {
	case "Voleur":
		bonus = 3  // Voleur = rapide
	case "Mage":
		bonus = 1  // Mage = réactif
	case "Guerrier":
		bonus = 0  // Guerrier = normal
	case "Occultiste":
		bonus = 2  // Occultiste = intuition
	}
	
	c.Initiative = baseRoll + bonus
	ui.PrintInfo(fmt.Sprintf("🎲 %s lance l'initiative : %d (base) + %d (classe) = %d", 
		c.Name, baseRoll, bonus, c.Initiative))
}

// RestoreHealth restaure les PV du personnage à leur maximum.
func (c *Character) RestoreHealth() {
	c.PvCurr = c.PvMax
}