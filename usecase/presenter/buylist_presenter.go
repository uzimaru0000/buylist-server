package presenter

import (
	"github.com/uzimaru0000/buylist/domain/model"
)

type BuyListPresenter interface {
	Response(list *model.BuyList) *model.BuyList
}
