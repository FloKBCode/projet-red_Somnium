package character

type Character  struct {
	Name string
	Race string
	Class string
	Level int
	PvMax int
	PvCurr int
	Inventory []string
	Money int
	Skills []string
	ManaMax int
	ManaCurr int
	Equipment Equipment
}

type Equipment struct {
    Head  string
    Chest string
    Feet  string
}

func InitCharacter(name, race, class string)Character{

}