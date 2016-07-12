package model

import (
	"strings"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

const Profile_Table = "profile"

type Service struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Icon        string `json:"icon" bson:"icon"`
}

type Testimonial struct {
	Content       string `json:"content" bson:"content"`
	WriterName    string `json:"writer_name" bson:"writer_name"`
	WriterCompany string `json:"writer_company" bson:"writer_company"`
}

type Profile struct {
	base         *Base
	ObjectID     bson.ObjectId `json:"_id" bson:"_id"`
	Login        string        `json:"login" bson:"login"`
	Description  string        `json:"description" bson:"description"`
	Services     []Service     `json:"services" bson:"services"`
	Testimonials []Testimonial `json:"testimonials" bson:"testimonials"`
}

func (profile *Profile) getCollectionName() string {
	return Profile_Table
}

func (profile *Profile) LoadByLogin(login string) error {
	data, err := profile.base.LoadByLogin(&profile, profile.getCollectionName(), login)

	if err != nil {
		return err
	}

	encoded, _ := json.Marshal(data)
	json.Unmarshal(encoded, &profile)

	return err
}

func (profile *Profile) Bind(c *gin.Context, params map[string]string) {
	c.BindWith(profile, binding.Form)
	profile.Login = params["login"]

	if bson.IsObjectIdHex(c.PostForm("id")) {
		profile.ObjectID = bson.ObjectIdHex(c.PostForm("id"))
	}

	serviceNames := c.Request.Form["Service.Name"]

	for index, value := range serviceNames {
		service := Service{}
		service.Name = value
		service.Description = c.Request.Form["Service.Description"][index]
		service.Icon = c.Request.Form["Service.Icon"][index]

		profile.Services = append(profile.Services, service)
	}

	testimonialContent := c.Request.Form["Testimonial.Content"]

	for index, value := range testimonialContent {
		testimonial := Testimonial{}
		testimonial.Content = value
		testimonial.WriterName = c.Request.Form["Testimonial.WriterName"][index]
		testimonial.WriterCompany = c.Request.Form["Testimonial.WriterCompany"][index]

		profile.Testimonials = append(profile.Testimonials, testimonial)
	}
}

func (profile *Profile) Validate() map[string]int {
	var err map[string]int = make(map[string]int)

	if strings.TrimSpace(profile.Description) == "" {
		err["Description"] = 1
	}

	return err
}

func (profile *Profile) Save() error {
	return profile.base.Save(profile.ObjectID, profile, profile.getCollectionName())
}
