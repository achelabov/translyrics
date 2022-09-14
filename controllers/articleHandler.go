package controllers

import (
	"fmt"
	"net/http"

	"github.com/achelabov/translyrics/models"
	"github.com/gin-gonic/gin"
)

type jsonArticleInput struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func CreateArticle(ctx *gin.Context) {
	inp := new(jsonArticleInput)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//	user := ctx.MustGet("user").(*models.User)

	err := dbArticles.CreateArticle(ctx.Request.Context(), &models.Article{
		Title: inp.Title,
		Text:  inp.Text,
	}, &models.User{ID: "228"})
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func GetAllArticles(ctx *gin.Context) {
	articles, err := dbArticles.GetArticles(ctx.Request.Context())
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")

	article, err := dbArticles.GetArticleByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	inp := new(jsonArticleInput)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//user := ctx.MustGet("user").(*models.User)

	err := dbArticles.UpdateArticle(ctx.Request.Context(),
		&models.Article{Title: inp.Title, Text: inp.Text}, &models.User{ID: "228"}, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	//user := ctx.MustGet("user").(*models.User)

	if err := dbArticles.DeleteArticle(ctx.Request.Context(), &models.User{ID: "228"}, id); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
