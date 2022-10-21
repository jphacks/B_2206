package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type valueController struct {
	u usecase.ValueUseCase
}

type ValueController interface {
	IndexValue(echo.Context) error
	ShowValue(echo.Context) error
	CreateValue(echo.Context) error
	UpdateValue(echo.Context) error
	DestroyValue(echo.Context) error
	IndexValueWithSponsorAndStyle(echo.Context) error
}

func NewValueController(u usecase.ValueUseCase) ValueController {
	return &valueController{u}
}

// Index
func (a *valueController) IndexValue(c echo.Context) error {
	activities, err := a.u.GetValues(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *valueController) ShowValue(c echo.Context) error {
	id := c.Param("id")
	value, err := a.u.GetValueByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, value)
}

// Create
func (a *valueController) CreateValue(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreateValue(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created Value")
}

// Update
func (a *valueController) UpdateValue(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdateValue(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated Value")
}

// Destroy
func (a *valueController) DestroyValue(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyValue(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy Value")
}

// For admin view
func (a *valueController) IndexValueWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetValuesWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

