package localstorage

import (
	"context"
	"sync"

	"github.com/achelabov/translyrics/models"
)

type ArticleLocalStorage struct {
	articles map[string]*models.Article
	mutex    *sync.Mutex
}

func NewArticleLocalStorage() *ArticleLocalStorage {
	return &ArticleLocalStorage{
		articles: make(map[string]*models.Article),
		mutex:    new(sync.Mutex),
	}
}

func (s *ArticleLocalStorage) CreateArticle(ctx context.Context, article *models.Article, user *models.User) error {
	article.UserID = user.ID

	s.mutex.Lock()
	s.articles[article.ID] = article
	s.mutex.Unlock()

	return nil
}

func (s *ArticleLocalStorage) GetArticles(ctx context.Context) ([]*models.Article, error) {
	articles := make([]*models.Article, 0)

	s.mutex.Lock()
	for _, value := range s.articles {
		articles = append(articles, value)
	}
	s.mutex.Unlock()

	return articles, nil
}

func (s *ArticleLocalStorage) GetArticlesByUserID(ctx context.Context, user *models.User) ([]*models.Article, error) {
	articles := make([]*models.Article, 0)

	s.mutex.Lock()
	for _, article := range s.articles {
		if article.ID == user.ID {
			articles = append(articles, article)
		}
	}
	s.mutex.Unlock()

	return articles, nil
}

func (s *ArticleLocalStorage) GetArticleByID(ctx context.Context, id string) (*models.Article, error) {
	return s.articles[id], nil
}

func (s *ArticleLocalStorage) UpdateArticle(ctx context.Context, newArticle *models.Article, id string) error {
	s.mutex.Lock()
	s.articles[id] = newArticle
	s.mutex.Unlock()

	return nil
}

func (s *ArticleLocalStorage) DeleteArticle(ctx context.Context, id string) error {
	s.mutex.Lock()
	delete(s.articles, id)
	s.mutex.Unlock()

	return nil
}
