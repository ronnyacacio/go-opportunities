package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ronnyacacio/go-opportunities/handlers"
)

func Initialize() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/openings", handlers.ListOpeningsHandler)
		v1.POST("/openings", handlers.CreateOpeningsHandler)
	}

	r.Run(":3000")
}
