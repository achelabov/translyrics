package controllers

import (
	"github.com/achelabov/translyrics/database"
)

var dbArticles = database.MongoAccess.ArticleStorage
var dbUsers = database.MongoAccess.UserStorage
