package presenter

import (
	"github.com/uzimaru0000/buylist/domain/model"
)

type FoodPresenter interface {
	Get(int) (*model.Food, error)
}
