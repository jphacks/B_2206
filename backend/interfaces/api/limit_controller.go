package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type LimitController struct {
	Interactor usecase.LimitInteractor
}

func NewLimitController(sqlHandler database.SqlHandler) *LimitController {
	return &LimitController{
		Interactor: usecase.LimitInteractor{
			LimitRepository: &database.LimitRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /limits
func (controller *LimitController) CreateLimit(c echo.Context, name string, maxValueId int, minValueId int) (err error) {
	u := domain.Limit{
		Name:       name,
		MaxValueId: maxValueId,
		MinValueId: minValueId,
	}
	c.Bind(&u)
	limit, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, limit)
	return
}

// endpoint GET /limits/:id/
func (controller *LimitController) GetLimit(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	limit, err := controller.Interactor.LimitById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, limit)
	return
}

// endpoint GET /limits/
func (controller *LimitController) GetLimits(c echo.Context) (err error) {
	limits, err := controller.Interactor.Limits()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, limits)
	return
}

// endpoint UPDATE /limits/:id/
func (controller *LimitController) UpdateLimit(c echo.Context, name string, maxValueId int, minValueId int) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))

	// idをLimit構造体のIDフィールドに格納
	limit := domain.Limit{
		ID:         id,
		Name:       name,
		MaxValueId: maxValueId,
		MinValueId: minValueId,
	}
	// limitをUpdate()に代入
	limit, err = controller.Interactor.Update(limit)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, limit)
	return
}

// endpoint DELETE /limits/:id/
func (controller *LimitController) DeleteLimit(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをLimit構造体のIDフィールドに格納
	limit := domain.Limit{ID: id}
	// limitをDeleteByLimit()に代入
	err = controller.Interactor.DeleteByLimit(limit)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, limit)
	return
}
