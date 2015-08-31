package game_data

// 1
type ConfigEquip struct {
	EquipId   int
	EquipType int
	ReqLv     int
}

var ConfigEquips = map[int]ConfigEquip{
	1: ConfigEquip{EquipId: 1, EquipType: 1, ReqLv: 10},
	2: ConfigEquip{EquipId: 1, EquipType: 1, ReqLv: 10},
}

func FindEquip(key int) ConfigEquip {
	return ConfigEquips[key]
}
