package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r * gin.Engine){
	todoRoutes := r.Group("/todos")
	todoRoutes.GET("/", controllers.GetTodos)
	todoRoutes.POST("/", controllers.CreateTodo)
	todoRoutes.PUT("/:id", controllers.UpdateTodo)
	todoRoutes.DELETE("/:id", controllers.DeleteTodo)
}