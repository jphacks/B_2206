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
	companyInfoController := controllers.NewCompanyInfoController(NewSqlHandler())
	detailController := controllers.NewDetailController(NewSqlHandler())
	requestController := controllers.NewRequestController(NewSqlHandler())
	areaController := controllers.NewAreaController(NewSqlHandler())
	classificationController := controllers.NewClassificationController(NewSqlHandler())
	tagController := controllers.NewTagController(NewSqlHandler())
	matchingController := controllers.NewMatchingController(NewSqlHandler())
	limitController := controllers.NewLimitController(NewSqlHandler())

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
		isPurchase, _ := strconv.ParseBool(c.QueryParam("is_purchase"))
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		detailId, _ := strconv.Atoi(c.QueryParam("detail_id"))
		return watchListController.CreateWatchList(c, description, isPurchase, userId, detailId)
	})
	e.GET("/watch-list", func(c echo.Context) error { return watchListController.GetWatchLists(c) })
	e.GET("/watch-list/:id", func(c echo.Context) error { return watchListController.GetWatchList(c) })
	e.PUT("/watch-list/:id", func(c echo.Context) error {
		description := c.QueryParam("description")
		isPurchase, _ := strconv.ParseBool(c.QueryParam("is_purchase"))
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		detailId, _ := strconv.Atoi(c.QueryParam("detail_id"))
		return watchListController.UpdateWatchList(c, description, isPurchase, userId, detailId)
	})
	e.DELETE("/watch-list/:id", func(c echo.Context) error { return watchListController.DeleteWatchList(c) })

	//companyInfos
	e.POST("/company-info", func(c echo.Context) error {
		name := c.QueryParam("name")
		phoneNumber := c.QueryParam("phone_number")
		postCode := c.QueryParam("post_code")
		addressNumber := c.QueryParam("address_number")
		buildingName := c.QueryParam("building_name")
		website := c.QueryParam("website")
		return companyInfoController.CreateCompanyInfo(c, name, phoneNumber, postCode, addressNumber, buildingName, website)
	})
	e.GET("/company-info", func(c echo.Context) error { return companyInfoController.GetCompanyInfos(c) })
	e.GET("/company-info/:id", func(c echo.Context) error { return companyInfoController.GetCompanyInfo(c) })
	e.PUT("/company-info/:id", func(c echo.Context) error {
		name := c.QueryParam("name")
		phoneNumber := c.QueryParam("phone_number")
		postCode := c.QueryParam("post_code")
		addressNumber := c.QueryParam("address_number")
		buildingName := c.QueryParam("building_name")
		website := c.QueryParam("website")
		return companyInfoController.UpdateCompanyInfo(c, name, phoneNumber, postCode, addressNumber, buildingName, website)
	})
	e.DELETE("/company-info/:id", func(c echo.Context) error { return companyInfoController.DeleteCompanyInfo(c) })

	//details
	e.POST("/detail", func(c echo.Context) error {
		areaId, _ := strconv.Atoi(c.QueryParam("area_id"))
		return detailController.CreateDetail(c, areaId)
	})
	e.GET("/detail", func(c echo.Context) error { return detailController.GetDetails(c) })
	e.GET("/detail/:id", func(c echo.Context) error { return detailController.GetDetail(c) })
	e.PUT("/detail/:id", func(c echo.Context) error {
		areaId, _ := strconv.Atoi(c.QueryParam("area_id"))
		return detailController.UpdateDetail(c, areaId)
	})
	e.DELETE("/detail/:id", func(c echo.Context) error { return detailController.DeleteDetail(c) })

	// request
	e.POST("/request", func(c echo.Context) error {
		description := c.QueryParam("description")
		isPurchase, _ := strconv.ParseBool(c.QueryParam("is_purchase"))
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		detailId, _ := strconv.Atoi(c.QueryParam("detail_id"))
		return requestController.CreateRequest(c, description, isPurchase, userId, detailId)
	})
	e.GET("/request", func(c echo.Context) error { return requestController.GetRequests(c) })
	e.GET("/request/:id", func(c echo.Context) error { return requestController.GetRequest(c) })
	e.PUT("/request/:id", func(c echo.Context) error {
		description := c.QueryParam("description")
		isPurchase, _ := strconv.ParseBool(c.QueryParam("is_purchase"))
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		detailId, _ := strconv.Atoi(c.QueryParam("detail_id"))
		return requestController.UpdateRequest(c, description, isPurchase, userId, detailId)
	})
	e.DELETE("/request/:id", func(c echo.Context) error { return requestController.DeleteRequest(c) })

	//areas
	e.POST("/area", func(c echo.Context) error {
		postCode := c.QueryParam("post_code")
		prefecture := c.QueryParam("prefecture")
		city := c.QueryParam("city")
		addressNumber := c.QueryParam("address_number")
		buildingName := c.QueryParam("building_name")
		return areaController.CreateArea(c, postCode, prefecture, city, addressNumber, buildingName)
	})
	e.GET("/area", func(c echo.Context) error { return areaController.GetAreas(c) })
	e.GET("/area/:id", func(c echo.Context) error { return areaController.GetArea(c) })
	e.PUT("/area/:id", func(c echo.Context) error {
		postCode := c.QueryParam("post_code")
		prefecture := c.QueryParam("prefecture")
		city := c.QueryParam("city")
		addressNumber := c.QueryParam("address_number")
		buildingName := c.QueryParam("building_name")
		return areaController.UpdateArea(c, postCode, prefecture, city, addressNumber, buildingName)
	})
	e.DELETE("/area/:id", func(c echo.Context) error { return areaController.DeleteArea(c) })

	//classification
	e.POST("/classification", func(c echo.Context) error {
		name := c.QueryParam("name")
		return classificationController.CreateClassification(c, name)
	})
	e.GET("/classification", func(c echo.Context) error { return classificationController.GetClassifications(c) })
	e.GET("/classification/:id", func(c echo.Context) error { return classificationController.GetClassification(c) })
	e.PUT("/classification/:id", func(c echo.Context) error {
		name := c.QueryParam("name")
		return classificationController.UpdateClassification(c, name)
	})
	e.DELETE("/classification/:id", func(c echo.Context) error { return classificationController.DeleteClassification(c) })

	//tag
	e.POST("/tag", func(c echo.Context) error {
		name := c.QueryParam("name")
		return tagController.CreateTag(c, name)
	})
	e.GET("/tag", func(c echo.Context) error { return tagController.GetTags(c) })
	e.GET("/tag/:id", func(c echo.Context) error { return tagController.GetTag(c) })
	e.PUT("/tag/:id", func(c echo.Context) error {
		name := c.QueryParam("name")
		return tagController.UpdateTag(c, name)
	})
	e.DELETE("/tag/:id", func(c echo.Context) error { return tagController.DeleteTag(c) })
	// matching
	e.POST("/matching", func(c echo.Context) error {
		status := c.QueryParam("status")
		sellerMessage := c.QueryParam("seller_message")
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		requestId, _ := strconv.Atoi(c.QueryParam("request_id"))
		buyerId, _ := strconv.Atoi(c.QueryParam("buyer_id"))
		sellerId, _ := strconv.Atoi(c.QueryParam("seller_id"))
		return matchingController.CreateMatching(c, status, sellerMessage, userId, requestId, buyerId, sellerId)
	})
	e.GET("/matching", func(c echo.Context) error { return matchingController.GetMatchings(c) })
	e.GET("/matching/:id", func(c echo.Context) error { return matchingController.GetMatching(c) })
	e.PUT("/matching/:id", func(c echo.Context) error {
		status := c.QueryParam("status")
		sellerMessage := c.QueryParam("seller_message")
		userId, _ := strconv.Atoi(c.QueryParam("user_id"))
		requestId, _ := strconv.Atoi(c.QueryParam("request_id"))
		buyerId, _ := strconv.Atoi(c.QueryParam("buyer_id"))
		sellerId, _ := strconv.Atoi(c.QueryParam("seller_id"))
		return matchingController.UpdateMatching(c, status, sellerMessage, userId, requestId, buyerId, sellerId)
	})
	e.DELETE("/matching/:id", func(c echo.Context) error { return matchingController.DeleteMatching(c) })

	//limit
	e.POST("/limit", func(c echo.Context) error {
		name := c.QueryParam("name")
		maxValueId, _ := strconv.Atoi(c.QueryParam("max_value_id"))
		minValueId, _ := strconv.Atoi(c.QueryParam("min_value_id"))
		return limitController.CreateLimit(c, name, maxValueId, minValueId)
	})
	e.GET("/limit", func(c echo.Context) error { return limitController.GetLimits(c) })
	e.GET("/limit/:id", func(c echo.Context) error { return limitController.GetLimit(c) })
	e.PUT("/limit/:id", func(c echo.Context) error {
		name := c.QueryParam("name")
		maxValueId, _ := strconv.Atoi(c.QueryParam("max_value_id"))
		minValueId, _ := strconv.Atoi(c.QueryParam("min_value_id"))
		return limitController.UpdateLimit(c, name, maxValueId, minValueId)
	})
	e.DELETE("/limit/:id", func(c echo.Context) error { return limitController.DeleteLimit(c) })

	e.Logger.Fatal(e.Start(":1323"))

}
