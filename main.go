package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/kubeblog/backend/article"
	"github.com/kubeblog/backend/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const uploadDir = "uploads"

func main() {
	db := database.ConnectDB()
	database.AutoMigrate(db)

	articleRepository := article.NewArticleRepository(db)
	articleSvc := article.NewArticleService(articleRepository)

	e := echo.New()

	articleSvc.Route(e)

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create upload directory: %v", err)
		}
	}

	e.POST("/upload", uploadFile)
	e.Static("/uploads", uploadDir)

	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":8080"))
}

func uploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid file"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot open file"})
	}
	defer src.Close()

	filename := uuid.New().String()

	filePath := filepath.Join(uploadDir, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save file"})
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to copy file"})
	}

	fileURL := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)
	return c.JSON(http.StatusOK, echo.Map{"message": "File uploaded successfully", "url": fileURL})
}
