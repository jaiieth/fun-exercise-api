package main

import (
	"github.com/KKGo-Software-engineering/fun-exercise-api/helper"
	"github.com/KKGo-Software-engineering/fun-exercise-api/middleware"
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	db, err := postgres.New()

	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger)
	e.Validator = helper.NewValidator()

	walletHandler := wallet.New(db)
	walletHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
