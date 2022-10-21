package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type detailController struct {
	u usecase.DetailUseCase
}

type DetailController interface {
	IndexDetail(echo.Context) error
	ShowDetail(echo.Context) error
	CreateDetail(echo.Context) error
	UpdateDetail(echo.Context) error
	DestroyDetail(echo.Context) error
}

func NewDetailController(u usecase.DetailUseCase) DetailController {
	return &detailController{u}
}

// Index
func (d *detailController) IndexDetail(c echo.Context) error {
	details, err := d.u.GetDetails(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, details)
}

// Show
func (d *detailController) ShowDetail(c echo.Context) error {
	id := c.Param("id")
	detail, err := d.u.GetDetailByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, detail)
}

// Create
func (d *detailController) CreateDetail(c echo.Context) error {
	name := c.QueryParam("name")
	err := d.u.CreateDetail(c.Request().Context(), name)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created Detail")
}

// Update
func (d *detailController) UpdateDetail(c echo.Context) error {
	id := c.Param("id")
	name := c.QueryParam("name")
	err := d.u.UpdateDetail(c.Request().Context(), id, name)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated Detail")
}

// Destroy
func (d *detailController) DestroyDetail(c echo.Context) error {
	id := c.Param("id")
	err := d.u.DestroyDetail(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy Detail")
}
