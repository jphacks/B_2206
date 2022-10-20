package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type WatchListController struct {
	Interactor usecase.WatchListInteractor
}

func NewWatchListController(sqlHandler database.SqlHandler) *WatchListController {
	return &WatchListController{
		Interactor: usecase.WatchListInteractor{
			WatchListRepository: &database.WatchListRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /watchLists
func (controller *WatchListController) CreateWatchList(c echo.Context, description string, isPurchase bool, userId int, detailId int) (err error) {
	u := domain.WatchList{
		Description: description,
		IsPurchase:  isPurchase,
		UserId:      userId,
		DetailId:    detailId,
	}
	c.Bind(&u)
	watchList, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, watchList)
	return
}

// endpoint GET /watchLists/:id/
func (controller *WatchListController) GetWatchList(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	watchList, err := controller.Interactor.WatchListById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, watchList)
	return
}

// endpoint GET /watchLists/
func (controller *WatchListController) GetWatchLists(c echo.Context) (err error) {
	watchLists, err := controller.Interactor.WatchLists()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, watchLists)
	return
}

// endpoint UPDATE /watchLists/:id/
func (controller *WatchListController) UpdateWatchList(c echo.Context, description string, isPurchase bool, userId int, detailId int) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))

	// idをWatchList構造体のIDフィールドに格納
	watchList := domain.WatchList{
		ID:          id,
		Description: description,
		IsPurchase:  isPurchase,
		UserId:      userId,
		DetailId:    detailId,
	}
	// watchListをUpdate()に代入
	watchList, err = controller.Interactor.Update(watchList)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, watchList)
	return
}

// endpoint DELETE /watchLists/:id/
func (controller *WatchListController) DeleteWatchList(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをWatchList構造体のIDフィールドに格納
	watchList := domain.WatchList{ID: id}
	// watchListをDeleteByWatchList()に代入
	err = controller.Interactor.DeleteByWatchList(watchList)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, watchList)
	return
}
