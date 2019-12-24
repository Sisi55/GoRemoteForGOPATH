package app

import (
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/controllers"
	"net/http"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	//router = gin.New()
}

func StartApp() {
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
