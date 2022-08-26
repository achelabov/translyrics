package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	h "github.com/achelabov/translyrics/controllers/handlers"
	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	//	TODO init db here

	return &App{}
}

func (a *App) Run(port string) error {
	router := initRouter()

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
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api")
	{
		articles := api.Group("/articles")
		{
			articles.POST("/", h.CreateArticle)
			articles.GET("/", h.GetAllArticles)
			articles.GET("/:id", h.GetArticleById)
			articles.GET("/:id", h.UpdateArticle)
			articles.GET("/:id", h.DeleteArticle)
		}
	}

	return router
}
