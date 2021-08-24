package main

import (
	"fsclone/conf"
	"fsclone/logs"
	"fsclone/router"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"time"
)

func main() {
	conf.Init()
	logs.InitLogger()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	v1 := r.Group("/api/v1")
	router.RCloneManager(v1.Group("/migration"))
	router.CommonManager(v1.Group("/common"))
	err := r.Run(":8083")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}
