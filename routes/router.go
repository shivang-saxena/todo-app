package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/shivang-saxena/todo-app/controllers"
)

// SetupRouter configures and returns the application router
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Create controller instances
	todoController := controllers.NewTodoController(db)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Todo routes
		todos := v1.Group("/todos")
		{
			todos.GET("/", todoController.GetAllTodos)
			todos.GET("/:id", todoController.GetTodo)
			todos.POST("/", todoController.CreateTodo)
			todos.PUT("/:id", todoController.UpdateTodo)
			todos.DELETE("/:id", todoController.DeleteTodo)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return r
}
