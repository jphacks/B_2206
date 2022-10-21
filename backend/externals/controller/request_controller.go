package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type requestController struct {
	u usecase.RequestUseCase
}

type RequestController interface {
	IndexRequest(echo.Context) error
	ShowRequest(echo.Context) error
	CreateRequest(echo.Context) error
	UpdateRequest(echo.Context) error
	DestroyRequest(echo.Context) error
	IndexRequestWithSponsorAndStyle(echo.Context) error
}

func NewRequestController(u usecase.RequestUseCase) RequestController {
	return &requestController{u}
}

// Index
func (a *requestController) IndexRequest(c echo.Context) error {
	activities, err := a.u.GetRequests(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *requestController) ShowRequest(c echo.Context) error {
	id := c.Param("id")
	request, err := a.u.GetRequestByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, request)
}

// Create
func (a *requestController) CreateRequest(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreateRequest(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created Request")
}

// Update
func (a *requestController) UpdateRequest(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdateRequest(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated Request")
}

// Destroy
func (a *requestController) DestroyRequest(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyRequest(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy Request")
}

// For admin view
func (a *requestController) IndexRequestWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetRequestsWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

