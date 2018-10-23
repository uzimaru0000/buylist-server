package service

import (
	"time"

	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/usecase/presenter"
	"github.com/uzimaru0000/buylist/usecase/repository"
)

type buyListService struct {
	Repository repository.BuyListRepository
	Presenter  presenter.RecipePresenter
}

type BuyListService interface {
	Create(urls []string) (*model.BuyList, error)
	Get(list *model.BuyList) (*model.BuyList, error)
}

func NewBuyListService(repo repository.BuyListRepository, pre presenter.RecipePresenter) BuyListService {
	return &buyListService{repo, pre}
}

func (service *buyListService) Create(urls []string) (*model.BuyList, error) {
	recipe, err := service.Presenter.Responce(urls)
	if err != nil {
		return nil, err
	}

	list := model.BuyList{
		Ingredients: *recipe,
		CreatedAt:   time.Now(),
	}

	return service.Repository.Store(&list)
}

func (service *buyListService) Get(list *model.BuyList) (*model.BuyList, error) {
	result, err := service.Repository.Find(list)
	if err != nil {
		return nil, err
	}

	return result, nil
}
