package service

import (
	"github.com/uzimaru0000/buylist/usecase/presenter"
)

type recipeService struct {
	Presenter presenter.RecipePresenter
}

type RecipeService interface {
	Get(url string) (*map[string]string, error)
}

func NewRecipeService(pre presenter.RecipePresenter) RecipeService {
	return &recipeService{pre}
}

func (service *recipeService) Get(url string) (*map[string]string, error) {
	recipe, err := service.Presenter.Responce(url)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
