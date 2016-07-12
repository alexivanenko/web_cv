package main

import (
	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/alexivanenko/web_cv/config"
	"github.com/alexivanenko/web_cv/controller"
	"github.com/alexivanenko/web_cv/model"
)

func main() {
	defer model.GetSession().Close()

	if config.Is("debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/", controller.IndexController)
	router.GET("/profile", controller.ProfileController)
	router.GET("/resume", controller.ResumeController)
	router.GET("/portfolio", controller.PortfolioController)
	router.GET("/contact", controller.ContactController)
	router.GET("/cv_pdf", controller.DownloadController)

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		config.String("auth.login"): config.String("auth.password"),
	}))

	authorized.GET("/", controller.AdminController)
	authorized.POST("/", controller.AdminController)
	authorized.GET("/profile", controller.AdminController)
	authorized.POST("/profile", controller.AdminController)
	authorized.GET("/resume", controller.AdminController)
	authorized.POST("/resume", controller.AdminController)
	authorized.GET("/portfolio", controller.AdminController)
	authorized.POST("/portfolio", controller.AdminController)

	//async
	authorized.GET("/remove_project", controller.RemoveProject)
	authorized.GET("/get_project_data", controller.LoadProject)

	router.NoRoute(controller.IndexController)

	router.Static("/static", fmt.Sprintf("%s/static", config.GetRootDir()))

	host := config.String("addr")
	config.Log(fmt.Sprintf("Running %s", host))

	router.Run(host)
}
