package mongo

import (
	"context"

	"github.com/achelabov/translyrics/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"userId"`
	Title  string             `bson:"title"`
	Text   string             `bson:"text"`
}

type ArticleMongoStorage struct {
	db *mongo.Collection
}

func NewArticleMongoStorage(db *mongo.Database, collection string) *ArticleMongoStorage {
	return &ArticleMongoStorage{
		db: db.Collection(collection),
	}
}

func (s *ArticleMongoStorage) CreateArticle(ctx context.Context, article *models.Article, user *models.User) error {
	article.UserID = user.ID

	res, err := s.db.InsertOne(ctx, func(a *models.Article) *Article {
		uid, _ := primitive.ObjectIDFromHex(a.UserID)

		return &Article{
			UserID: uid,
			Title:  a.Title,
			Text:   a.Text,
		}
	}(article))
	if err != nil {
		return err
	}

	article.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (s *ArticleMongoStorage) GetArticles(ctx context.Context) ([]*models.Article, error) {
	articles := make([]*Article, 0)

	cur, err := s.db.Find(ctx, bson.D{{}}, options.Find().SetLimit(50))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		article := new(Article)
		err := cur.Decode(&article)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toArticles(articles), nil
}

func (s *ArticleMongoStorage) GetArticlesByUserID(ctx context.Context, user *models.User) ([]*models.Article, error) {
	articles := make([]*Article, 0)
	uid, _ := primitive.ObjectIDFromHex(user.ID)

	cur, err := s.db.Find(ctx, bson.M{"userId": uid})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		article := new(Article)
		err := cur.Decode(&article)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toArticles(articles), nil
}

func (s *ArticleMongoStorage) GetArticleByID(ctx context.Context, id string) (*models.Article, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	var article *models.Article
	err := s.db.FindOne(ctx, bson.M{"_id": objId}).Decode(&article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

//TODO
func (s *ArticleMongoStorage) UpdateArticle(ctx context.Context, newArticle *models.Article, id string) error {
	return nil
}

func (s *ArticleMongoStorage) DeleteArticle(ctx context.Context, user *models.User, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	_, err := s.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
	return err
}

func toArticle(a *Article) *models.Article {
	return &models.Article{
		ID:     a.ID.Hex(),
		UserID: a.UserID.Hex(),
		Title:  a.Title,
		Text:   a.Text,
	}
}

func toArticles(as []*Article) []*models.Article {
	articles := make([]*models.Article, len(as))

	for i, b := range as {
		articles[i] = toArticle(b)
	}

	return articles
}
