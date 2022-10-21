package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type AreaController struct {
	Interactor usecase.AreaInteractor
}

func NewAreaController(sqlHandler database.SqlHandler) *AreaController {
	return &AreaController{
		Interactor: usecase.AreaInteractor{
			AreaRepository: &database.AreaRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /area
func (controller *AreaController) CreateArea(c echo.Context, postCode string, prefecture string, city string, addressNumber string, buildingName string, detailId int) (err error) {
	u := domain.Area{
		PostCode:      postCode,
		Prefecture:    prefecture,
		City:          city,
		AddressNumber: addressNumber,
		BuildingName:  buildingName,
		DetailId:      detailId,
	}
	c.Bind(&u)
	area, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, area)
	return
}

// endpoint GET /area/:id/
func (controller *AreaController) GetArea(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	area, err := controller.Interactor.AreaById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, area)
	return
}

// endpoint GET /area/
func (controller *AreaController) GetAreas(c echo.Context) (err error) {
	areas, err := controller.Interactor.Areas()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, areas)
	return
}

// endpoint UPDATE /area/:id/
func (controller *AreaController) UpdateArea(c echo.Context, postCode string, prefecture string, city string, addressNumber string, buildingName string, detailId int) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをArea構造体のIDフィールドに格納
	area := domain.Area{
		ID:            id,
		PostCode:      postCode,
		Prefecture:    prefecture,
		City:          city,
		AddressNumber: addressNumber,
		BuildingName:  buildingName,
		DetailId:      detailId,
	}
	// areaをUpdate()に代入
	area, err = controller.Interactor.Update(area)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, area)
	return
}

// endpoint DELETE /area/:id/
func (controller *AreaController) DeleteArea(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをArea構造体のIDフィールドに格納
	area := domain.Area{ID: id}
	// areaをDeleteByArea()に代入
	err = controller.Interactor.DeleteByArea(area)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, area)
	return
}
