package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type classificationController struct {
	u usecase.ClassificationUseCase
}

type ClassificationController interface {
	IndexClassification(echo.Context) error
	ShowClassification(echo.Context) error
	CreateClassification(echo.Context) error
	UpdateClassification(echo.Context) error
	DestroyClassification(echo.Context) error
	//classificationに紐づくyearとsourceの取得
	ShowClassificationWithYearAndSource(echo.Context) error
}

func NewClassificationController(u usecase.ClassificationUseCase) ClassificationController {
	return &classificationController{u}
}

// Index
func (b *classificationController) IndexClassification(c echo.Context) error {
	classifications, err := b.u.GetClassifications(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, classifications)
}

// Show
func (b *classificationController) ShowClassification(c echo.Context) error {
	id := c.Param("id")
	classification, err := b.u.GetClassificationByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, classification)
}

// Create
func (b *classificationController) CreateClassification(c echo.Context) error {
	price := c.QueryParam("price")
	yearID := c.QueryParam("year_id")
	sourceID := c.QueryParam("source_id")
	err := b.u.CreateClassification(c.Request().Context(), price, yearID, sourceID)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created Classification")
}

// Update
func (b *classificationController) UpdateClassification(c echo.Context) error {
	id := c.Param("id")
	price := c.QueryParam("price")
	yearID := c.QueryParam("year_id")
	sourceID := c.QueryParam("source_id")
	err := b.u.UpdateClassification(c.Request().Context(), id, price, yearID, sourceID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated Classification")
}

// Destroy
func (b *classificationController) DestroyClassification(c echo.Context) error {
	id := c.Param("id")
	err := b.u.DestroyClassification(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy Classification")
}

//Show Classification with year and source 
func (b *classificationController) ShowClassificationWithYearAndSource(c echo.Context) error {
	id := c.Param("id")
	classificationyearsource ,err :=b.u.GetClassificationWithYearAndSource(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, classificationyearsource)
} 
