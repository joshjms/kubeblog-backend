package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kubeblog/backend/models"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	db.AutoMigrate(&models.Article{})

	return &ArticleRepository{DB: db}
}

func (r *ArticleRepository) CreateArticle(article *models.Article) error {
	return r.DB.Create(article).Error
}

func (r *ArticleRepository) GetArticleByID(id uuid.UUID) (*models.Article, error) {
	var article models.Article
	err := r.DB.Preload("Author").First(&article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.DB.Preload("Author").Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *ArticleRepository) GetArticleByAuthorID(id uuid.UUID) ([]models.Article, error) {
	var articles []models.Article
	err := r.DB.Preload("Author").Find(&articles, "author_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *ArticleRepository) GetFeaturedArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.DB.Where("featured = ?", true).Preload("Author").Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *ArticleRepository) UpdateArticle(article *models.Article) error {
	return r.DB.Save(article).Error
}

func (r *ArticleRepository) DeleteArticle(id uuid.UUID) error {
	result := r.DB.Delete(&models.Article{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return errors.New("article not found")
	}
	return nil
}
