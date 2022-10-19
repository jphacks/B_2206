package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type CompanyInfoController struct {
	Interactor usecase.CompanyInfoInteractor
}

func NewCompanyInfoController(sqlHandler database.SqlHandler) *CompanyInfoController {
	return &CompanyInfoController{
		Interactor: usecase.CompanyInfoInteractor{
			CompanyInfoRepository: &database.CompanyInfoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /companyInfos
func (controller *CompanyInfoController) CreateCompanyInfo(c echo.Context, name string, phoneNumber int, postCode int, addressNumber int, buildingName string, website string) (err error) {
	u := domain.CompanyInfo{
		Name:          name,
		PhoneNumber:   phoneNumber,
		PostCode:      postCode,
		AddressNumber: addressNumber,
		BuildingName:  buildingName,
		Website:       website,
	}
	c.Bind(&u)
	companyInfo, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, companyInfo)
	return
}

// endpoint GET /companyInfos/:id/
func (controller *CompanyInfoController) GetCompanyInfo(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	companyInfo, err := controller.Interactor.CompanyInfoById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, companyInfo)
	return
}

// endpoint GET /companyInfos/
func (controller *CompanyInfoController) GetCompanyInfos(c echo.Context) (err error) {
	companyInfos, err := controller.Interactor.CompanyInfos()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, companyInfos)
	return
}

// endpoint UPDATE /companyInfos/:id/
func (controller *CompanyInfoController) UpdateCompanyInfo(c echo.Context, name string, phoneNumber int, postCode int, addressNumber int, buildingName string, website string) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをCompanyInfo構造体のIDフィールドに格納
	companyInfo := domain.CompanyInfo{
		ID:            id,
		Name:          name,
		PhoneNumber:   phoneNumber,
		PostCode:      postCode,
		AddressNumber: addressNumber,
		BuildingName:  buildingName,
		Website:       website,
	}
	// companyInfoをUpdate()に代入
	companyInfo, err = controller.Interactor.Update(companyInfo)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, companyInfo)
	return
}

// endpoint DELETE /companyInfos/:id/
func (controller *CompanyInfoController) DeleteCompanyInfo(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをCompanyInfo構造体のIDフィールドに格納
	companyInfo := domain.CompanyInfo{ID: id}
	// companyInfoをDeleteByCompanyInfo()に代入
	err = controller.Interactor.DeleteByCompanyInfo(companyInfo)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, companyInfo)
	return
}
