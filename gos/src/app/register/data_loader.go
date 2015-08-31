//Generated by generate_tables

package register

import (
	. "app/consts"
	. "app/models"
	"gslib"
	"gslib/store"
)

func RegisterDataLoader() {
	store.RegisterDataLoader("equips", func(playerId string, ctx interface{}) interface{} {
		var datas []Equip
		var models []*EquipModel
		ctx.(*gslib.Player).Store.Db.Select(&datas, "SELECT * FROM equips where user_id=?", playerId)
		for i := 0; i < len(datas); i++ {
			data := datas[i]
			models = append(models, &EquipModel{gslib.BaseModel{"equips", data.Uuid, ctx.(*gslib.Player)}, &data})
		}
		return models
	})

	store.RegisterDataLoader("users", func(playerId string, ctx interface{}) interface{} {
		var datas []User
		var models []*UserModel
		ctx.(*gslib.Player).Store.Db.Select(&datas, "SELECT * FROM users where user_id=?", playerId)
		for i := 0; i < len(datas); i++ {
			data := datas[i]
			models = append(models, &UserModel{gslib.BaseModel{"users", data.Uuid, ctx.(*gslib.Player)}, &data})
		}
		return models
	})

}