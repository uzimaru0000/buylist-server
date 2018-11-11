package controller

import (
	"github.com/uzimaru0000/buylist/interface/presenter"
	"github.com/uzimaru0000/buylist/usecase/service"
)

type foodController struct {
	Presenter presenter.FoodPresenter
}

type FoodController interface {
	service.FoodService
}

func NewFoodController(presenter presenter.FoodPresenter) FoodController {
	return service.NewFoodService(presenter)
}
