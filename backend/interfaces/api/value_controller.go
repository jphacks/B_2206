package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ValueController struct {
	Interactor usecase.ValueInteractor
}

func NewValueController(sqlHandler database.SqlHandler) *ValueController {
	return &ValueController{
		Interactor: usecase.ValueInteractor{
			ValueRepository: &database.ValueRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /values
func (controller *ValueController) CreateValue(c echo.Context, name string ,number int) (err error) {
	u := domain.Value{
		Name: name,
		Number: number,
	}
	c.Bind(&u)
	value, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, value)
	return
}

// endpoint GET /values/:id/
func (controller *ValueController) GetValue(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	value, err := controller.Interactor.ValueById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, value)
	return
}

// endpoint GET /values/
func (controller *ValueController) GetValues(c echo.Context) (err error) {
	values, err := controller.Interactor.Values()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, values)
	return
}

// endpoint UPDATE /values/:id/
func (controller *ValueController) UpdateValue(c echo.Context, name string, number int) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))

	// idをValue構造体のIDフィールドに格納
	value := domain.Value{
		ID:   id,
		Name: name,
		Number: number,
	}
	// valueをUpdate()に代入
	value, err = controller.Interactor.Update(value)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, value)
	return
}

// endpoint DELETE /values/:id/
func (controller *ValueController) DeleteValue(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをValue構造体のIDフィールドに格納
	value := domain.Value{ID: id}
	// valueをDeleteByValue()に代入
	err = controller.Interactor.DeleteByValue(value)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, value)
	return
}
