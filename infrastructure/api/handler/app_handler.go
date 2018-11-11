package handler

type AppHandler interface {
	BuyListHandler
	FoodHandler
}

type appHandler struct {
	BuyListHandler
	FoodHandler
}

func NewAppHandler(buylistHandler BuyListHandler, foodHandler FoodHandler) AppHandler {
	return &appHandler{BuyListHandler: buylistHandler, FoodHandler: foodHandler}
}
