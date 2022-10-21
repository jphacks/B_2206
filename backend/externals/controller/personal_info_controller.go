package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type personalInfoController struct {
	u usecase.PersonalInfoUseCase
}

type PersonalInfoController interface {
	IndexPersonalInfo(echo.Context) error
	ShowPersonalInfo(echo.Context) error
	CreatePersonalInfo(echo.Context) error
	UpdatePersonalInfo(echo.Context) error
	DestroyPersonalInfo(echo.Context) error
	IndexPersonalInfoWithSponsorAndStyle(echo.Context) error
}

func NewPersonalInfoController(u usecase.PersonalInfoUseCase) PersonalInfoController {
	return &personalInfoController{u}
}

// Index
func (a *personalInfoController) IndexPersonalInfo(c echo.Context) error {
	activities, err := a.u.GetPersonalInfos(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

// Show
func (a *personalInfoController) ShowPersonalInfo(c echo.Context) error {
	id := c.Param("id")
	personalInfo, err := a.u.GetPersonalInfoByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, personalInfo)
}

// Create
func (a *personalInfoController) CreatePersonalInfo(c echo.Context) error {
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.CreatePersonalInfo(c.Request().Context(), sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created PersonalInfo")
}

// Update
func (a *personalInfoController) UpdatePersonalInfo(c echo.Context) error {
	id := c.Param("id")
	sponsorStyleID := c.QueryParam("sponsor_style_id")
	userID := c.QueryParam("user_id")
	isDone := c.QueryParam("is_done")
	sponsorID := c.QueryParam("sponsor_id")
	err := a.u.UpdatePersonalInfo(c.Request().Context(), id, sponsorStyleID, userID, isDone, sponsorID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated PersonalInfo")
}

// Destroy
func (a *personalInfoController) DestroyPersonalInfo(c echo.Context) error {
	id := c.Param("id")
	err := a.u.DestroyPersonalInfo(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy PersonalInfo")
}

// For admin view
func (a *personalInfoController) IndexPersonalInfoWithSponsorAndStyle(c echo.Context) error {
	activities, err := a.u.GetPersonalInfosWithSponsorAndStyle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activities)
}

