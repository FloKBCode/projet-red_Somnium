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
	GoblinXP = 25  // Récompense XP pour tuer un gobelin
)

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
	CurrentLayer  int
	IsShielded    bool
	Initiative    int
}

type Equipment struct {
	Head  string
	Chest string
	Feet  string
}

func InitCharacter(name, race, class string, pvMax, manaMax int) Character {
	return Character{
		Name:          name,
		Race:          race,
		Class:         class,
		Level:         1,
		PvMax:         pvMax,
		PvCurr:        int(math.Round(float64(pvMax) * 0.5)),
		Inventory:     []string{"Potion de vie", "Potion de vie", "Potion de vie"},
		InventorySize: 10,
		Money:         100,
		Skills:        []string{"Coup de poing"},
		ManaCurr:      manaMax,
		ManaMax:       manaMax,
		Equipment:     Equipment{},
		XPCurr:        0,
		XPMax:         100,
		XPUpgrades:    0,
		CurrentLayer:  1,
		IsShielded:    false,
		Initiative:    0,
	}
}

func (c *Character) DisplayInfo() {
	fmt.Println("═══════════════ LABYRINTHE DES CAUCHEMARS ═══════════════")
	fmt.Printf("Esprit : %s, %s %s, errant entre les couches de conscience.\n",
		c.Name, c.Race, c.Class)
	fmt.Printf("Niveau de conscience : %d — chaque pas pourrait être le dernier.\n", c.Level)
	c.DisplayXPInfo()
	// HP et Mana
	fmt.Printf("Vitalité : %d/%d — votre essence vacille entre existence et néant.\n", c.PvCurr, c.PvMax)
	fmt.Printf("Énergie magique : %d/%d — le flux onirique vous soutient.\n", c.ManaCurr, c.ManaMax)

	// Argent / ressources
	fmt.Printf("Fragments de mémoire : %d — précieux pour survivre.\n", c.Money)

	// Compétences
	if len(c.Skills) > 0 {
		fmt.Printf("Talents de l’esprit éveillé : %v.\n", c.Skills)
	} else {
		fmt.Println("Aucun talent conscient pour l’instant, le sommeil est encore lourd.")
	}

	// Équipement
	if c.Equipment.Head == "" && c.Equipment.Chest == "" && c.Equipment.Feet == "" {
		fmt.Println("Aucun artefact ne protège votre enveloppe spectrale.")
	} else {
		fmt.Printf("Équipements trouvés dans ce rêve — tête: %s, torse: %s, pieds: %s.\n",
			c.Equipment.Head, c.Equipment.Chest, c.Equipment.Feet)
	}

	// Inventaire
	if len(c.Inventory) == 0 {
		fmt.Println("Le sac de votre esprit est vide, attendant les reliques des rêves futurs.")
	} else {
		fmt.Printf("Dans votre sac éthéré : %v.\n", c.Inventory)
	}

	// Mort et danger
	if c.IsDead() {
		fmt.Println("\n⚠️  Votre essence vacille… la mort dans le Labyrinthe est bien réelle !")
	} else {
		fmt.Println("\nVotre esprit flotte dans l’obscurité, prêt pour le prochain niveau.")
	}

	fmt.Println("══════════════════════════════════════════════════════════")
}
func (c *Character) IsDead() bool {
	return c.PvCurr <= 0
}

// Resurrect ressuscite le personnage à 50% des HP et Mana max.
func (c *Character) Resurrect() {
	c.PvCurr = c.PvMax / 2
	c.ManaCurr = c.ManaMax / 2
	fmt.Println("Le personnage a été ressuscité !")
}

func (c *Character) PoisonEffect() {
	for i := 0; i < 3; i++ {
		c.PvCurr -= 10
		if c.PvCurr < 0 {
			c.PvCurr = 0
		}
		fmt.Printf("Poison : -10 PV (%d/%d)\n", c.PvCurr, c.PvMax)
		time.Sleep(1 * time.Second)
		if c.IsDead() {
			fmt.Println("⚠️ Votre esprit succombe au poison !")
			break
		}
	}
}

func (c *Character) DisplayInventory() {
	if len(c.Inventory) == 0 {
		fmt.Println("Votre sac est vide.")
		return
	}
	fmt.Println("Inventaire :")
	for i, item := range c.Inventory {
		fmt.Printf("%d) %s\n", i+1, item)
	}
}

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
				fmt.Println("Vous utilisez", itemName)
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	fmt.Println("Objet introuvable :", itemName)
	return false
}

func (c *Character) CountItem(itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

func (c *Character) AddToInventory(item string) bool {
	if len(c.Inventory) >= c.InventorySize {
		fmt.Println("Inventaire plein!")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func (c *Character) TakeItem(itemName string) {
	switch strings.ToLower(itemName) {
	case "potion de vie":
		if c.PvCurr >= c.PvMax {
			fmt.Println("💊 Vos PV sont déjà au maximum !")
			return
		}
		if c.CountItem("Potion de vie") > 0 {
			c.UseItem("Potion de vie")
			c.PvCurr += 30
			if c.PvCurr > c.PvMax {
				c.PvCurr = c.PvMax
			}
			fmt.Printf("💖 Vous buvez une potion de vie. PV: %d/%d\n", c.PvCurr, c.PvMax)
		} else {
			fmt.Println("❌ Aucune potion de vie disponible !")
		}
	case "potion de mana":
		if c.ManaCurr >= c.ManaMax {
			fmt.Println("🔮 Votre énergie magique est déjà au maximum !")
			return
		}
		if c.CountItem("Potion de mana") > 0 {
			c.UseItem("Potion de mana")
			c.ManaCurr += 20
			if c.ManaCurr > c.ManaMax {
				c.ManaCurr = c.ManaMax
			}
			fmt.Printf("✨ Vous buvez une potion de mana. Mana: %d/%d\n", c.ManaCurr, c.ManaMax)
		} else {
			fmt.Println("❌ Aucune potion de mana disponible !")
		}
	default:
		fmt.Println("❌ Cet objet ne peut pas être utilisé directement !")
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

	// Calculer l'excès d'XP pour le niveau suivant
	excessXP := c.XPCurr - c.XPMax

	// Montée de niveau
	c.Level++
	c.XPUpgrades++

	// Bonus stats selon les consignes du projet
	oldMaxHP := c.PvMax
	oldMaxMana := c.ManaMax

	c.PvMax += 20   // +20 MaxHP par niveau
	c.ManaMax += 10 // +10 MaxMana par niveau

	// Soigner complètement au level up (bonus)
	c.PvCurr = c.PvMax
	c.ManaCurr = c.ManaMax

	// Calculer nouvelle XP requise pour le niveau suivant
	c.XPMax = c.CalculateXPNeeded(c.Level)
	c.XPCurr = excessXP // Reporter l'excès

	// Message de level up thématique Somnium
	fmt.Println("\n🌟 ═══ ÉVEIL DE L'ESPRIT ═══ 🌟")
	fmt.Printf("Votre conscience s'élève... Niveau %d atteint !\n", c.Level)
	fmt.Printf("💖 Vitalité : %d → %d (+20)\n", oldMaxHP, c.PvMax)
	fmt.Printf("🔮 Essence : %d → %d (+10)\n", oldMaxMana, c.ManaMax)
	fmt.Println("Votre esprit et corps sont restaurés par cette révélation !")

	if excessXP > 0 {
		fmt.Printf("📊 XP en excès reportée : %d/%d\n", c.XPCurr, c.XPMax)
	}
	fmt.Println("═══════════════════════════════════════")
}

// CalculateXPNeeded calcule l'XP nécessaire pour un niveau donné
func (c *Character) CalculateXPNeeded(level int) int {
	// Formule : Level * 100 (selon consignes)
	// Niveau 1 → 100 XP
	// Niveau 2 → 200 XP
	// Niveau 3 → 300 XP, etc.
	return level * 100
}

// DisplayXPInfo affiche les informations d'XP (utilitaire)
func (c *Character) DisplayXPInfo() {
	nextLevelXP := c.CalculateXPNeeded(c.Level + 1)
	fmt.Printf("📊 Expérience : %d/%d (Niveau %d)\n", c.XPCurr, nextLevelXP, c.Level)
	if c.Level < 10 { // Limite arbitraire
		remaining := nextLevelXP - c.XPCurr
		fmt.Printf("   Prochain niveau dans : %d XP\n", remaining)
	}
}

func (c *Character) GetName() string {
	return c.Name
}

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
	fmt.Printf("🎲 %s lance l'initiative : %d (base) + %d (classe) = %d\n", 
		c.Name, baseRoll, bonus, c.Initiative)
}

func (c *Character) RestoreHealth() {
	c.PvCurr = c.PvMax
}