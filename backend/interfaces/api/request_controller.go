package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type RequestController struct {
	Interactor usecase.RequestInteractor
}

func NewRequestController(sqlHandler database.SqlHandler) *RequestController {
	return &RequestController{
		Interactor: usecase.RequestInteractor{
			RequestRepository: &database.RequestRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /requests
func (controller *RequestController) CreateRequest(c echo.Context, description string, isPurchase bool, userId int, detailId int) (err error) {
	u := domain.Request{
		Description: description,
		IsPurchase:  isPurchase,
		UserId:      userId,
		DetailId:    detailId,
	}
	c.Bind(&u)
	request, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, request)
	return
}

// endpoint GET /requests/:id/
func (controller *RequestController) GetRequest(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	request, err := controller.Interactor.RequestById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, request)
	return
}

// endpoint GET /requests/
func (controller *RequestController) GetRequests(c echo.Context) (err error) {
	requests, err := controller.Interactor.Requests()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, requests)
	return
}

// endpoint UPDATE /requests/:id/
func (controller *RequestController) UpdateRequest(c echo.Context, description string, isPurchase bool, userId int, detailId int) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))

	// idをRequest構造体のIDフィールドに格納
	request := domain.Request{
		ID:          id,
		Description: description,
		IsPurchase:  isPurchase,
		UserId:      userId,
		DetailId:    detailId,
	}
	// requestをUpdate()に代入
	request, err = controller.Interactor.Update(request)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, request)
	return
}

// endpoint DELETE /requests/:id/
func (controller *RequestController) DeleteRequest(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをRequest構造体のIDフィールドに格納
	request := domain.Request{ID: id}
	// requestをDeleteByRequest()に代入
	err = controller.Interactor.DeleteByRequest(request)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, request)
	return
}
