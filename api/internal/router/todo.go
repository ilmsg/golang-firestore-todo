package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilmsg/eloop-api/internal/firebase"
	"github.com/ilmsg/eloop-api/internal/repository"
	"github.com/ilmsg/eloop-api/pkg/model"
)

func RouterTodo(app *gin.Engine) {
	client := firebase.NewFirestore()
	todoRepo := repository.NewTodoRepo(client)

	todo := app.Group("/todo")
	todo.GET("/", func(ctx *gin.Context) {
		todos, err := todoRepo.GetTodos()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"todos": todos,
		})
	})

	todo.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		todo, err := todoRepo.GetTodo(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"todo": todo,
		})
	})

	todo.PUT("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var todo *model.Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := todoRepo.UpdateTodo(id, todo); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"todo": todo,
		})
	})

	todo.POST("/", func(ctx *gin.Context) {
		var todo *model.Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := todoRepo.CreateTodo(todo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		todo.ID = id
		ctx.JSON(200, gin.H{
			"todo": todo,
		})
	})

	todo.DELETE("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := todoRepo.DeleteTodo(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"msg": "delete success.",
		})
	})
}
