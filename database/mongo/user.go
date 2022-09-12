package mongo

import (
	"context"

	"github.com/achelabov/translyrics/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type UserMongoStorage struct {
	db *mongo.Collection
}

func NewUserMongoStorage(db *mongo.Database, collection string) *UserMongoStorage {
	return &UserMongoStorage{
		db: db.Collection(collection),
	}
}

func (s *UserMongoStorage) CreateUser(ctx context.Context, user *models.User) error {
	mongoUser := toMongoUser(user)

	res, err := s.db.InsertOne(ctx, mongoUser)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (s *UserMongoStorage) GetUser(ctx context.Context, username string) (*models.User, error) {
	user := new(User)

	err := s.db.FindOne(ctx, bson.M{
		"username": username,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return toModelUser(user), nil
}

//TODO
func (s *UserMongoStorage) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func toMongoUser(u *models.User) *User {
	return &User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

func toModelUser(u *User) *models.User {
	return &models.User{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
