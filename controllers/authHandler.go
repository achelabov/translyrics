package controllers

import (
	"net/http"

	"github.com/achelabov/translyrics/auth"
	"github.com/achelabov/translyrics/models"
	"github.com/gin-gonic/gin"
)

type signInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(ctx *gin.Context) {
	inp := new(signInput)

	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := &models.User{
		Username: inp.Username,
		Email:    inp.Email,
		Password: inp.Password,
	}

	if err := user.HashPassword(); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := dbUsers.CreateUser(ctx.Request.Context(), user); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func SignIn(ctx *gin.Context) {
	inp := new(signInput)
	if err := ctx.ShouldBindJSON(&inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// check if email exists and password is correct
	user, err := dbUsers.GetUser(ctx.Request.Context(), inp.Username)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	credentialError := user.CheckPassword(inp.Password)
	if credentialError != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
