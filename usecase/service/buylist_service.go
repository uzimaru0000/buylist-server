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
	Store(list *model.BuyList) (*model.BuyList, error)
	Get(list *model.BuyList) (*model.BuyList, error)
	Add(base *model.BuyList, ingredients []string) (*model.BuyList, error)
	Merge(base *model.BuyList, list *model.BuyList) (*model.BuyList, error)
	Delete(list *model.BuyList) error
}

func NewBuyListService(repo repository.BuyListRepository, pre presenter.RecipePresenter) BuyListService {
	return &buyListService{repo, pre}
}

func (service *buyListService) Create(urls []string) (*model.BuyList, error) {
	recipe, err := service.Presenter.Response(urls)
	if err != nil {
		return nil, err
	}

	list := &model.BuyList{
		Ingredients: recipe,
		CreatedAt:   time.Now(),
	}

	return list, nil
}

func (service *buyListService) Store(list *model.BuyList) (*model.BuyList, error) {
	return service.Repository.Store(list)
}

func (service *buyListService) Get(list *model.BuyList) (*model.BuyList, error) {
	result, err := service.Repository.Find(list)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *buyListService) Add(base *model.BuyList, ingredients []string) (*model.BuyList, error) {
	base.Ingredients = append(base.Ingredients, ingredients...)

	return service.Repository.Update(base)
}

func (service *buyListService) Merge(base *model.BuyList, list *model.BuyList) (*model.BuyList, error) {
	base.Ingredients = append(base.Ingredients, list.Ingredients...)

	return service.Repository.Update(base)
}

func (service *buyListService) Delete(list *model.BuyList) error {
	return service.Repository.Delete(list)
}
