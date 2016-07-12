package model

import (
	"time"

	"strings"

	"fmt"

	"encoding/json"

	"github.com/alexivanenko/web_cv/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

const Dev_Table = "developer"

type Status struct {
	Name string
}

type Developer struct {
	base             *Base
	ObjectID         bson.ObjectId `json:"_id" bson:"_id"`
	Login            string        `json:"login" bson:"login"`
	Status           string        `json:"status" bson:"status"`
	Technologies     string        `json:"technologies" bson:"technologies"`
	Description      string        `json:"description" bson:"description"`
	Email            string        `json:"email" bson:"email"`
	Phone            string        `json:"phone" bson:"phone"`
	BirthDate        time.Time     `json:"birth_date" bson:"birth_date"`
	Languages        string        `json:"languages" bson:"languages"`
	TechnologiesList []string
}

func (developer *Developer) getCollectionName() string {
	return Dev_Table
}

func (developer *Developer) LoadByLogin(login string) error {
	data, err := developer.base.LoadByLogin(&developer, developer.getCollectionName(), login)

	if err != nil {
		return err
	}

	encoded, _ := json.Marshal(data)
	json.Unmarshal(encoded, &developer)

	developer.TechnologiesList = strings.Split(developer.Technologies, ",")

	return err
}

func (developer *Developer) GetStatus(login string) string {
	db := GetDB()
	b := db.C(developer.getCollectionName())

	var result []struct {
		Status string `bson:"status"`
	}
	err := b.Find(bson.M{"login": "alexivanenko"}).Select(bson.M{"status": 1}).All(&result)

	if err != nil {
		config.Log(fmt.Sprintf("Load developer error"))
	}

	return result[0].Status
}

func (developer *Developer) Bind(c *gin.Context, params map[string]string) {
	c.BindWith(developer, binding.Form)
	developer.Login = params["login"]
	developer.BirthDate = time.Date(1983, time.December, 15, 10, 0, 0, 0, time.Local)

	if bson.IsObjectIdHex(c.PostForm("id")) {
		developer.ObjectID = bson.ObjectIdHex(c.PostForm("id"))
	}
}

func (developer *Developer) Validate() map[string]int {
	var err map[string]int = make(map[string]int)

	if strings.TrimSpace(developer.Technologies) == "" {
		err["Technologies"] = 1
	}

	if strings.TrimSpace(developer.Description) == "" {
		err["Description"] = 1
	}

	if strings.TrimSpace(developer.Email) == "" {
		err["Email"] = 1
	}

	if strings.TrimSpace(developer.Phone) == "" {
		err["Phone"] = 1
	}

	if strings.TrimSpace(developer.Languages) == "" {
		err["Languages"] = 1
	}

	return err
}

func (developer *Developer) Save() error {
	return developer.base.Save(developer.ObjectID, developer, developer.getCollectionName())
}

func (developer *Developer) GetStatusList() []Status {
	return []Status{
		Status{
			Name: "Full-Time Available",
		},
		Status{
			Name: "Part-Time Available",
		},
		Status{
			Name: "Not Available",
		},
		Status{
			Name: "Traveling",
		},
	}
}
