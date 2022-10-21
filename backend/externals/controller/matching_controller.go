package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type matchingController struct {
	u usecase.MatchingUseCase
}

type MatchingController interface {
	IndexMatching(echo.Context) error
	ShowMatching(echo.Context) error
	CreateMatching(echo.Context) error
	UpdateMatching(echo.Context) error
	DestroyMatching(echo.Context) error
	IndexMatchingWithSponsorAndStyle(echo.Context) error
}

func NewMatchingController(u usecase.MatchingUseCase) MatchingController {
	return &matchingController{u}
}

// Index
func (a *matchingController) IndexMatching(c echo.Context) error {
	activities, err := a.u.GetMatchings(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *matchingController) ShowMatching(c echo.Context) error {
	id := c.Param("id")
	matching, err := a.u.GetMatchingByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, matching)
}

// Create
func (a *matchingController) CreateMatching(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreateMatching(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created Matching")
}

// Update
func (a *matchingController) UpdateMatching(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdateMatching(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated Matching")
}

// Destroy
func (a *matchingController) DestroyMatching(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyMatching(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy Matching")
}

// For admin view
func (a *matchingController) IndexMatchingWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetMatchingsWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

