package database

import (
	"context"
	"log"
	"time"

	"github.com/achelabov/translyrics/config"
	ldb "github.com/achelabov/translyrics/database/local"
	mdb "github.com/achelabov/translyrics/database/mongo"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseAccess struct {
	ArticleStorage
	UserStorage
}

var mongodb *mongo.Database = initMongoDB()
var MongoAccess *DatabaseAccess = newMongoDatabaseAccess()

func newMongoDatabaseAccess() *DatabaseAccess {
	return &DatabaseAccess{
		ArticleStorage: mdb.NewArticleMongoStorage(mongodb, viper.GetString("mongo.article_collection")),
		UserStorage:    mdb.NewUserMongoStorage(mongodb, viper.GetString("mongo.user_collection")),
	}
}

func newLocalDatabaseAccess() *DatabaseAccess {
	return &DatabaseAccess{
		ArticleStorage: ldb.NewArticleLocalStorage(),
		UserStorage:    ldb.NewUserLocalStorage(),
	}
}

func initMongoDB() *mongo.Database {
	if err := config.Init(); err != nil {
		log.Fatalf(err.Error())
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("mongo.name"))
}
