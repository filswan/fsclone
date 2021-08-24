package router

import (
	"fsclone/domain"
	"fsclone/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommonManager(router *gin.RouterGroup) {
	image := router.Group("/info")
	image.GET("/version", GetMigratinVersion)
}

func GetMigratinVersion(c *gin.Context) {
	c.JSON(http.StatusOK, domain.NewSuccessResponse(service.GetVersion()))
}
