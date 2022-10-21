package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type areaController struct {
	u usecase.AreaUseCase
}

type AreaController interface {
	IndexArea(echo.Context) error
	ShowArea(echo.Context) error
	CreateArea(echo.Context) error
	UpdateArea(echo.Context) error
	DestroyArea(echo.Context) error
	IndexAreaWithSponsorAndStyle(echo.Context) error
}

func NewAreaController(u usecase.AreaUseCase) AreaController {
	return &areaController{u}
}

// Index
func (a *areaController) IndexArea(c echo.Context) error {
	activities, err := a.u.GetActivities(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *areaController) ShowArea(c echo.Context) error {
	id := c.Param("id")
	area, err := a.u.GetAreaByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, area)
}

// Create
func (a *areaController) CreateArea(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreateArea(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created Area")
}

// Update
func (a *areaController) UpdateArea(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdateArea(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated Area")
}

// Destroy
func (a *areaController) DestroyArea(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyArea(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy Area")
}

// For admin view
func (a *areaController) IndexAreaWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetActivitiesWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

