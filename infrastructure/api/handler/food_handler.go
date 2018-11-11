package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/interface/controller"
)

type foodHandler struct {
	Controller controller.FoodController
}

type FoodHandler interface {
	GetFood(c *gin.Context)
}

func NewFoodHandler(controller controller.FoodController) FoodHandler {
	return &foodHandler{Controller: controller}
}

func (h *foodHandler) GetFood(c *gin.Context) {
	code, err := strconv.Atoi(c.Param("code"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: "Code should be integer."})
		return
	}

	food, err := h.Controller.Get(code)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, food)

}
