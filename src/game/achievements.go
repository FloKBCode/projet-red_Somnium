// Pas utilisé
package game

type Achievement struct {
	ID          string
	Name        string
	Description string
	Unlocked    bool
}

var gameAchievements = []Achievement{
	{
		ID:          "first_victory",
		Name:        "Premier Pas",
		Description: "Gagner votre premier combat",
	},
	{
		ID:          "collector",
		Name:        "Collectionneur",
		Description: "Obtenir 10 objets différents",
	},
}
