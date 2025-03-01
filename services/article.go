package services

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kubeblog/backend/auth"
	"github.com/kubeblog/backend/middleware"
	"github.com/kubeblog/backend/models"
	"github.com/kubeblog/backend/repositories"
	"github.com/labstack/echo/v4"
)

type ArticleService struct {
	repository *repositories.ArticleRepository
}

func NewArticleService(r *repositories.ArticleRepository) *ArticleService {
	return &ArticleService{repository: r}
}

func (h *ArticleService) Route(e *echo.Echo, authMiddleware *middleware.AuthMiddleware) {
	e.GET("/articles", h.GetAllArticles)
	e.GET("/articles/:id", h.GetArticleByID)
	e.GET("/articles/featured", h.GetFeaturedArticles)

	// Protected routes
	e.POST("/articles", h.CreateArticle, authMiddleware.ValidateGoogleTokenMiddleware)
	e.PUT("/articles/:id", h.UpdateArticle, authMiddleware.ValidateGoogleTokenMiddleware)
	e.DELETE("/articles/:id", h.DeleteArticle, authMiddleware.ValidateGoogleTokenMiddleware)
}

func (h *ArticleService) GetAllArticles(c echo.Context) error {
	articles, err := h.repository.GetAllArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleService) CreateArticle(c echo.Context) error {
	var article models.Article
	article.ID = uuid.New()
	article.AuthorID = c.Get("user").(*auth.User).ID
	article.CreatedAt = time.Now().Unix()
	article.UpdatedAt = time.Now().Unix()

	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Store image at /uploads from base64 string
	if article.Image != "" {
		filename := uuid.New().String()
		filePath := filepath.Join("uploads", filename)
		if err := saveBase64Image(article.Image, filePath); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		article.Image = fmt.Sprintf("http://localhost:8080/%s", filePath)
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

func (h *ArticleService) GetArticleByAuthorID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	articles, err := h.repository.GetArticleByAuthorID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleService) GetFeaturedArticles(c echo.Context) error {
	articles, err := h.repository.GetFeaturedArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleService) UpdateArticle(c echo.Context) error {
	var article models.Article
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

func saveBase64Image(base64Image, filePath string) error {
	// Remove base64 metadata if present (e.g., "data:image/png;base64,")
	if idx := strings.Index(base64Image, ","); idx != -1 {
		base64Image = base64Image[idx+1:]
	}

	// Decode base64 string
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return fmt.Errorf("failed to decode base64 image: %w", err)
	}

	// Create the file
	err = os.WriteFile(filePath, imageData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}
