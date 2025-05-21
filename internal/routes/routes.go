package routes

import (
	"github.com/armanesenbaiuly/go-rest-gorm/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	routes := r.Group("api/v1/tasks")
	{
		routes.POST("/", handlers.Create)
		routes.PUT("/:id", handlers.Update)
		routes.GET("/", handlers.GetAll)
		routes.GET("/:id", handlers.GetById)
		routes.DELETE("/:id", handlers.DeleteById)
	}

	return r
}
