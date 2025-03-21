package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.POST("/api/save", SaveData)
	router.GET("/api/data", FetchData)
	router.POST("/api/format", FormatData)
}
