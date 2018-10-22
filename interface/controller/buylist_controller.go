package controller

import (
	"github.com/uzimaru0000/buylist/interface/storage"
	pre "github.com/uzimaru0000/buylist/usecase/presenter"
	"github.com/uzimaru0000/buylist/usecase/service"
)

type buyListController struct {
	Storage   storage.Storage
	Presenter pre.RecipePresenter
}

type BuyListController interface {
	service.BuyListService
}

func NewBuyListController(storage storage.Storage, pre pre.RecipePresenter) BuyListController {
	return service.NewBuyListService(storage, pre)
}
