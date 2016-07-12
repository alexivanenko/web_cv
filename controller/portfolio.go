package controller

import (
	"github.com/alexivanenko/web_cv/config"
	"github.com/alexivanenko/web_cv/model"
	"github.com/gin-gonic/gin"
)

func PortfolioController(c *gin.Context) {
	portfolio := new(model.Portfolio)
	portfolio.LoadByLogin(config.String("view.nickname"))

	render(c, "portfolio.html", gin.H{"portfolio": portfolio, "categoriesList": portfolio.GetCategoriesList()})
}
