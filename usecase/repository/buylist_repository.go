package repository

import "github.com/uzimaru0000/buylist/domain/model"

type BuyListRepository interface {
	Store(list *model.BuyList) (*model.BuyList, error)
	Find(list *model.BuyList) (*model.BuyList, error)
	Update(list *model.BuyList) (*model.BuyList, error)
	Delete(list *model.BuyList) error
}
