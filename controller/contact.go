package controller

import (
	"github.com/alexivanenko/web_cv/config"
	"github.com/alexivanenko/web_cv/model"
	"github.com/gin-gonic/gin"
)

func ContactController(c *gin.Context) {
	developer := new(model.Developer)
	developer.LoadByLogin(config.String("view.nickname"))

	render(c, "contact.html", gin.H{"developer": developer})
}
