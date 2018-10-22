package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/buylist/infrastructure/api/router"
	"github.com/uzimaru0000/buylist/registry"
	"google.golang.org/api/option"
)

func main() {
	// Firebase App initialize
	opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}

	// gin Engine initialize
	engine := gin.Default()

	config := cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders: []string{"Authorization"},
	}

	engine.Use(cors.New(config))

	// interacter initialize
	units := []string{
		"g",
		"cc",
		"個",
		"個分",
	}
	interacter := registry.NewInteractor(app, units)

	// handler initialize
	handler := interacter.NewAppHandler()

	router.NewRoute(engine, app, handler)

	engine.Run(":5000")
}
