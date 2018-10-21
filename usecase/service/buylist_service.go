package service

import (
	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/usecase/presenter"
	"github.com/uzimaru0000/buylist/usecase/repository"
)

type buyListService struct {
	Repository repository.BuyListRepository
	Presenter  presenter.BuyListPresenter
}

type BuyListService interface {
	Create(list *model.BuyList) error
	Get(list *model.BuyList) (*model.BuyList, error)
}

func NewBuyListService(repo repository.BuyListRepository, pre presenter.BuyListPresenter) BuyListService {
	return &buyListService{repo, pre}
}

func (service *buyListService) Create(list *model.BuyList) error {
	return service.Repository.Store(list)
}

func (service *buyListService) Get(list *model.BuyList) (*model.BuyList, error) {
	result, err := service.Repository.Find(list)
	if err != nil {
		return nil, err
	}

	return service.Presenter.Response(result), nil
}
