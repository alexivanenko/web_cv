package model

import (
	"encoding/json"

	"strconv"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

const Resume_Table = "resume"

type Education struct {
	From        int    `json:"from" bson:"from"`
	To          string `json:"to" bson:"to"`
	Grade       string `json:"grade" bson:"grade"`
	Science     string `json:"science" bson:"science"`
	Institution string `json:"institution" bson:"institution"`
	Description string `json:"description" bson:"description"`
}

type Employment struct {
	From        int    `json:"from" bson:"from"`
	To          string `json:"to" bson:"to"`
	Company     string `json:"company" bson:"company"`
	Position    string `json:"position" bson:"position"`
	Description string `json:"description" bson:"description"`
}

type Skill struct {
	Name       string `json:"name" bson:"name"`
	LevelInPct int    `json:"level_in_pct" bson:"level_in_pct"`
	Color      string `json:"color" bson:"color"`
}

type Fact struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type Resume struct {
	base        *Base
	ObjectID    bson.ObjectId `json:"_id" bson:"_id"`
	Login       string        `json:"login" bson:"login"`
	Educations  []Education   `json:"educations" bson:"educations"`
	Employments []Employment  `json:"employments" bson:"employments"`
	Skills      []Skill       `json:"skills" bson:"skills"`
	Facts       []Fact        `json:"facts" bson:"facts"`
}

func (resume *Resume) getCollectionName() string {
	return Resume_Table
}

func (resume *Resume) LoadByLogin(login string) error {
	data, err := resume.base.LoadByLogin(&resume, resume.getCollectionName(), login)

	if err != nil {
		return err
	}

	encoded, _ := json.Marshal(data)
	json.Unmarshal(encoded, &resume)

	fmt.Println(resume)
	return err
}

func (resume *Resume) Bind(c *gin.Context, params map[string]string) {
	c.BindWith(resume, binding.Form)
	resume.Login = params["login"]

	if bson.IsObjectIdHex(c.PostForm("id")) {
		resume.ObjectID = bson.ObjectIdHex(c.PostForm("id"))
	}

	educationInstitutions := c.Request.Form["Education.Institution"]

	for index, value := range educationInstitutions {
		education := Education{}
		education.Institution = value

		education.From, _ = strconv.Atoi(c.Request.Form["Education.From"][index])
		education.To = c.Request.Form["Education.To"][index]
		education.Grade = c.Request.Form["Education.Grade"][index]
		education.Science = c.Request.Form["Education.Science"][index]
		education.Description = c.Request.Form["Education.Description"][index]

		resume.Educations = append(resume.Educations, education)
	}

	employmentCompanies := c.Request.Form["Employment.Company"]

	for index, value := range employmentCompanies {
		employment := Employment{}
		employment.Company = value

		employment.From, _ = strconv.Atoi(c.Request.Form["Employment.From"][index])
		employment.To = c.Request.Form["Employment.To"][index]
		employment.Position = c.Request.Form["Employment.Position"][index]
		employment.Description = c.Request.Form["Employment.Description"][index]

		resume.Employments = append(resume.Employments, employment)
	}

	skillNames := c.Request.Form["Skill.Name"]

	for index, value := range skillNames {
		skill := Skill{}
		skill.Name = value

		skill.LevelInPct, _ = strconv.Atoi(c.Request.Form["Skill.LevelInPct"][index])
		skill.Color = c.Request.Form["Skill.Color"][index]

		resume.Skills = append(resume.Skills, skill)
	}

	factNames := c.Request.Form["Fact.Name"]

	for index, value := range factNames {
		fact := Fact{}
		fact.Name = value

		fact.Value = c.Request.Form["Fact.Value"][index]

		resume.Facts = append(resume.Facts, fact)
	}
}

func (resume *Resume) Validate() map[string]int {
	var err map[string]int = make(map[string]int)
	return err
}

func (resume *Resume) Save() error {
	return resume.base.Save(resume.ObjectID, resume, resume.getCollectionName())
}
