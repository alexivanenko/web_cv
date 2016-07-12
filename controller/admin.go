package controller

import (
	"bytes"
	"strings"

	"github.com/alexivanenko/web_cv/config"
	"github.com/alexivanenko/web_cv/model"
	"github.com/gin-gonic/gin"
)

func AdminController(c *gin.Context) {
	var section string = strings.Replace(c.Request.URL.Path, "/admin/", "", 1)
	params := gin.H{}
	errors := make(map[string]int)
	var dataModel interface{}

	if section == "" {
		section = "index"
	}

	errors, params, dataModel = AdminAction(c, section)

	if c.Request.Method == "POST" {
		redirect(c, c.Request.URL.Path)
		return
	} else {
		params["model"] = dataModel
		params["errors"] = errors
	}

	var tplName bytes.Buffer
	tplName.WriteString("admin/")
	tplName.WriteString(section)
	tplName.WriteString(".html")

	render(c, tplName.String(), params)
}

func AdminAction(c *gin.Context, section string) (map[string]int, gin.H, interface{}) {
	errors := make(map[string]int)
	var dataModel interface{}
	params := gin.H{}

	switch section {
	case "profile":
		dataModel = new(model.Profile)
	case "resume":
		dataModel = new(model.Resume)
	case "portfolio":
		dataModel = new(model.Portfolio)
	default:
		dataModel = new(model.Developer)
	}

	if c.Request.Method == "POST" {

		details := make(map[string]string)
		details["login"] = config.String("view.nickname")
		details["rootDir"] = config.GetRootDir()

		dataModel.(model.BaseModel).Bind(c, details)

		if dataErrors := dataModel.(model.BaseModel).Validate(); len(dataErrors) > 0 {
			errors = dataErrors
		} else {
			dataModel.(model.BaseModel).Save()

			return nil, nil, nil
		}

	} else {
		dataModel.(model.BaseModel).LoadByLogin(config.String("view.nickname"))

		if section == "index" {
			params["statusList"] = dataModel.(*model.Developer).GetStatusList()
		} else if section == "portfolio" {
			params["categoriesList"] = dataModel.(*model.Portfolio).GetCategoriesList()
		}
	}

	return errors, params, dataModel
}

func RemoveProject(c *gin.Context) {
	id := c.Query("id")

	portfolio := new(model.Portfolio)
	err := portfolio.DeleteProject(id)

	if err != nil {
		config.Log("There are some errors in deleting project process (ID:" + id + ")")
	}

	c.JSON(200, gin.H{"id": id})
}

func LoadProject(c *gin.Context) {
	portfolioId := c.Query("id")
	projectId := c.Query("project_id")

	portfolio := new(model.Portfolio)
	project := portfolio.LoadProject(portfolioId, projectId)

	c.JSON(200, project)
}
