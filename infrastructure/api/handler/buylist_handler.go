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
	URLs        []string `json:"urls"`
	Ingredients []string `json:"ingredients"`
}

type BuyListHandler interface {
	CreateBuyList(c *gin.Context)
	GetBuyList(c *gin.Context)
	AddList(c *gin.Context)
	DeleteList(c *gin.Context)
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
	list, err = handler.controllre.Store(list)
	if err != nil {
		log.Printf("Store Error : %s\n", err.Error())
		c.JSON(500, model.ResponseError{Message: "Database stored error"})
	}

	c.JSON(http.StatusOK, list)
}

func (handler *buylistHandler) GetBuyList(c *gin.Context) {
	id := c.Param("id")

	list := &model.BuyList{ID: id}
	list, err := handler.controllre.Get(list)
	if err != nil {
		log.Printf("Get Error : %s\n", err.Error())
		c.JSON(http.StatusNotFound, model.ResponseError{Message: "List getting error"})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (handler *buylistHandler) AddList(c *gin.Context) {
	request := &request{}

	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: "Request is not in the correct json format."})
		return
	}

	id := c.Param("id")
	base := &model.BuyList{ID: id}
	base, err := handler.controllre.Get(base)
	if err != nil {
		log.Printf("Get Error : %s\n", err.Error())
		c.JSON(http.StatusNotFound, model.ResponseError{Message: "ID is not found."})
		return
	}

	base, err = handler.controllre.Add(base, request.Ingredients)
	if err != nil {
		log.Printf("Get Error : %s\n", err.Error())
		c.JSON(500, model.ResponseError{Message: "Update Error"})
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
		c.JSON(500, model.ResponseError{Message: "Update Error"})
		return
	}

	base, err = handler.controllre.Merge(base, list)
	if err != nil {
		log.Printf("Update Error : %s\n", err.Error())
		c.JSON(500, model.ResponseError{Message: "Update Error"})
		return
	}

	c.JSON(http.StatusOK, base)
}

func (handler *buylistHandler) DeleteList(c *gin.Context) {
	id := c.Param("id")
	list := &model.BuyList{ID: id}
	err := handler.controllre.Delete(list)
	if err != nil {
		log.Printf("Delete Error : %s\n", err.Error())
		c.JSON(500, model.ResponseError{Message: "Delete Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Delete Success"})
}
