package article

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ArticleService struct {
	repository *ArticleRepository
}

func NewArticleService(r *ArticleRepository) *ArticleService {
	return &ArticleService{repository: r}
}

func (h *ArticleService) Route(e *echo.Echo) {
	e.GET("/articles", h.GetAllArticles)
	e.POST("/articles", h.CreateArticle)
	e.GET("/articles/:id", h.GetArticleByID)
	e.PUT("/articles/:id", h.UpdateArticle)
	e.DELETE("/articles/:id", h.DeleteArticle)
}

func (h *ArticleService) GetAllArticles(c echo.Context) error {
	articles, err := h.repository.GetAllArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleService) CreateArticle(c echo.Context) error {
	var article Article
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.repository.CreateArticle(&article); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, article)
}

func (h *ArticleService) GetArticleByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	article, err := h.repository.GetArticleByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, article)
}

func (h *ArticleService) UpdateArticle(c echo.Context) error {
	var article Article
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.repository.UpdateArticle(&article); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, article)
}

func (h *ArticleService) DeleteArticle(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.repository.DeleteArticle(id); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.NoContent(http.StatusNoContent)
}
