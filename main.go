package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/buylist/config"
	"github.com/uzimaru0000/buylist/infrastructure/api/router"
	"github.com/uzimaru0000/buylist/registry"
	"google.golang.org/api/option"
)

func main() {
	// Firebase App initialize
	opt := option.WithCredentialsFile(config.Get().Firebase.AcountKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}

	// gin Engine initialize
	engine := gin.Default()

	// interacter initialize
	interacter := registry.NewInteractor(app)

	// handler initialize
	handler := interacter.NewAppHandler()

	router.NewRoute(engine, app, handler)

	engine.Run(config.Get().Server.Port)
}
