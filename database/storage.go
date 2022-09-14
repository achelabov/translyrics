package database

import (
	"context"

	"github.com/achelabov/translyrics/models"
)

type ArticleStorage interface {
	CreateArticle(ctx context.Context, article *models.Article, user *models.User) error
	GetArticles(ctx context.Context) ([]*models.Article, error)
	GetArticlesByUserID(ctx context.Context, user *models.User) ([]*models.Article, error)
	GetArticleByID(ctx context.Context, id string) (*models.Article, error)
	UpdateArticle(ctx context.Context, newArticle *models.Article, user *models.User, id string) error
	DeleteArticle(ctx context.Context, user *models.User, id string) error
}

type UserStorage interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}
