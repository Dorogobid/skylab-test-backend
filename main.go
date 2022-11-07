package main

import (
	_ "github.com/dorogobid/skylab-test-backend/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//159.223.220.58

// @title SkyLab test application API
// @version 1.0.0
// @host 159.223.220.58:8080
// @BasePath /
func main() {
	if err := initConfig(); err != nil {
		panic("Error initializing configs")
	}

	db := &DBManager{}
	db.ConnectToDb(getConfig())

	var h HandlerInterface = &Handler{db: db}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n"}))

	e.GET("/quiz/category", h.GetAllCategories)
	e.POST("/quiz/category", h.AddCategory)

	e.GET("/quiz/question/:category_id", h.GetQuestionsByCategoryID)
	e.POST("/quiz/question", h.AddQuestion)

	e.GET("/api/docs/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(viper.GetString("port")))
}