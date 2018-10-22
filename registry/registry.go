package registry

import (
	"context"

	"firebase.google.com/go"
	"github.com/uzimaru0000/buylist/infrastructure/api/handler"
	"github.com/uzimaru0000/buylist/interface/client"
	"github.com/uzimaru0000/buylist/interface/controller"
	"github.com/uzimaru0000/buylist/interface/perser"
	"github.com/uzimaru0000/buylist/interface/presenter"
	"github.com/uzimaru0000/buylist/interface/storage"
)

type interactor struct {
	firebase *firebase.App
	units    []string
}

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(app *firebase.App, units []string) Interactor {
	return &interactor{app, units}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return i.newBuyListHandler()
}

func (i *interactor) newBuyListHandler() handler.BuyListHandler {
	return handler.NewBuyListHandler(i.newBuyListController())
}

func (i *interactor) newBuyListController() controller.BuyListController {
	return controller.NewBuyListController(i.newStorage(), i.newRecipePresenter())
}

func (i *interactor) newStorage() storage.Storage {
	client, err := i.firebase.Firestore(context.Background())
	if err != nil {
		panic(err)
	}

	return storage.NewStorage(client)
}

func (i *interactor) newRecipePresenter() presenter.RecipePresenter {
	return presenter.NewRecipePresenter(i.newRecipeClient(), i.newPerser())
}

func (i *interactor) newRecipeClient() client.RecipeClient {
	return client.NewClient()
}

func (i *interactor) newPerser() perser.Perser {
	return perser.NewPerser(i.units)
}
