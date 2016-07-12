package model

import (
	"encoding/json"

	"io"
	"os"

	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

const Portfolio_Table = "portfolio"
const Image_Dir = "/static/img/folio/"

type Category struct {
	Name string
}

type Project struct {
	ObjectID         bson.ObjectId `json:"_id" bson:"_id"`
	Name             string        `json:"name" bson:"name"`
	Category         string        `json:"category" bson:"category"`
	Url              string        `json:"url" bson:"url"`
	ShortDescription string        `json:"short_description" bson:"short_description"`
	Description      string        `json:"description" bson:"description"`
	Image            string        `json:"image" bson:"image"`
	Order            string        `json:"order" bson:"order"`
}

type Portfolio struct {
	base          *Base
	ObjectID      bson.ObjectId `json:"_id" bson:"_id"`
	Login         string        `json:"login" bson:"login"`
	Projects      []Project     `json:"projects" bson:"projects"`
	UpdateProject bool
}

// ByOrder implements sort.Interface for []Project based on Order field
type ByOrder []Project

func (a ByOrder) Len() int      { return len(a) }
func (a ByOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool {
	e1, _ := strconv.Atoi(a[i].Order)
	e2, _ := strconv.Atoi(a[j].Order)
	return e1 > e2
}

func (portfolio *Portfolio) getCollectionName() string {
	return Portfolio_Table
}

func (portfolio *Portfolio) LoadByLogin(login string) error {
	data, err := portfolio.base.LoadByLogin(&portfolio, portfolio.getCollectionName(), login)

	if err != nil {
		return err
	}

	encoded, _ := json.Marshal(data)
	json.Unmarshal(encoded, &portfolio)

	sort.Stable(ByOrder(portfolio.Projects))

	return err
}

func (portfolio *Portfolio) LoadProject(portfolioId string, projectId string) Project {
	db := GetDB()
	b := db.C(portfolio.getCollectionName())

	var idHex interface{}
	idHex = ""

	if bson.IsObjectIdHex(portfolioId) {
		idHex = bson.ObjectIdHex(portfolioId)
	}

	err := b.Find(bson.M{"_id": idHex}).One(&portfolio)
	var result Project

	if err == nil {
		for _, project := range portfolio.Projects {
			if project.ObjectID.Hex() == projectId {
				result = project
				break
			}
		}
	}

	return result
}

func (portfolio *Portfolio) Bind(c *gin.Context, params map[string]string) {
	portfolio.LoadByLogin(params["login"])

	portfolio.Login = params["login"]

	if bson.IsObjectIdHex(c.PostForm("id")) {
		portfolio.ObjectID = bson.ObjectIdHex(c.PostForm("id"))
	}

	project := Project{}
	c.BindWith(&project, binding.Form)

	if bson.IsObjectIdHex(c.PostForm("projectId")) {
		project.ObjectID = bson.ObjectIdHex(c.PostForm("projectId"))
		portfolio.UpdateProject = true
	} else {
		project.ObjectID = bson.NewObjectId()
	}

	_, _, uploadErr := c.Request.FormFile("Image")

	if c.PostForm("projectImage") != "" && uploadErr != nil {
		project.Image = c.PostForm("projectImage")
	} else {
		err, filename := portfolio.uploadImage(c, params["rootDir"])

		if err == nil {
			project.Image = filename
		}
	}

	//Clear Projects and add only one (updated)
	if portfolio.UpdateProject {
		portfolio.Projects = []Project{}
	}

	portfolio.Projects = append(portfolio.Projects, project)
}

func (portfolio *Portfolio) uploadImage(c *gin.Context, rootDir string) (error, string) {
	file, header, uploadErr := c.Request.FormFile("Image")
	filename := ""
	var err error

	if uploadErr == nil {
		filename = header.Filename
		out, err := os.Create(rootDir + Image_Dir + filename)

		if err == nil {
			defer out.Close()
			_, err = io.Copy(out, file)
		}
	}

	return err, filename
}

func (portfolio *Portfolio) Validate() map[string]int {
	var err map[string]int = make(map[string]int)
	return err
}

func (portfolio *Portfolio) Save() error {

	if portfolio.UpdateProject {
		return portfolio.updateProject(portfolio.Projects[0])
	} else {
		return portfolio.base.Save(portfolio.ObjectID, portfolio, portfolio.getCollectionName())
	}
}

func (portfolio *Portfolio) updateProject(project Project) error {
	db := GetDB()
	b := db.C(portfolio.getCollectionName())

	colQuery := bson.M{"projects._id": project.ObjectID}

	change := bson.M{"$set": bson.M{
		"projects.$.name":              project.Name,
		"projects.$.category":          project.Category,
		"projects.$.url":               project.Url,
		"projects.$.short_description": project.ShortDescription,
		"projects.$.description":       project.Description,
		"projects.$.image":             project.Image,
		"projects.$.order":             project.Order,
	}}

	err := b.Update(colQuery, change)

	return err
}

func (portfolio *Portfolio) DeleteProject(id string) error {
	db := GetDB()
	b := db.C(portfolio.getCollectionName())

	var idHex interface{}
	idHex = ""

	if bson.IsObjectIdHex(id) {
		idHex = bson.ObjectIdHex(id)
	}

	colQuery := bson.M{"projects._id": idHex}
	change := bson.M{"$pull": bson.M{"projects": bson.M{"_id": idHex}}}
	err := b.Update(colQuery, change)

	return err
}

func (portfolio *Portfolio) GetCategoriesList() []Category {
	return []Category{
		Category{
			Name: "Coding",
		},
		Category{
			Name: "Maintenance",
		},
	}
}
