package article

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{DB: db}
}

func (r *ArticleRepository) CreateArticle(article *Article) error {
	return r.DB.Create(article).Error
}

func (r *ArticleRepository) GetArticleByID(id uuid.UUID) (*Article, error) {
	var article Article
	err := r.DB.First(&article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) GetAllArticles() ([]Article, error) {
	var articles []Article
	err := r.DB.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *ArticleRepository) UpdateArticle(article *Article) error {
	return r.DB.Save(article).Error
}

func (r *ArticleRepository) DeleteArticle(id uuid.UUID) error {
	result := r.DB.Delete(&Article{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return errors.New("article not found")
	}
	return nil
}
