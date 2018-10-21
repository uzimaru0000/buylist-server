package storage_test

import (
	"context"
	"os"
	"testing"

	"firebase.google.com/go"
	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/interface/storage"
	"google.golang.org/api/option"
)

func TestStore(t *testing.T) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		t.Fatalf("Firebase app create error : %v\n", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		t.Fatalf("Firestore create error : %v\n", err)
	}

	storage := storage.NewStorage(client)

	list := model.BuyList{
		Ingredients: map[string]string{
			"かぼちゃ":  "300g",
			"生クリーム": "100cc",
		},
	}

	err = storage.Store(ctx, &list)
	if err != nil {
		t.Fatalf("Firestora store err : %v\n", err)
	}
}
