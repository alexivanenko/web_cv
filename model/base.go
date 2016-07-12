package model

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type BaseModel interface {
	Bind(c *gin.Context, params map[string]string)
	LoadByLogin(login string) error
	Validate() map[string]int
	Save() error
}

type Base struct {
	ObjectID bson.ObjectId
}

func (base *Base) LoadByLogin(object interface{}, collection string, login string) (interface{}, error) {
	db := GetDB()
	b := db.C(collection)

	err := b.Find(bson.M{"login": login}).One(&object)

	return object, err
}

func (base *Base) Save(id interface{}, model interface{}, collection string) error {
	db := GetDB()
	b := db.C(collection)

	idStr := reflect.ValueOf(model).Elem().FieldByName("ObjectID").String()

	var storedId interface{}
	storedId = id

	if len(idStr) < 1 {
		storedId = bson.NewObjectId()
		reflect.ValueOf(model).Elem().FieldByName("ObjectID").Set(reflect.ValueOf(storedId))
	}

	_, err := b.UpsertId(storedId, model)

	return err
}
