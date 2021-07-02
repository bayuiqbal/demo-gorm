package routes

import (
	"github.com/bayuiqballl/demo/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/people", controllers.Create)
	r.GET("/people", controllers.GetAll)
	r.GET("/people/:id", controllers.GetByID)
	r.DELETE("/people/:id", controllers.Delete)
	r.PUT("/people/:id", controllers.Update)

	return r
}
