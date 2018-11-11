package registry

import (
	"context"

	"firebase.google.com/go"
	"github.com/uzimaru0000/buylist/infrastructure/api/handler"
	"github.com/uzimaru0000/buylist/interface/client"
	"github.com/uzimaru0000/buylist/interface/controller"
	"github.com/uzimaru0000/buylist/interface/presenter"
	"github.com/uzimaru0000/buylist/interface/storage"
)

type interactor struct {
	firebase    *firebase.App
	yahooAPIKey string
}

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(app *firebase.App, apiKey string) Interactor {
	return &interactor{firebase: app, yahooAPIKey: apiKey}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return handler.NewAppHandler(i.newBuyListHandler(), i.newFoodHandler())
}

func (i *interactor) newBuyListHandler() handler.BuyListHandler {
	return handler.NewBuyListHandler(i.newBuyListController())
}

func (i *interactor) newFoodHandler() handler.FoodHandler {
	return handler.NewFoodHandler(i.newFoodController())
}

func (i *interactor) newBuyListController() controller.BuyListController {
	return controller.NewBuyListController(i.newStorage(), i.newRecipePresenter())
}

func (i *interactor) newFoodController() controller.FoodController {
	return controller.NewFoodController(i.newFoodPresenter())
}

func (i *interactor) newStorage() storage.Storage {
	client, err := i.firebase.Firestore(context.Background())
	if err != nil {
		panic(err)
	}

	return storage.NewStorage(client)
}

func (i *interactor) newRecipePresenter() presenter.RecipePresenter {
	return presenter.NewRecipePresenter(i.newRecipeClient())
}

func (i *interactor) newFoodPresenter() presenter.FoodPresenter {
	return presenter.NewFoodPresenter(i.yahooAPIKey)
}

func (i *interactor) newRecipeClient() client.RecipeClient {
	return client.NewClient()
}
