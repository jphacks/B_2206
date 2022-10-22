package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type MatchingController struct {
	Interactor usecase.MatchingInteractor
}

func NewMatchingController(sqlHandler database.SqlHandler) *MatchingController {
	return &MatchingController{
		Interactor: usecase.MatchingInteractor{
			MatchingRepository: &database.MatchingRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /matchings
func (controller *MatchingController) CreateMatching(c echo.Context, status string, sellerMessage string, userId int, requestId int, buyerId int, sellerId int) (err error) {
	u := domain.Matching{
		Status:        status,
		SellerMessage: sellerMessage,
		UserId:        userId,
		RequestId:     requestId,
		BuyerId:       buyerId,
		SellerId:      sellerId,
	}
	c.Bind(&u)
	matching, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, matching)
	return
}

// endpoint GET /matchings/:id/
func (controller *MatchingController) GetMatching(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	matching, err := controller.Interactor.MatchingById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, matching)
	return
}

// endpoint GET /matchings/
func (controller *MatchingController) GetMatchings(c echo.Context) (err error) {
	matchings, err := controller.Interactor.Matchings()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, matchings)
	return
}

// endpoint UPDATE /matchings/:id/
func (controller *MatchingController) UpdateMatching(c echo.Context, status string, sellerMessage string, userId int, requestId int, buyerId int, sellerId int) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))

	// idをMatching構造体のIDフィールドに格納
	matching := domain.Matching{
		ID:            id,
		Status:        status,
		SellerMessage: sellerMessage,
		UserId:        userId,
		RequestId:     requestId,
		BuyerId:       buyerId,
		SellerId:      sellerId,
	}
	// matchingをUpdate()に代入
	matching, err = controller.Interactor.Update(matching)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, matching)
	return
}

// endpoint DELETE /matchings/:id/
func (controller *MatchingController) DeleteMatching(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをMatching構造体のIDフィールドに格納
	matching := domain.Matching{ID: id}
	// matchingをDeleteByMatching()に代入
	err = controller.Interactor.DeleteByMatching(matching)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, matching)
	return
}
