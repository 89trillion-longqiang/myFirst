package soldierInfo

type Info struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	UnlockArena string `json:"UnlockArena"`
	Rarity       string `json:"rarity"`
	CombatPoints string `json:"combatPoints"`
	Desc         string `json:"desc"`
}

