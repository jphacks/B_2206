package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type watchListController struct {
	u usecase.WatchListUseCase
}

type WatchListController interface {
	IndexWatchList(echo.Context) error
	ShowWatchList(echo.Context) error
	CreateWatchList(echo.Context) error
	UpdateWatchList(echo.Context) error
	DestroyWatchList(echo.Context) error
	IndexWatchListWithSponsorAndStyle(echo.Context) error
}

func NewWatchListController(u usecase.WatchListUseCase) WatchListController {
	return &watchListController{u}
}

// Index
func (a *watchListController) IndexWatchList(c echo.Context) error {
	activities, err := a.u.GetWatchLists(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *watchListController) ShowWatchList(c echo.Context) error {
	id := c.Param("id")
	watchList, err := a.u.GetWatchListByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, watchList)
}

// Create
func (a *watchListController) CreateWatchList(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreateWatchList(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created WatchList")
}

// Update
func (a *watchListController) UpdateWatchList(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdateWatchList(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated WatchList")
}

// Destroy
func (a *watchListController) DestroyWatchList(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyWatchList(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy WatchList")
}

// For admin view
func (a *watchListController) IndexWatchListWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetWatchListsWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

