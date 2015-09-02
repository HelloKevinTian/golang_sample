package game_data

// 2
type ConfigForge struct {
	EquipId int
	Health  int
	Attack  int
}

var ConfigForges = map[int]ConfigForge{
	1: ConfigForge{1, 2, 3},
	2: ConfigForge{2, 3, 4},
}

func FindForge(key int) ConfigForge {
	return ConfigForges[key]
}
