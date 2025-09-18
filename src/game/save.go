package game

import (
	"encoding/json"
	"os"
	"somnium/character"
)

// Sauvgarde la partie
func SaveGame(player *character.Character) error {
	data, err := json.Marshal(player)
	if err != nil {
		return err
	}
	return os.WriteFile("save.json", data, 0644)
}

// Charger la partie
func LoadGame() (*character.Character, error) {
	data, err := os.ReadFile("save.json")
	if err != nil {
		return nil, err
	}

	var player character.Character
	err = json.Unmarshal(data, &player)
	return &player, err
}
