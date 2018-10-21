package main

import (
	"context"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}

	router := gin.Default()

	config := cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders: []string{"Authorization"},
	}

	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	{
		private := v1.Group("/private", FirebaseAuthMiddleWear(app))
		{
			private.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Success"})
			})
		}

		public := v1.Group("/public")
		{
			public.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Success"})
			})
		}
	}

	router.Run(":5000")
}

// FirebaseAuthMiddleWear Authorized from Firebase
func FirebaseAuthMiddleWear(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		header := c.GetHeader("Authorization")
		idToken := strings.Replace(header, "Bearer ", "", 1)

		token, err := auth.VerifyIDToken(context.Background(), idToken)

		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorization error"})
			c.Abort()
		}

		log.Printf("Verify ID token: %v\n", token)
		c.Next()
	}
}
