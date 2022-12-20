package router

import (
	"github.com/gin-gonic/gin"
)

func RouterIndex(app *gin.Engine) {
	index := app.Group("/")

	index.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "api",
		})
	})

	index.GET("/hi", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "hi",
		})
	})
}
