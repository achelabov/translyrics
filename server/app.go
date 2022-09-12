package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/achelabov/translyrics/config"
	handler "github.com/achelabov/translyrics/controllers"
	"github.com/achelabov/translyrics/database"
	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server

	db database.DatabaseAccess
}

func NewApp() *App {
	return &App{
		db: *database.MongoAccess,
	}
}

func (a *App) Run(port string) error {
	router := initRouter()
	var _ = config.Init()
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", handler.SignIn)
		auth.POST("/sign-up", handler.SignUp)
	}

	api := router.Group("/api")
	{
		articles := api.Group("/articles")
		{
			articles.POST("/", handler.CreateArticle)
			articles.GET("/", handler.GetAllArticles)
			articles.GET("/:id", handler.GetArticleById)
			articles.PUT("/:id", handler.UpdateArticle)
			articles.DELETE("/:id", handler.DeleteArticle)
		}
	}

	return router
}
