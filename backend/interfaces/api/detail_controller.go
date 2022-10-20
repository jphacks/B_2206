package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type DetailController struct {
	Interactor usecase.DetailInteractor
}

func NewDetailController(sqlHandler database.SqlHandler) *DetailController {
	return &DetailController{
		Interactor: usecase.DetailInteractor{
			DetailRepository: &database.DetailRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /details
func (controller *DetailController) CreateDetail(c echo.Context, areaId int) (err error) {
	u := domain.Detail{
		AreaId: areaId,
	}
	c.Bind(&u)
	detail, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, detail)
	return
}

// endpoint GET /details/:id/
func (controller *DetailController) GetDetail(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	detail, err := controller.Interactor.DetailById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, detail)
	return
}

// endpoint GET /details/
func (controller *DetailController) GetDetails(c echo.Context) (err error) {
	details, err := controller.Interactor.Details()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, details)
	return
}

// endpoint UPDATE /detail/:id/
func (controller *DetailController) UpdateDetail(c echo.Context, areaId int) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))

	// idをDetail構造体のIDフィールドに格納
	detail := domain.Detail{
		ID:     id,
		AreaId: areaId,
	}
	// detailをUpdate()に代入
	detail, err = controller.Interactor.Update(detail)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, detail)
	return
}

// endpoint DELETE /details/:id/
func (controller *DetailController) DeleteDetail(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをDetail構造体のIDフィールドに格納
	detail := domain.Detail{ID: id}
	// detailをDeleteByDetail()に代入
	err = controller.Interactor.DeleteByDetail(detail)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, detail)
	return
}
