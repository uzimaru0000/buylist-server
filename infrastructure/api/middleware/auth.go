package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		auth, err := app.Auth(ctx)
		if err != nil {
			panic(err)
		}

		header := c.GetHeader("Authorization")
		idToken := strings.Replace(header, "Bearer ", "", 1)

		_, err = auth.VerifyIDToken(ctx, idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
		}

		c.Next()
	}
}
