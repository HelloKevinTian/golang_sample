package register

import (
// . "app/consts"
// . "app/models"
// "gslib"
// "gslib/store"
)

// For overwirte auto generate data_loader
func CustomRegisterDataLoader() {
	// store.RegisterDataLoader("equips", func(playerId string, ctx interface{}) interface{} {
	// 	var datas []Equip
	// 	var models []*EquipModel
	// 	ctx.(*gslib.Player).Store.Db.Select(&datas, "SELECT * FROM equips where user_id=?", playerId)
	// 	for i := 0; i < len(datas); i++ {
	// 		data := datas[i]
	// 		models = append(models, &EquipModel{gslib.BaseModel{"equips", data.Uuid, ctx.(*gslib.Player)}, &data})
	// 	}
	// 	return models
	// })
}
