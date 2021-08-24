package router

import (
	"fmt"
	"fsclone/cloud/aws"
	"fsclone/cloud/google"
	"fsclone/cloud/microsoft"
	"fsclone/domain"
	"fsclone/logs"
	"fsclone/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RCloneManager(router *gin.RouterGroup) {
	image := router.Group("/sync")
	image.GET("/", GetConnection)
	image.POST("/GS2SyncAWS3", GS2SyncAWS3)
	image.POST("/GS2SyncOneDrive", GS2SyncOneDrive)
	image.POST("/GS2SyncGoogleDrive", GS2SyncGoogleDrive)
}

func GetConnection(c *gin.Context) {
	c.JSON(http.StatusOK, domain.NewSuccessResponse("DataSync service connected!"))
}

func GS2SyncAWS3(c *gin.Context) {
	var gs2AWS domain.AWStoGS2
	err := c.BindJSON(&gs2AWS)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Resolve AWS path failed! "))
		return
	}

	if len(gs2AWS.GS2.Bucket) <= 0 || len(gs2AWS.AWStest.Bucket) <= 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("AWS or GS2 Bucket name not Found! "))
	} else {
		err = service.CreateAWStoGS2Conf(gs2AWS)
		if err != nil {
			fmt.Println("err:\t", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Create AWS conf file failed!2 "))
			return
		}
		err = aws.SyncDataWithAWS3(gs2AWS.GS2.Bucket, gs2AWS.AWStest.Bucket, gs2AWS.GS2.ConfName, gs2AWS.AWStest.ConfName)
		if err != nil {
			fmt.Println("err:\t", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Sync with AWS files failed! "))
			return
		}
		c.JSON(http.StatusOK, domain.NewSuccessResponse("Sync AWS bucket success!"))
	}
}
func GS2SyncOneDrive(c *gin.Context) {

	var gs2OneDrive domain.OneDrivetoGS2
	err := c.BindJSON(&gs2OneDrive)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Resolve One Drive path failed! "))
		return
	}

	if len(gs2OneDrive.GS2.Bucket) <= 0 || len(gs2OneDrive.OneDrive.Bucket) <= 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("One Drive or GS2 Bucket name not Found! "))
	} else {
		err = service.CreateOneDrivetoGS2Conf(gs2OneDrive)
		if err != nil {
			fmt.Println("err:\t", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Create One Drive conf file failed! "))
			return
		}

		err = microsoft.SyncDataWithOneDrive(gs2OneDrive.GS2.Bucket, gs2OneDrive.OneDrive.Bucket, gs2OneDrive.GS2.ConfName, gs2OneDrive.OneDrive.ConfName)
		if err != nil {
			fmt.Println("err:\t", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Sync with OneDrive files failed! "))
			return
		}
		c.JSON(http.StatusOK, domain.NewSuccessResponse("Sync One Drive bucket success!"))
	}
}
func GS2SyncGoogleDrive(c *gin.Context) {
	var gs2GoogleDrive domain.GoogleDrivetoGS2

	err := c.BindJSON(&gs2GoogleDrive)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Resolve Google Drive path failed! "))
		return
	}

	if len(gs2GoogleDrive.GS2.Bucket) <= 0 || len(gs2GoogleDrive.GoogleDrive.Bucket) <= 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Google Drive or GS2 Bucket name not Found! "))
	} else {
		err = service.CreateGoogleDrivetoGS2Conf(gs2GoogleDrive)
		if err != nil {
			fmt.Println("err:\t", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Create Google Drive conf file failed! "))
			return
		}
		err = google.SyncDataWithGoogleDrive(gs2GoogleDrive.GS2.Bucket, gs2GoogleDrive.GoogleDrive.Bucket, gs2GoogleDrive.GS2.ConfName, gs2GoogleDrive.GoogleDrive.ConfName)
		if err != nil {
			fmt.Println("err:\t", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewErrorResponse("Sync with GoogleDrive files failed! "))
			return
		}
		c.JSON(http.StatusOK, domain.NewSuccessResponse("Sync Google Drive bucket success!"))
	}
}
