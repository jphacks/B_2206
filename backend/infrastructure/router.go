package infrastructure

import (
	"github.com/jphacks/B_2206/backend/interfaces/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func Init() {
	e := echo.New()
	userController := controllers.NewUserController(NewSqlHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//users
	e.POST("/users", func(c echo.Context) error {
		name := c.QueryParam("name")
		email := c.QueryParam("email")
		password := c.QueryParam("password")
		personalInfoId := c.QueryParam("personal_info_id")
		companyInfoId := c.QueryParam("company_info_id")
		pInfoId, _ := strconv.Atoi(personalInfoId)
		cInfoId, _ := strconv.Atoi(companyInfoId)
		return userController.CreateUser(c, name, email, password, pInfoId, cInfoId)
	})
	e.GET("/users", func(c echo.Context) error { return userController.GetUsers(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.GetUser(c) })
	e.PUT("/users/:id", func(c echo.Context) error {
		name := c.QueryParam("name")
		email := c.QueryParam("email")
		password := c.QueryParam("password")
		personalInfoId := c.QueryParam("personal_info_id")
		companyInfoId := c.QueryParam("company_info_id")
		pInfoId, _ := strconv.Atoi(personalInfoId)
		cInfoId, _ := strconv.Atoi(companyInfoId)
		return userController.UpdateUser(c, name, email, password, pInfoId, cInfoId)
	})
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.DeleteUser(c) })

	e.Logger.Fatal(e.Start(":1323"))
}
