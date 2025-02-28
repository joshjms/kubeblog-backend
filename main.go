package main

import (
	"github.com/kubeblog/backend/article"
	"github.com/kubeblog/backend/database"
	"github.com/labstack/echo/v4"
)

func main() {
	db := database.ConnectDB()
	database.AutoMigrate(db)

	articleRepository := article.NewArticleRepository(db)
	articleSvc := article.NewArticleService(articleRepository)

	e := echo.New()

	articleSvc.Route(e)

	e.Logger.Fatal(e.Start(":8080"))
}
