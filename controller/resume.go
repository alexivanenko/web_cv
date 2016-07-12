package controller

import (
	"github.com/alexivanenko/web_cv/config"
	"github.com/alexivanenko/web_cv/model"
	"github.com/gin-gonic/gin"
)

func ResumeController(c *gin.Context) {
	resume := new(model.Resume)
	resume.LoadByLogin(config.String("view.nickname"))

	render(c, "resume.html", gin.H{"resume": resume})
}
