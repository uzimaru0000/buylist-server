package service

import (
	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/usecase/presenter"
)

type foodService struct {
	Presenter presenter.FoodPresenter
}

type FoodService interface {
	Get(int) (*model.Food, error)
}

func (s *foodService) Get(code int) (*model.Food, error) {
	return s.Presenter.Get(code)
}

func NewFoodService(presenter presenter.FoodPresenter) FoodService {
	return &foodService{Presenter: presenter}
}
