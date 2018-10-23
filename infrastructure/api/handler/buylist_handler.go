package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/interface/controller"
)

type buylistHandler struct {
	controllre controller.BuyListController
}

type request struct {
	URLs []string `json:"urls"`
}

type BuyListHandler interface {
	CreateBuyList(c *gin.Context)
	GetBuyList(c *gin.Context)
}

func NewBuyListHandler(controller controller.BuyListController) BuyListHandler {
	return &buylistHandler{controller}
}

func (handler *buylistHandler) CreateBuyList(c *gin.Context) {
	request := &request{}

	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: "Request is not in the correct json format."})
		return
	}

	for _, url := range request.URLs {
		if !strings.HasPrefix(url, "https://") {
			c.JSON(http.StatusBadRequest, model.ResponseError{Message: "Request is not in the correct URL format."})
			return
		}
	}

	list, err := handler.controllre.Create(request.URLs)
	if err != nil {
		log.Printf("Create Error : %s\n", err.Error())
		c.JSON(500, model.ResponseError{Message: "List create error"})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (handler *buylistHandler) GetBuyList(c *gin.Context) {
	id := c.Param("id")

	list := &model.BuyList{ID: id}
	list, err := handler.controllre.Get(list)
	if err != nil {
		log.Printf("Get Error : %s\n", err.Error())
		c.JSON(500, model.ResponseError{Message: "List getting error"})
		return
	}

	c.JSON(http.StatusOK, list)
}
