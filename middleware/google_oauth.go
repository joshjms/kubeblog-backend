package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

func ValidateGoogleTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing authorization header"})
		}

		token := authHeader[len("Bearer "):]

		payload, err := idtoken.Validate(context.Background(), token, "257566191758-cd6jv7m3st9dumkgoqa9bn8eqae53pk0.apps.googleusercontent.com")
		if err != nil {
			log.Println("Invalid token:", err)
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
		}

		// Store user information in the request context
		c.Set("user", payload)

		return next(c)
	}
}
