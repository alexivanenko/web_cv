package model

import (
	"crypto/md5"
	"fmt"
	"io"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ObjectID bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Login    string        `json:"login" bson:"login"`
	Password string        `json:"password" bson:"password"`
}

func (user *User) LoadByLogin(login string) error {
	db := GetDB()
	b := db.C("users")

	return b.Find(bson.M{"login": login}).One(&user)
}

func encryptPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	salt1 := "@#$%"
	salt2 := "^&*()"

	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	return fmt.Sprintf("%x", h.Sum(nil))
}
