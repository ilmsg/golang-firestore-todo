package functions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilmsg/eloop-api/internal/router"
)

var app *gin.Engine

func init() {
	app = gin.Default()
	router.RouterIndex(app)
	router.RouterTodo(app)
}

func EloopApi(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
