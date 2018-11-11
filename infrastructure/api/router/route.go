package router

import (
	"firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/buylist/infrastructure/api/handler"
)

func NewRoute(engine *gin.Engine, app *firebase.App, handler handler.AppHandler) {
	config := cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders: []string{"Authorization"},
	}

	engine.Use(cors.New(config))
	// v1 := engine.Group("/api/v1", middleware.AuthMiddleWare(app))
	v1 := engine.Group("/api/v1")
	{
		v1.POST("/list", handler.CreateBuyList)
		v1.GET("/list/:id", handler.GetBuyList)
		v1.PATCH("/list/:id", handler.AddList)
		v1.DELETE("/list/:id", handler.DeleteList)
		v1.GET("/food/:code", handler.GetFood)
	}
}
