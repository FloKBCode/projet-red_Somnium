package combat

type Difficulty int

const (
	Easy Difficulty = iota
	Normal
	Hard
)

func AdjustMonsterStats(monster *Monster, difficulty Difficulty) {
	switch difficulty {
	case Easy:
		monster.PvMax = int(float64(monster.PvMax) * 0.8)
		monster.Attack = int(float64(monster.Attack) * 0.8)
	case Hard:
		monster.PvMax = int(float64(monster.PvMax) * 1.2)
		monster.Attack = int(float64(monster.Attack) * 1.2)
	}
	monster.PvCurr = monster.PvMax
}
