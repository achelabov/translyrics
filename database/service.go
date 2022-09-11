package database

import (
	"context"
	"log"
	"time"

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
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
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
