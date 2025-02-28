package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kubeblog/backend/database"
	"github.com/kubeblog/backend/user"
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

		// Create user if not exists in database
		userRepository := user.NewUserRepository(database.Db)
		currentUser, err := userRepository.GetUserByEmail(payload.Claims["email"].(string))
		if err != nil {
			currentUser = &user.User{
				ID:          uuid.New(),
				Email:       payload.Claims["email"].(string),
				Username:    payload.Claims["email"].(string),
				DisplayName: payload.Claims["name"].(string),
			}
			if err := userRepository.CreateUser(currentUser); err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create user"})
			}
		}

		// Store user in the request context
		c.Set("user", currentUser)

		return next(c)
	}
}
