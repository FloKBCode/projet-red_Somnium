package character

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type Character struct {
	Name      string
	Race      string
	Class     string
	Level     int
	PvMax     int
	PvCurr    int
	Inventory []string
	Money     int
	Skills    []string
	ManaMax   int
	ManaCurr  int
	Equipment Equipment
}

type Equipment struct {
	Head  string
	Chest string
	Feet  string
}

func InitCharacter(name, race, class string, pvMax, manaMax int) Character {
	return Character{
		Name:      name,
		Race:      race, // ✅ Ajouté (était manquant)
		Class:     class,
		Level:     1,
		PvMax:     pvMax, // ✅ Utiliser les paramètres
		PvCurr:    int(math.Round(float64(pvMax) * 0.5)),
		Inventory: []string{"Potion de vie", "Potion de vie", "Potion de vie"}, // ✅ Consignes : 3 potions
		Money:     100,                                                         // ✅ Consignes : 100 pièces d'or
		Skills:    []string{"Coup de poing"},
		ManaCurr:  manaMax, // ✅ Utiliser les paramètres
		ManaMax:   manaMax,
		Equipment: Equipment{},
	}
}

func (c *Character) DisplayInfo() {
	fmt.Println("═══════════════ LABYRINTHE DES CAUCHEMARS ═══════════════")
	fmt.Printf("Esprit : %s, %s %s, errant entre les couches de conscience.\n",
		c.Name, c.Race, c.Class)
	fmt.Printf("Niveau de conscience : %d — chaque pas pourrait être le dernier.\n", c.Level)

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
			case "potion de rêve":
				c.PvCurr += 20
				if c.PvCurr > c.PvMax {
					c.PvCurr = c.PvMax
				}
			case "amulette du souvenir":
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
