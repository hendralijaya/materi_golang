package main

import (
	"log"
	"materi-golang/app/config"
	"materi-golang/app/routes"
	"materi-golang/helper"
	"materi-golang/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	err := godotenv.Load()
	helper.PanicIfError(err)
	router := setupRouter()
	log.Fatal(router.Run(":" + os.Getenv("GO_PORT")))
}

func setupRouter() *gin.Engine {
	err := godotenv.Load()
	helper.PanicIfError(err)
	router := gin.Default()
	if os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if os.Getenv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	routes.NewBookRoutes(db, router)
	router.Use(middleware.ErrorHandler)
	router.Use(cors.Default())
	return router
}