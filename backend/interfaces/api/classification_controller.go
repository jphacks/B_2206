package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ClassificationController struct {
	Interactor usecase.ClassificationInteractor
}

func NewClassificationController(sqlHandler database.SqlHandler) *ClassificationController {
	return &ClassificationController{
		Interactor: usecase.ClassificationInteractor{
			ClassificationRepository: &database.ClassificationRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /classifications
func (controller *ClassificationController) CreateClassification(c echo.Context, name string) (err error) {
	u := domain.Classification{
		Name: name,
	}
	c.Bind(&u)
	classification, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, classification)
	return
}

// endpoint GET /classifications/:id/
func (controller *ClassificationController) GetClassification(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	classification, err := controller.Interactor.ClassificationById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, classification)
	return
}

// endpoint GET /classifications/
func (controller *ClassificationController) GetClassifications(c echo.Context) (err error) {
	classifications, err := controller.Interactor.Classifications()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, classifications)
	return
}

// endpoint UPDATE /classifications/:id/
func (controller *ClassificationController) UpdateClassification(c echo.Context, name string) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))

	// idをClassification構造体のIDフィールドに格納
	classification := domain.Classification{
		ID:   id,
		Name: name,
	}
	// classificationをUpdate()に代入
	classification, err = controller.Interactor.Update(classification)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, classification)
	return
}

// endpoint DELETE /classifications/:id/
func (controller *ClassificationController) DeleteClassification(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをClassification構造体のIDフィールドに格納
	classification := domain.Classification{ID: id}
	// classificationをDeleteByClassification()に代入
	err = controller.Interactor.DeleteByClassification(classification)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, classification)
	return
}
