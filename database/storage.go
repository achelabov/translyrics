package database

import (
	"context"

	"github.com/achelabov/translyrics/models"
)

type ArticleStorage interface {
	CreateArticle(ctx context.Context, user *models.User, article *models.Article) error
	GetArticles(ctx context.Context) ([]*models.Article, error)
	GetArticlesByUserID(ctx context.Context, user *models.User) ([]*models.Article, error)
	GetArticlesByID(ctx context.Context, id string) (*models.Article, error)
	UpdateArticle(ctx context.Context, newArticle *models.Article, id string) error
	DeleteArticle(ctx context.Context, id string) error
}

type UserStorage interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}