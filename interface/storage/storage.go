package storage

import (
	"cloud.google.com/go/firestore"
	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/usecase/repository"
	"golang.org/x/net/context"
)

type storage struct {
	firestore *firestore.Client
}

type Storage interface {
	repository.BuyListRepository
}

func NewStorage(firestore *firestore.Client) Storage {
	return &storage{firestore}
}

func (storage *storage) Store(list *model.BuyList) (*model.BuyList, error) {
	ctx := context.Background()
	data := make(map[string][]string)
	data["ingredients"] = list.Ingredients
	ref := storage.firestore.Collection("buylist").NewDoc()
	_, err := ref.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	list.ID = ref.ID
	return list, err
}

func (storage *storage) Find(list *model.BuyList) (*model.BuyList, error) {
	ref := storage.firestore.Collection("buylist").Doc(list.ID)
	ss, err := ref.Get(context.Background())
	if err != nil {
		return nil, err
	}

	data, err := ss.DataAt("ingredients")
	if err != nil {
		return nil, err
	}
	list.Ingredients = convert(data)

	return list, nil
}

func (storage *storage) Update(list *model.BuyList) (*model.BuyList, error) {
	ref := storage.firestore.Collection("buylist").Doc(list.ID)
	_, err := ref.Update(context.Background(), []firestore.Update{{Path: "ingredients", Value: list.Ingredients}})
	if err != nil {
		return nil, err
	}
	return list, err
}

func (storage *storage) Delete(list *model.BuyList) error {
	ref := storage.firestore.Collection("buylist").Doc(list.ID)
	_, err := ref.Delete(context.Background())

	return err
}

func convert(data interface{}) []string {
	result := make([]string, 0)
	for _, val := range data.([]interface{}) {
		result = append(result, val.(string))
	}

	return result
}
