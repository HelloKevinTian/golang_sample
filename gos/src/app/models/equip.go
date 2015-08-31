package models

import (
	. "app/consts"
	"fmt"
	"gslib"
)

type EquipModel struct {
	gslib.BaseModel
	Data *Equip
}

func (e *EquipModel) Load(heroId string) {
	fmt.Println("Load equip to: ", heroId)
}
