package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilmsg/eloop-api/internal/router"
)

func main() {
	app := gin.Default()
	router.RouterIndex(app)
	router.RouterTodo(app)

	app.Run(":3090")
}
