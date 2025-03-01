package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/kubeblog/backend/models"
	"github.com/kubeblog/backend/repositories"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

type AuthMiddleware struct {
	userRepository *repositories.UserRepository
}

func NewAuthMiddleware(r *repositories.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		userRepository: r,
	}
}

func (mw *AuthMiddleware) ValidateGoogleTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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

		user, err := mw.userRepository.GetUserByEmail(payload.Claims["email"].(string))
		if err != nil {
			user = models.NewUser(payload.Claims["email"].(string), payload.Claims["name"].(string))
			if err := mw.userRepository.CreateUser(user); err != nil {
				log.Println("Failed to create user:", err)
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create user"})
			}
		}

		c.Set("user", user)

		return next(c)
	}
}
