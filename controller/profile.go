package controller

import (
	"github.com/alexivanenko/web_cv/config"
	"github.com/alexivanenko/web_cv/model"
	"github.com/gin-gonic/gin"
)

func ProfileController(c *gin.Context) {
	profile := new(model.Profile)
	profile.LoadByLogin(config.String("view.nickname"))

	render(c, "profile.html", gin.H{"profile": profile})
}
