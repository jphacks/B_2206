package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type limitController struct {
	u usecase.LimitUseCase
}

type LimitController interface {
	IndexLimit(echo.Context) error
	ShowLimit(echo.Context) error
	CreateLimit(echo.Context) error
	UpdateLimit(echo.Context) error
	DestroyLimit(echo.Context) error
	IndexLimitWithSponsorAndStyle(echo.Context) error
}

func NewLimitController(u usecase.LimitUseCase) LimitController {
	return &limitController{u}
}

// Index
func (a *limitController) IndexLimit(c echo.Context) error {
	activities, err := a.u.GetLimits(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *limitController) ShowLimit(c echo.Context) error {
	id := c.Param("id")
	limit, err := a.u.GetLimitByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, limit)
}

// Create
func (a *limitController) CreateLimit(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreateLimit(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created Limit")
}

// Update
func (a *limitController) UpdateLimit(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdateLimit(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated Limit")
}

// Destroy
func (a *limitController) DestroyLimit(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyLimit(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy Limit")
}

// For admin view
func (a *limitController) IndexLimitWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetLimitsWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

