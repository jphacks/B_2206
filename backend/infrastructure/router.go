package infrastructure

import (
	controllers "github.com/jphacks/B_2206/backend/interfaces/api"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	userController := controllers.NewUserController(NewSqlHandler())
	personalInfoController := controllers.NewPersonalInfoController(NewSqlHandler())
	watchListController := controllers.NewWatchListController(NewSqlHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//users
	e.POST("/users", func(c echo.Context) error {
		name := c.QueryParam("name")
		email := c.QueryParam("email")
		password := c.QueryParam("password")
		PersonalInfoId, _ := strconv.Atoi(c.QueryParam("personal_info_id"))
		return userController.CreateUser(c, name, email, password, PersonalInfoId)
	})
	e.GET("/users", func(c echo.Context) error { return userController.GetUsers(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.GetUser(c) })
	e.PUT("/users/:id", func(c echo.Context) error {
		name := c.QueryParam("name")
		email := c.QueryParam("email")
		password := c.QueryParam("password")
		PersonalInfoId, _ := strconv.Atoi(c.QueryParam("personal_info_id"))
		return userController.UpdateUser(c, name, email, password, PersonalInfoId)
	})
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.DeleteUser(c) })

	//personal_info
	e.POST("/personal-info", func(c echo.Context) error {
		familyName := c.QueryParam("family_name")
		firstName := c.QueryParam("first_name")
		familyNameKana := c.QueryParam("family_name_kana")
		firstNameKana := c.QueryParam("first_name_kana")
		birthday := c.QueryParam("birthday")
		phoneNumber := c.QueryParam("phone_number")
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		return personalInfoController.CreatePersonalInfo(c, familyName, firstName, familyNameKana, firstNameKana, birthday, phoneNumber, userId)
	})
	e.GET("/personal-info", func(c echo.Context) error { return personalInfoController.GetPersonalInfos(c) })
	e.GET("/personal-info/:id", func(c echo.Context) error { return personalInfoController.GetPersonalInfo(c) })
	e.PUT("/personal-info/:id", func(c echo.Context) error {
		familyName := c.QueryParam("family_name")
		firstName := c.QueryParam("first_name")
		familyNameKana := c.QueryParam("family_name_kana")
		firstNameKana := c.QueryParam("first_name_kana")
		birthday := c.QueryParam("birthday")
		phoneNumber := c.QueryParam("phone_number")
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		return personalInfoController.UpdatePersonalInfo(c, familyName, firstName, familyNameKana, firstNameKana, birthday, phoneNumber, userId)
	})
	e.DELETE("/personal-info/:id", func(c echo.Context) error { return personalInfoController.DeletePersonalInfo(c) })

	// watch_list
	e.POST("/watch-list", func(c echo.Context) error {
		description := c.QueryParam("description")
		isPurchase,_ := strconv.ParseBool(c.QueryParam("is_purchase")	)
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		detailId, _ := strconv.Atoi(c.QueryParam("detail_id"))
		return watchListController.CreateWatchList(c, description, isPurchase, userId, detailId)
	})
	e.GET("/watch-list", func(c echo.Context) error { return watchListController.GetWatchLists(c) })
	e.GET("/watch-list/:id", func(c echo.Context) error { return watchListController.GetWatchList(c) })
	e.PUT("/watch-list/:id", func(c echo.Context) error {
		description := c.QueryParam("description")
		isPurchase,_ := strconv.ParseBool(c.QueryParam("is_purchase")	)
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		detailId, _ := strconv.Atoi(c.QueryParam("detail_id"))
		return watchListController.UpdateWatchList(c, description, isPurchase, userId, detailId)
	})
	e.DELETE("/watch-list/:id", func(c echo.Context) error { return watchListController.DeleteWatchList(c) })

	e.Logger.Fatal(e.Start(":1323"))

}
