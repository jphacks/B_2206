package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type tagController struct {
	u usecase.TagUseCase
}

type TagController interface {
	IndexTag(echo.Context) error
	ShowTag(echo.Context) error
	CreateTag(echo.Context) error
	UpdateTag(echo.Context) error
	DestroyTag(echo.Context) error
	IndexTagWithSponsorAndStyle(echo.Context) error
}

func NewTagController(u usecase.TagUseCase) TagController {
	return &tagController{u}
}

// Index
func (a *tagController) IndexTag(c echo.Context) error {
	activities, err := a.u.GetTags(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *tagController) ShowTag(c echo.Context) error {
	id := c.Param("id")
	tag, err := a.u.GetTagByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tag)
}

// Create
func (a *tagController) CreateTag(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreateTag(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created Tag")
}

// Update
func (a *tagController) UpdateTag(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdateTag(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated Tag")
}

// Destroy
func (a *tagController) DestroyTag(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyTag(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy Tag")
}

// For admin view
func (a *tagController) IndexTagWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetTagsWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

