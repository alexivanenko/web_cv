package controller

import (
	"html/template"
	"path/filepath"

	"github.com/alexivanenko/web_cv/config"
	"github.com/alexivanenko/web_cv/model"

	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var templates map[string]*template.Template

func render(c *gin.Context, tplName string, params gin.H) {
	params["url_path"] = c.Request.URL.Path
	params["author"] = config.String("view.author")

	developer := new(model.Developer)
	params["status"] = developer.GetStatus(config.String("view.nickname"))

	var isAdmin int = 0
	if strings.Contains(c.Request.URL.Path, "admin") {
		isAdmin = 1
	}
	params["is_admin"] = isAdmin

	c.Writer.WriteHeader(http.StatusOK)

	if err := templates[tplName].Execute(c.Writer, params); err != nil {
		panic(err)
	}
}

func redirect(c *gin.Context, url string) {
	http.Redirect(c.Writer, c.Request, url, http.StatusFound)
}

func getTplFuncMap() template.FuncMap {
	var funcMap = template.FuncMap{}

	funcMap["format_date"] = func(value time.Time) string {
		return value.Format(config.String("view.date_format"))
	}
	funcMap["no_escape"] = func(s string) template.HTML {
		return template.HTML(s)
	}
	funcMap["mod"] = func(i, j int) bool {
		return i%j == 0
	}

	return funcMap
}

func init() {
	partialFiles, _ := filepath.Glob(config.GetRootDir() + "/templates/partial/*.html")
	partialFiles = append(partialFiles, fmt.Sprintf("%s/templates/base.html", config.GetRootDir()))

	baseTemplate := template.Must(template.New("base.html").Funcs(getTplFuncMap()).ParseFiles(partialFiles...))

	pageFiles, _ := filepath.Glob(config.GetRootDir() + "/templates/*.html")
	adminPageFiles, _ := filepath.Glob(config.GetRootDir() + "/templates/admin/*.html")
	pageFiles = append(pageFiles, adminPageFiles...)
	templates = make(map[string]*template.Template)

	var prefix string
	var buffer bytes.Buffer

	for _, tplName := range pageFiles {
		if strings.Contains(tplName, "base.html") {
			continue
		}

		if strings.Contains(tplName, "/admin/") {
			prefix = "admin/"
		} else {
			prefix = ""
		}

		buffer.Reset()
		buffer.WriteString(prefix)
		buffer.WriteString(filepath.Base(tplName))

		templates[buffer.String()] = template.Must(template.Must(baseTemplate.Clone()).ParseFiles(
			tplName,
		))
	}
}
