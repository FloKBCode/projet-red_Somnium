package character

import (
	"fmt"
	"math"
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

func InitCharacter(name, race, class string) Character {
	PvMax, ManaMax := getBaseStats(race, class)
	return Character{
		Name:      name,
		Class:     class,
		Level:     1,
		PvMax:     PvMax,
		PvCurr:    int(math.Round(float64(PvMax) * 0.5)),
		Inventory: []string{},
		Money:     20,
		Skills:    []string{"Coup de poing"},
		ManaCurr:  ManaMax,
		ManaMax:   ManaMax,
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
