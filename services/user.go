package services

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kubeblog/backend/middleware"
	"github.com/kubeblog/backend/models"
	"github.com/kubeblog/backend/repositories"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{userRepository: r}
}

func (h *UserService) Route(e *echo.Echo, authMiddleware *middleware.AuthMiddleware) {
	e.POST("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUserByID)
	e.GET("/users/email/:email", h.GetUserByEmail)
	e.GET("/users/username/:username", h.GetUserByUsername, authMiddleware.ValidateGoogleTokenMiddleware)
	e.PUT("/users/:id", h.UpdateUser, authMiddleware.ValidateGoogleTokenMiddleware)
	e.DELETE("/users/:id", h.DeleteUser, authMiddleware.ValidateGoogleTokenMiddleware)
	e.GET("/users/me", h.GetCurrentUserInfo, authMiddleware.ValidateGoogleTokenMiddleware)
}

func (h *UserService) GetCurrentUserInfo(c echo.Context) error {
	user := c.Get("user").(*models.User)
	return c.JSON(http.StatusOK, user)
}

func (h *UserService) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.userRepository.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserService) GetUserByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := h.userRepository.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserService) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")

	user, err := h.userRepository.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserService) GetUserByUsername(c echo.Context) error {
	username := c.Param("username")

	user, err := h.userRepository.GetUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserService) UpdateUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user.ID = id
	if err := h.userRepository.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserService) DeleteUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.userRepository.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}
