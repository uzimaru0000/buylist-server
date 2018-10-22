package presenter

import (
	"strconv"

	"github.com/uzimaru0000/buylist/interface/client"
	"github.com/uzimaru0000/buylist/interface/perser"
	"github.com/uzimaru0000/buylist/usecase/presenter"
)

type recipePresenter struct {
	client client.RecipeClient
	perser perser.Perser
}

type RecipePresenter interface {
	presenter.RecipePresenter
}

func NewRecipePresenter(client client.RecipeClient, perser perser.Perser) RecipePresenter {
	return &recipePresenter{client, perser}
}

func (presenter *recipePresenter) Responce(url string) (*map[string]string, error) {
	strs, err := presenter.client.Get(url)
	if err != nil {
		return nil, err
	}

	data := presenter.perser.Perse(strs)
	ingre := make(map[string]string)
	for key, val := range data {
		ingre[key] = strconv.FormatInt(int64(val.Amount), 10) + val.Unit
	}

	return &ingre, nil
}
