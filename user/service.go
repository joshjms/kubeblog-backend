package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
	return &UserService{userRepository: r}
}

func (h *UserService) Route(e *echo.Echo) {
	e.POST("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUserByID)
	e.GET("/users/email/:email", h.GetUserByEmail)
	e.GET("/users/username/:username", h.GetUserByUsername)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
}

func (h *UserService) CreateUser(c echo.Context) error {
	var user User
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

	var user User
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
