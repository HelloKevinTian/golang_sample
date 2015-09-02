package store

import (
	. "app/consts"
	"app/register/tables"
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type Filter func(elem interface{}) bool

type Store map[string]interface{}

type dataLoader func(modelName string, ctx interface{}) interface{}

var dataLoaderMap = map[string]dataLoader{}

type Ets struct {
	store Store
	Db    *gorp.DbMap
	Ctx   interface{}
}

const (
	GET     = 1
	LOAD    = 2
	SET     = 3
	DEL     = 4
	FIND    = 5
	SELECT  = 6
	PERSIST = 7
)

func Test(ets *Ets) {
	equip := &Equip{}
	ets.Db.SelectOne(&equip, "select * from equips where uuid='54A3927E2B89780A1491F441'")
	fmt.Println("Store Test:", equip)

	key := "54A3927E2B89780A1491F441"
	namespaces := []string{"54A3927E2B89780A1491F43C", "equips"}
	ets.Load(namespaces, key, equip)

	fmt.Println("Get: ", ets.Get(namespaces, key))

	equip.Level = 0
	ets.Set(namespaces, key, equip)
	ets.Del(namespaces, "1")
	ets.Persist([]string{"54A3927E2B89780A1491F43C"})
}

var sharedDBInstance *gorp.DbMap

func InitDB() {
	db, err := sql.Open("mysql", "root:@/game_server_development")
	if err != nil {
		panic(err.Error())
	}
	sharedDBInstance = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	tables.RegisterTables(sharedDBInstance)
}

func GetSharedDBInstance() *gorp.DbMap {
	return sharedDBInstance
}

func New(ctx interface{}) *Ets {
	e := &Ets{
		store: make(Store),
		Db:    GetSharedDBInstance(),
		Ctx:   ctx,
	}
	return e
}

func RegisterDataLoader(modelName string, loader dataLoader) {
	dataLoaderMap[modelName] = loader
}

func (e *Ets) LoadData(modelName, playerId string) {
	handler, ok := dataLoaderMap[modelName]
	if ok {
		values := reflect.ValueOf(handler(playerId, e.Ctx))
		for i := 0; i < values.Len(); i++ {
			model := values.Index(i)
			value := model.Elem()
			data := value.FieldByName("Data")
			key := data.Elem().FieldByName("Uuid").String()
			e.Load([]string{"models", modelName}, key, model.Interface())
		}
	}
}

func (e *Ets) Get(namespaces []string, key string) interface{} {
	if ctx := e.getCtx(namespaces); ctx != nil {
		return ctx[key]
	} else {
		return nil
	}
}

func (e *Ets) Load(namespaces []string, key string, value interface{}) {
	ctx := e.makeCtx(namespaces)
	ctx[key] = value
}

func (e *Ets) Set(namespaces []string, key string, value interface{}) {
	ctx := e.makeCtx(namespaces)
	if ctx[key] == nil {
		e.UpdateStatus(namespaces, key, STATUS_CREATE)
	} else {
		e.UpdateStatus(namespaces, key, STATUS_UPDATE)
	}
	ctx[key] = value
}

func (e *Ets) Del(namespaces []string, key string) {
	if ctx := e.getCtx(namespaces); ctx != nil {
		e.UpdateStatus(namespaces, key, STATUS_DELETE)
		delete(ctx, key)
	}
}

func (e *Ets) Find(namespaces []string, filter Filter) interface{} {
	if ctx := e.getCtx(namespaces); ctx != nil {
		for _, v := range ctx {
			if filter(v) {
				return v
			}
		}
	}
	return nil
}

func (e *Ets) Select(namespaces []string, filter Filter) interface{} {
	var elems []interface{}
	if ctx := e.getCtx(namespaces); ctx != nil {
		for _, v := range ctx {
			if filter(v) {
				elems = append(elems, v)
			}
		}
		return elems
	}
	return elems
}

func (e *Ets) Count(namespaces []string) int {
	if ctx := e.getCtx(namespaces); ctx != nil {
		return len(ctx)
	} else {
		return 0
	}
}

func (e *Ets) Persist(namespaces []string) {
	trans, err := e.Db.Begin()
	if err != nil {
		panic(err.Error())
	}
	for tableName, tableCtx := range e.getCtx(namespaces) {
		status := e.allStatus([]string{namespaces[0], tableName})
		executeSql(trans, tableName, status, tableCtx.(Store))
	}
	err = trans.Commit()
	if err != nil {
		panic(err.Error())
	}
	e.cleanStatus(namespaces)
}

func (e *Ets) getCtx(namespaces []string) Store {
	var ctx Store = nil
	for _, namespace := range namespaces {
		if ctx == nil {
			vctx, ok := e.store[namespace]
			if !ok {
				return nil
			}
			ctx = vctx.(Store)
		} else {
			vctx, ok := ctx[namespace]
			if !ok {
				return nil
			}
			ctx = vctx.(Store)
		}
	}
	return ctx
}

func (e *Ets) makeCtx(namespaces []string) Store {
	var ctx Store = nil
	for _, namespace := range namespaces {
		if ctx == nil {
			vctx, ok := e.store[namespace]
			if !ok {
				ctx = make(Store)
				e.store[namespace] = ctx
			} else {
				ctx = vctx.(Store)
			}
		} else {
			vctx, ok := ctx[namespace]
			if !ok {
				vctx = make(Store)
				ctx[namespace] = vctx
			}
			ctx = vctx.(Store)
		}
	}
	return ctx
}

func (e *Ets) UpdateStatus(namespaces []string, key string, status int) {
	statusKey := getStatusKey(namespaces)
	ctx, ok := e.store[statusKey]
	if !ok {
		ctx = make(Store)
		e.store[statusKey] = ctx
	}
	ctx.(Store)[key] = status
}

func (e *Ets) getStatus(namespaces []string, key string) int {
	statusKey := getStatusKey(namespaces)
	ctx, ok := e.store[statusKey]
	if !ok {
		return STATUS_ORIGIN
	} else {
		return ctx.(Store)[key].(int)
	}
}

func (e *Ets) allStatus(namespaces []string) Store {
	ctx := e.store[getStatusKey(namespaces)]
	if ctx != nil {
		return ctx.(Store)
	} else {
		return nil
	}
}

func (e *Ets) cleanStatus(namespaces []string) {
	delete(e.store, getStatusKey(namespaces))
}

func getStatusKey(namespaces []string) string {
	return strings.Join(append(namespaces, "status"), "_")
}

func executeSql(trans *gorp.Transaction, tableName string, status Store, tableCtx Store) {
	for k, v := range status {
		switch v.(int) {
		case STATUS_UPDATE:
			fmt.Println("STATUS_UPDATE: ", reflect.ValueOf(tableCtx[k]).Elem().FieldByName("Data").Interface())
			_, err := trans.Update(reflect.ValueOf(tableCtx[k]).Elem().FieldByName("Data").Interface())
			if err != nil {
				panic(err.Error())
			}
		case STATUS_DELETE:
			_, err := trans.Exec(fmt.Sprintf("DELETE FROM `%s` WHERE `uuid`='%s'", tableName, k))
			if err != nil {
				panic(err.Error())
			}
		case STATUS_CREATE:
			err := trans.Insert(reflect.ValueOf(tableCtx[k]).Elem().FieldByName("Data").Interface())
			if err != nil {
				panic(err.Error())
			}
		}
	}
}
