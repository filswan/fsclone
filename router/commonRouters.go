package router

import (
	"github.com/gin-gonic/gin"
	"migrates3/domain"
	"migrates3/service"
	"net/http"
)

func CommonManager(router *gin.RouterGroup) {
	image := router.Group("/info")
	image.GET("/version", GetMigratinVersion)
}

func GetMigratinVersion(c *gin.Context) {
	c.JSON(http.StatusOK, domain.NewSuccessResponse(service.GetVersion()))
}
