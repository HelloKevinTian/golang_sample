package gslib

import (
	. "app/consts"
	// "fmt"
	"reflect"
)

type BaseModel struct {
	TableName string
	Uuid      string
	Ctx       *Player
}

func (self *BaseModel) Save() {
	self.Ctx.Store.UpdateStatus([]string{"models", self.TableName}, self.Uuid, STATUS_UPDATE)
}

func (self *BaseModel) Delete() {
	self.Ctx.Store.Del([]string{"models", self.TableName}, self.Uuid)
}

//gslib.CreateModel(self.Ctx, &EquipModel{Data: &Equip{Uuid: uuid}})
func CreateModel(ctx *Player, model interface{}) {
	rv := reflect.ValueOf(model).Elem()
	data := rv.FieldByName("Data").Elem()
	uuid := data.FieldByName("Uuid").String()
	tableName := StructToTableNameMap[data.Type().Name()]
	base := BaseModel{tableName, uuid, ctx}
	rv.Field(0).Set(reflect.ValueOf(base))
	ctx.Store.Set([]string{"models", tableName}, uuid, model)
}

func FindModel(ctx *Player, tableName, uuid string) interface{} {
	return ctx.Store.Get([]string{"models", tableName}, uuid)
}
