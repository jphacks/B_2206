package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type wishListController struct {
	u usecase.WishListUseCase
}

type WishListController interface {
	IndexWishList(echo.Context) error
	ShowWishList(echo.Context) error
	CreateWishList(echo.Context) error
	UpdateWishList(echo.Context) error
	DestroyWishList(echo.Context) error
	IndexWishListWithSponsorAndStyle(echo.Context) error
}

func NewWishListController(u usecase.WishListUseCase) WishListController {
	return &wishListController{u}
}

// Index
func (a *wishListController) IndexWishList(c echo.Context) error {
	activities, err := a.u.GetWishLists(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *wishListController) ShowWishList(c echo.Context) error {
	id := c.Param("id")
	wishList, err := a.u.GetWishListByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, wishList)
}

// Create
func (a *wishListController) CreateWishList(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreateWishList(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created WishList")
}

// Update
func (a *wishListController) UpdateWishList(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdateWishList(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated WishList")
}

// Destroy
func (a *wishListController) DestroyWishList(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyWishList(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy WishList")
}

// For admin view
func (a *wishListController) IndexWishListWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetWishListsWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

