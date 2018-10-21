package presenter

type RecipePresenter interface {
	Responce(url string) (*map[string]string, error)
}
