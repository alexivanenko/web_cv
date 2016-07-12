package controller

import (
	"fmt"

	"github.com/alexivanenko/web_cv/config"
	"github.com/gin-gonic/gin"
)

func DownloadController(c *gin.Context) {
	c.File(fmt.Sprintf("%s/files/AlexanderIvanenko.pdf", config.GetRootDir()))
}
