package router

import (
	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/buylist/infrastructure/api/handler"
)

func NewRoute(engine *gin.Engine, app *firebase.App, handler handler.AppHandler) {
	// v1 := engine.Group("/api/v1", middleware.AuthMiddleWare(app))
	v1 := engine.Group("/api/v1")
	{
		v1.POST("/list", handler.CreateBuyList)
		v1.GET("/list/:id", handler.GetBuyList)
	}
}
