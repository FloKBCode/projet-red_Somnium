package game

type Quest struct {
	ID          int
	Title       string
	Description string
	IsCompleted bool
	RewardGold  int
	RewardExp   int
}

var availableQuests = []Quest{
	{
		ID:          1,
		Title:       "Le Cauchemar Initial",
		Description: "Vaincre le premier boss du labyrinthe",
		RewardGold:  100,
		RewardExp:   50,
	},
	// Ajoutez d'autres quêtes
}
