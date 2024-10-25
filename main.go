package main

import (
	"pteridophyte-app/config"
	"pteridophyte-app/server"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := config.InitDB()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Validator = &config.CustomValidator{Validator: validator.New(validator.WithRequiredStructEnabled())}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	server.NewRouter(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
