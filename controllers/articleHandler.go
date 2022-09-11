package controllers

import (
	"net/http"

	"github.com/achelabov/translyrics/models"
	"github.com/gin-gonic/gin"
)

type jsonInput struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func CreateArticle(ctx *gin.Context) {
	inp := new(jsonInput)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := ctx.MustGet("user").(*models.User)

	err := dbArticles.CreateArticle(ctx.Request.Context(), &models.Article{
		Title: inp.Title,
		Text:  inp.Text,
	}, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func GetAllArticles(ctx *gin.Context) {
	articles, err := dbArticles.GetArticles(ctx.Request.Context())
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func GetArticleById(ctx *gin.Context) {

}

func UpdateArticle(ctx *gin.Context) {

}

func DeleteArticle(ctx *gin.Context) {

}
