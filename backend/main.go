package main

import (
	"net/http"

	"github.com/jphacks/B_2206/backend/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	infrastructure.Init()
	e.GET("/", healthCheck)
	e.Logger.Fatal(e.Start(":1323"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Health Check")
}
