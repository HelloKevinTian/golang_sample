package models

import (
	. "app/consts"
	"gslib"
)

type UserModel struct {
	gslib.BaseModel
	Data *User
}
